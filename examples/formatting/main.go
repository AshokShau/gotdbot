package main

import (
	_ "embed"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/AshokShau/gotdbot"
	"github.com/AshokShau/gotdbot/filters/callbackquery"
)

//go:generate go run ../../scripts/tools/get_tdjson.go

func main() {
	apiID := int32(6)
	apiHash := ""
	botToken := ""

	bot, err := gotdbot.NewClient(apiID, apiHash, botToken, &gotdbot.ClientOpts{
		DatabaseDirectory: "database",
	})

	if err != nil {
		log.Fatal(err)
	}
	defer bot.Close()

	bot.OnCommand("start", func(c *gotdbot.Client, command *gotdbot.Message) error {
		_, err := command.ReplyText(c, "Send a formatted message or media caption, then reply to it with /format.\n"+
			"Modes: /format html, md, mdv2, richhtml, richmd, all.\n"+
			"Use /format source to inspect the generated markup, or /sample for test messages.\n"+
			"Reply with /buttons md (or html/mdv2/none) to resend text or a caption with its buttons.\n"+
			"Use /buttons without a reply to see a button sample.\n"+
			"Rich messages support richhtml/richmd; legacy md loses newer styles and nesting.", &gotdbot.SendMessageOpts{TopicId: command.TopicId})
		return err
	})

	bot.OnCommand("format", formatReply)
	bot.OnCommand("sample", sendSamples)
	bot.OnCommand("buttons", sendButtons)

	bot.OnUpdateNewCallbackQuery(func(c *gotdbot.Client, query *gotdbot.UpdateNewCallbackQuery) error {
		return query.Answer(c, "Done", nil)
	}, callbackquery.Equal("formatting:done"))

	if err = bot.Start(); err != nil {
		log.Fatal(err)
	}

	bot.Logger.Info("Formatting bot started; send /sample or reply with /format")
	bot.Idle()
}

func formatReply(c *gotdbot.Client, command *gotdbot.Message) error {
	replyOpts := &gotdbot.SendMessageOpts{TopicId: command.TopicId}
	if command.ReplyToMessageID() == 0 {
		_, err := command.ReplyText(c, "Reply to text, media or a rich message with /format [mode].", replyOpts)
		return err
	}

	mode := strings.ToLower(strings.TrimSpace(command.Args()))
	switch mode {
	case "":
		mode = "all"
	case gotdbot.ParseModeHTML, "md", "mdv2", "richhtml", "richmd", "all", "source":
	default:
		_, err := command.ReplyText(c, "Modes: html, md, mdv2, richhtml, richmd, all, source.", replyOpts)
		return err
	}

	original, err := c.GetRepliedMessage(command.ChatId, command.Id)
	if err != nil {
		_, replyErr := command.ReplyText(c, "Cannot get the replied message: "+err.Error(), replyOpts)
		return replyErr
	}

	rich, richErr := original.GetRichMessage()
	if richErr == nil && !rich.IsFull {
		rich, err = c.GetFullRichMessage(original.ChatId, original.Id)
		if err != nil {
			_, replyErr := command.ReplyText(c, "Cannot load the full rich message: "+err.Error(), replyOpts)
			return replyErr
		}
		fullMessage := *original
		fullMessage.Content = &gotdbot.MessageRichMessage{Message: rich}
		original = &fullMessage
	}

	_, textErr := original.GetFormattedText()
	_, captionErr := original.GetFormattedCaption()
	if richErr != nil && textErr != nil && captionErr != nil {
		_, err := command.ReplyText(c, "This message has no text, caption or rich content to format.", replyOpts)
		return err
	}

	modes := []string{mode}
	if mode == "all" || mode == "source" {
		modes = []string{gotdbot.ParseModeHTML, "md", "mdv2", "richhtml", "richmd"}
		if richErr == nil {
			modes = []string{"richhtml", "richmd"}
		}
	}

	for _, selectedMode := range modes {
		if err := replyFormatted(c, command, original, selectedMode, mode == "source"); err != nil {
			_, replyErr := command.ReplyText(c, selectedMode+" failed: "+err.Error(), replyOpts)
			if replyErr != nil {
				return replyErr
			}
		}
	}
	return nil
}

func replyFormatted(c *gotdbot.Client, command, original *gotdbot.Message, mode string, showSource bool) error {
	var markup, parseMode string
	var media []gotdbot.InputRichMessageMedia
	var err error

	switch mode {
	case gotdbot.ParseModeHTML:
		parseMode = gotdbot.ParseModeHTML
	case "md":
		parseMode = gotdbot.ParseModeMarkdown
	case "mdv2":
		parseMode = gotdbot.ParseModeMarkdownV2
	case "richhtml":
		markup, media = original.OriginalRichHTML()
	case "richmd":
		markup, media = original.OriginalRichMD()
	}

	if parseMode != gotdbot.ParseModeNone {
		if _, err := original.GetRichMessage(); err == nil {
			return errors.New("use richhtml or richmd for rich messages")
		}
		markup, err = original.OriginalTextWithButtons(parseMode)
		if err != nil {
			return err
		}
	}

	if showSource {
		_, err := command.ReplyText(c, mode+":\n"+markup, &gotdbot.SendMessageOpts{TopicId: command.TopicId})
		return err
	}

	if parseMode == gotdbot.ParseModeNone {
		opts := &gotdbot.SendRichTextOpts{
			SendMessageOpts:       gotdbot.SendMessageOpts{TopicId: command.TopicId},
			Media:                 media,
			DetectAutomaticBlocks: true,
		}
		if rich, err := original.GetRichMessage(); err == nil {
			opts.IsRtl = rich.IsRtl
		}
		if mode == "richhtml" {
			_, err = command.ReplyRichMessageHTML(c, markup, opts)
		} else {
			_, err = command.ReplyRichMessageMarkdown(c, markup, opts)
		}
		return err
	}

	text, keyboard, err := gotdbot.ParseButtonMarkup(markup, parseMode)
	if err != nil {
		return err
	}
	if _, err := original.GetFormattedCaption(); err == nil {
		return replyMedia(c, command, original, text, parseMode, keyboard)
	}
	_, err = command.ReplyText(c, text, &gotdbot.SendMessageOpts{
		ParseMode: parseMode, ReplyMarkup: keyboard,
		TopicId: command.TopicId, DisableWebPagePreview: true,
	})
	return err
}

func replyMedia(c *gotdbot.Client, command, original *gotdbot.Message, caption, parseMode string, keyboard gotdbot.ReplyMarkup) error {
	file := original.GetFile()
	if file == nil {
		return errors.New("this media type cannot be resent by this example")
	}
	input := &gotdbot.InputFileId{Id: file.Id}
	var err error

	switch content := original.Content.(type) {
	case *gotdbot.MessagePhoto:
		opts := &gotdbot.SendPhotoOpts{
			Caption: caption, ParseMode: parseMode, ReplyMarkup: keyboard, TopicId: command.TopicId,
			HasSpoiler: content.HasSpoiler, ShowCaptionAboveMedia: content.ShowCaptionAboveMedia,
		}
		if content.Video != nil && content.Video.Video != nil {
			opts.Video = &gotdbot.InputFileId{Id: content.Video.Video.Id}
		}
		_, err = command.ReplyPhoto(c, input, opts)
	case *gotdbot.MessageVideo:
		_, err = command.ReplyVideo(c, input, &gotdbot.SendVideoOpts{
			Caption: caption, ParseMode: parseMode, ReplyMarkup: keyboard, TopicId: command.TopicId,
			HasSpoiler: content.HasSpoiler, ShowCaptionAboveMedia: content.ShowCaptionAboveMedia,
			Duration: content.Video.Duration, Width: content.Video.Width, Height: content.Video.Height,
			SupportsStreaming: content.Video.SupportsStreaming, StartTimestamp: content.StartTimestamp,
		})
	case *gotdbot.MessageAnimation:
		_, err = command.ReplyAnimation(c, input, &gotdbot.SendAnimationOpts{
			Caption: caption, ParseMode: parseMode, ReplyMarkup: keyboard, TopicId: command.TopicId,
			HasSpoiler: content.HasSpoiler, ShowCaptionAboveMedia: content.ShowCaptionAboveMedia,
			Duration: content.Animation.Duration, Width: content.Animation.Width, Height: content.Animation.Height,
		})
	case *gotdbot.MessageAudio:
		_, err = command.ReplyAudio(c, input, &gotdbot.SendAudioOpts{
			Caption: caption, ParseMode: parseMode, ReplyMarkup: keyboard, TopicId: command.TopicId,
			Duration: content.Audio.Duration, Title: content.Audio.Title, Performer: content.Audio.Performer,
		})
	case *gotdbot.MessageDocument:
		_, err = command.ReplyDocument(c, input, &gotdbot.SendDocumentOpts{
			Caption: caption, ParseMode: parseMode, ReplyMarkup: keyboard, TopicId: command.TopicId,
		})
	case *gotdbot.MessageVoiceNote:
		_, err = command.ReplyVoice(c, input, &gotdbot.SendVoiceOpts{
			Caption: caption, ParseMode: parseMode, ReplyMarkup: keyboard, TopicId: command.TopicId,
			Duration: content.VoiceNote.Duration, Waveform: content.VoiceNote.Waveform,
		})
	default:
		return errors.New("this media type cannot be resent by this example")
	}
	return err
}

func sendButtons(c *gotdbot.Client, command *gotdbot.Message) error {
	replyOpts := &gotdbot.SendMessageOpts{TopicId: command.TopicId}
	parseMode := gotdbot.ParseModeMarkdown
	heading := "*Choose an action*"

	switch strings.ToLower(strings.TrimSpace(command.Args())) {
	case gotdbot.ParseModeHTML:
		parseMode = gotdbot.ParseModeHTML
		heading = "<b>Choose an action</b>"
	case "mdv2":
		parseMode = gotdbot.ParseModeMarkdownV2
	case "none":
		parseMode = gotdbot.ParseModeNone
		heading = "Choose an action"
	case "", "md":
	default:
		_, err := command.ReplyText(c, "Use /buttons html, md, mdv2 or none. Reply to a message to resend its text and buttons.", replyOpts)
		return err
	}

	markup := heading + "\n[Docs](buttonurl:https://core.telegram.org|style=primary)" +
		"\n[Done](buttoncallback:formatting:done:same)" +
		"\n[Copy](buttoncopy:Hello from gotdbot)" +
		"\n[Disabled](buttondisabled:)"

	if command.ReplyToMessageID() != 0 {
		original, err := c.GetRepliedMessage(command.ChatId, command.Id)
		if err != nil {
			_, replyErr := command.ReplyText(c, "Cannot get the replied message: "+err.Error(), replyOpts)
			return replyErr
		}

		markup, err = original.OriginalTextWithButtons(parseMode)
		c.Logger.Infof("markup: %s", markup)
		if err != nil {
			_, replyErr := command.ReplyText(c, "Cannot format this message's text or caption: "+err.Error(), replyOpts)
			return replyErr
		}
	}

	text, keyboard, err := gotdbot.ParseButtonMarkup(markup, parseMode)
	if err != nil {
		_, replyErr := command.ReplyText(c, "Cannot parse the buttons: "+err.Error(), replyOpts)
		return replyErr
	}

	_, err = command.ReplyText(c, text, &gotdbot.SendMessageOpts{
		ParseMode: parseMode, ReplyMarkup: keyboard,
		TopicId: command.TopicId, DisableWebPagePreview: true,
	})

	if err != nil {
		_, replyErr := command.ReplyText(c, "Cannot send the message: "+err.Error(), replyOpts)
		return replyErr
	}
	return nil
}

const sampleHTML = `<b>Bold <i>nested italic</i></b>
<u>Underline</u> <s>Strike</s> <tg-spoiler>Spoiler</tg-spoiler>
<code>a_b * c ` + "`" + ` d \ e</code>
<pre><code class="language-go">fmt.Println("Hello 😀")</code></pre>
<a href="https://example.com/?a=1&amp;b=2">Link</a>
<a href="tg://user?id=777000">Mention</a>
<blockquote>Quote
<b>Second line</b></blockquote>
<blockquote expandable>Expandable
Hidden line</blockquote>
<tg-time unix="1792108800" format="r">A date</tg-time>
Emoji 😀 日本語 हिन्दी é
Special: &lt;tag&gt; &amp; _ * [ ] ( ) ~ ` + "`" + ` &gt; # + - = | { } . ! \`

//go:embed samples/rich.html
var sampleRichHTML string

//go:embed samples/rich.md
var sampleRichMarkdown string

func sendSamples(c *gotdbot.Client, command *gotdbot.Message) error {
	path, err := samplePhoto()
	if err != nil {
		return err
	}
	defer os.Remove(path)
	var failures []error

	if _, err = command.ReplyText(c, sampleHTML, &gotdbot.SendMessageOpts{ParseMode: gotdbot.ParseModeHTML, TopicId: command.TopicId, DisableWebPagePreview: true}); err != nil {
		failures = append(failures, fmt.Errorf("text sample: %w", err))
	}

	if _, err = command.ReplyPhoto(c, &gotdbot.InputFileLocal{Path: path}, &gotdbot.SendPhotoOpts{Caption: sampleHTML, ParseMode: gotdbot.ParseModeHTML, TopicId: command.TopicId}); err != nil {
		failures = append(failures, fmt.Errorf("caption sample: %w", err))
	}

	if _, err = command.ReplyRichMessageHTML(c, sampleRichHTML, &gotdbot.SendRichTextOpts{
		SendMessageOpts: gotdbot.SendMessageOpts{TopicId: command.TopicId}, DetectAutomaticBlocks: true,
	}); err != nil {
		failures = append(failures, fmt.Errorf("rich HTML sample: %w", err))
	}

	if _, err = command.ReplyRichMessageMarkdown(c, sampleRichMarkdown, &gotdbot.SendRichTextOpts{
		SendMessageOpts: gotdbot.SendMessageOpts{TopicId: command.TopicId}, DetectAutomaticBlocks: true,
	}); err != nil {
		failures = append(failures, fmt.Errorf("rich Markdown sample: %w", err))
	}

	if err = errors.Join(failures...); err != nil {
		_, replyErr := command.ReplyText(c, err.Error(), &gotdbot.SendMessageOpts{TopicId: command.TopicId})
		return replyErr
	}

	return nil
}

func samplePhoto() (string, error) {
	img := image.NewRGBA(image.Rect(0, 0, 320, 160))
	for y := range 160 {
		for x := range 320 {
			img.SetRGBA(x, y, color.RGBA{R: uint8(x * 255 / 320), G: uint8(y * 255 / 160), B: 180, A: 255})
		}
	}
	file, err := os.CreateTemp("", "gotdbot-formatting-*.png")
	if err != nil {
		return "", err
	}
	encodeErr := png.Encode(file, img)
	closeErr := file.Close()

	if err = errors.Join(encodeErr, closeErr); err != nil {
		os.Remove(file.Name())
		return "", err
	}
	return file.Name(), nil
}

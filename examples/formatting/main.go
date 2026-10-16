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
		_, err = command.ReplyText(c, "Send a formatted message or media caption, then reply to it with /format.\n"+
			"Modes: /format html, md, mdv2, richhtml, richmd, all.\n"+
			"Use /format source to inspect the generated markup, or /sample for test messages.\n"+
			"Rich messages support richhtml/richmd; legacy md loses newer styles and nesting.", &gotdbot.SendMessageOpts{TopicId: command.TopicId})
		return err
	})

	bot.OnCommand("format", formatReply)
	bot.OnCommand("sample", sendSamples)

	if err = bot.Start(); err != nil {
		log.Fatal(err)
	}

	bot.Logger.Info("Formatting bot started; send /sample or reply with /format")
	bot.Idle()
}

func formatReply(c *gotdbot.Client, command *gotdbot.Message) error {
	replyOpts := &gotdbot.SendMessageOpts{TopicId: command.TopicId}
	if command.ReplyToMessageID() == 0 {
		_, err := command.ReplyText(c, "Reply to a text message, caption or rich message with /format [mode].", replyOpts)
		return err
	}

	source, err := c.GetRepliedMessage(command.ChatId, command.Id)
	if err != nil {
		_, replyErr := command.ReplyText(c, "Cannot get replied message: "+err.Error(), replyOpts)
		return replyErr
	}

	rich, richErr := source.GetRichMessage()
	if richErr == nil && !rich.IsFull {
		// BOT CANT
		rich, err = c.GetFullRichMessage(source.ChatId, source.Id)
		if err != nil {
			_, replyErr := command.ReplyText(c, "Cannot load full rich message: "+err.Error(), replyOpts)
			return replyErr
		}

		copySource := *source
		copySource.Content = &gotdbot.MessageRichMessage{Message: rich}
		source = &copySource
	}

	_, textErr := source.GetFormattedText()
	_, captionErr := source.GetFormattedCaption()
	if richErr != nil && textErr != nil && captionErr != nil {
		_, err = command.ReplyText(c, "This message has no text, supported caption or rich content.", replyOpts)
		return err
	}

	mode := strings.ToLower(strings.TrimSpace(command.Args()))
	if mode == "" {
		mode = "all"
	}

	modes := []string{mode}
	if mode == "all" || mode == "source" {
		modes = []string{"html", "md", "mdv2", "richhtml", "richmd"}
		if richErr == nil {
			modes = []string{"richhtml", "richmd"}
		}
	}

	for _, variant := range modes {
		var markup, parseMode string
		switch variant {
		case "html":
			markup, parseMode = source.OriginalTextHTML(), gotdbot.ParseModeHTML
		case "md":
			markup, parseMode = source.OriginalTextMD(), gotdbot.ParseModeMarkdown
		case "mdv2":
			markup, parseMode = source.OriginalTextMDV2(), gotdbot.ParseModeMarkdownV2
		}

		var media []gotdbot.InputRichMessageMedia
		switch variant {
		case "html", "md", "mdv2":
			if richErr == nil {
				if _, err = command.ReplyText(c, "Rich messages use richhtml or richmd.", replyOpts); err != nil {
					return err
				}
				continue
			}
		case "richhtml":
			markup, media = source.OriginalRichHTML()
		case "richmd":
			markup, media = source.OriginalRichMD()
		default:
			_, err = command.ReplyText(c, "Modes: html, md, mdv2, richhtml, richmd, all, source.", replyOpts)
			return err
		}

		if mode == "source" {
			if _, err = command.ReplyText(c, fmt.Sprintf("%s (%d media references):\n%s", variant, len(media), markup), replyOpts); err != nil {
				return err
			}
			continue
		}

		if _, err = command.ReplyText(c, "Rendering via "+variant+":", replyOpts); err != nil {
			return err
		}
		var sendErr error
		switch variant {
		case "richhtml", "richmd":
			opts := &gotdbot.SendRichTextOpts{
				SendMessageOpts: gotdbot.SendMessageOpts{TopicId: command.TopicId},
				Media:           media, IsRtl: rich != nil && rich.IsRtl, DetectAutomaticBlocks: true,
			}

			if variant == "richhtml" {
				_, sendErr = command.ReplyRichMessageHTML(c, markup, opts)
			} else {
				_, sendErr = command.ReplyRichMessageMarkdown(c, markup, opts)
			}
		default:
			if captionErr == nil {
				showAbove := false
				switch content := source.Content.(type) {
				case *gotdbot.MessagePhoto:
					showAbove = content.ShowCaptionAboveMedia
				case *gotdbot.MessageVideo:
					showAbove = content.ShowCaptionAboveMedia
				case *gotdbot.MessageAnimation:
					showAbove = content.ShowCaptionAboveMedia
				}
				_, sendErr = c.CopyMessage(command.ChatId, source.ChatId, source.Id, &gotdbot.CopyMessageOpts{
					ReplaceCaption: true, NewCaption: markup, ParseMode: parseMode,
					ReplyToMessageID: command.Id, TopicId: command.TopicId,
					NewShowCaptionAboveMedia: showAbove,
				})
			} else {
				_, sendErr = command.ReplyText(c, markup, &gotdbot.SendMessageOpts{ParseMode: parseMode, TopicId: command.TopicId, DisableWebPagePreview: true})
			}
		}
		if sendErr != nil {
			if _, err := command.ReplyText(c, variant+" failed: "+sendErr.Error(), replyOpts); err != nil {
				return err
			}
		}
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

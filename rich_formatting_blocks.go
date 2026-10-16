package gotdbot

import (
	"fmt"
	"html"
	"strconv"
	"strings"
)

type richRenderer struct{ media []InputRichMessageMedia }

func (r *richRenderer) blocks(blocks []PageBlock, mode string) string {
	var out strings.Builder
	for _, block := range blocks {
		out.WriteString(r.block(block, mode))
	}
	return out.String()
}

func (r *richRenderer) block(block PageBlock, mode string) string {
	block = formattingValue(block)
	if block == nil {
		return ""
	}
	if mode == ParseModeMarkdown {
		switch v := block.(type) {
		case PageBlockParagraph:
			return RichTextMarkdown(v.Text) + "\n\n"
		case PageBlockSectionHeading:
			return strings.Repeat("#", max(1, min(6, int(v.Size)))) + " " + RichTextMarkdown(v.Text) + "\n\n"
		case PageBlockDivider:
			return "---\n\n"
		// Use HTML to preserve complex block structure.
		default:
			if markup := r.block(block, ParseModeHTML); markup != "" {
				return markup + "\n"
			}
			return ""
		}
	}
	textBlock := func(tag string, text RichText) string {
		return "<" + tag + ">" + RichTextHTML(text) + "</" + tag + ">\n"
	}
	credit := func(text RichText) string {
		if formattingValue(text) == nil {
			return ""
		}
		return "<cite>" + RichTextHTML(text) + "</cite>"
	}
	switch v := block.(type) {
	case nil:
		return ""
	case PageBlockParagraph:
		return textBlock("p", v.Text)
	case PageBlockSectionHeading:
		return textBlock(fmt.Sprintf("h%d", max(1, min(6, v.Size))), v.Text)
	case PageBlockTitle:
		return textBlock("h1", v.Title)
	case PageBlockSubtitle:
		return textBlock("h2", v.Subtitle)
	case PageBlockHeader:
		return textBlock("h1", v.Header)
	case PageBlockSubheader:
		return textBlock("h2", v.Subheader)
	case PageBlockKicker:
		return textBlock("p", v.Kicker)
	case PageBlockFooter:
		return textBlock("footer", v.Footer)
	case PageBlockThinking:
		return textBlock("tg-thinking", v.Text)
	case PageBlockAuthorDate:
		return textBlock("p", v.Author)
	case PageBlockPreformatted:
		language := ""
		if v.Language != "" {
			language = ` class="language-` + html.EscapeString(v.Language) + `"`
		}
		return "<pre><code" + language + ">" + html.EscapeString(RichTextContent(v.Text)) + "</code></pre>\n"
	case PageBlockDivider:
		return "<hr/>\n"
	case PageBlockAnchor:
		return `<a name="` + html.EscapeString(v.Name) + `"></a>` + "\n"
	case PageBlockMathematicalExpression:
		return "<tg-math-block>" + html.EscapeString(v.Expression) + "</tg-math-block>\n"
	case PageBlockBlockQuote:
		return "<blockquote>\n" + r.blocks(v.Blocks, ParseModeHTML) + credit(v.Credit) + "</blockquote>\n"
	case PageBlockExpandableBlockQuote:
		return "<blockquote expandable>" + RichTextHTML(v.Text) + credit(v.Credit) + "</blockquote>\n"
	case PageBlockPullQuote:
		return "<aside>" + RichTextHTML(v.Text) + credit(v.Credit) + "</aside>\n"
	case PageBlockDetails:
		open := ""
		if v.IsOpen {
			open = " open"
		}
		return "<details" + open + "><summary>" + RichTextHTML(v.Header) + "</summary>\n" + r.blocks(v.Blocks, ParseModeHTML) + "</details>\n"
	case PageBlockList:
		return r.listHTML(v.Items)
	case PageBlockTable:
		return richTableHTML(v)
	case PageBlockCover:
		return r.block(v.Cover, ParseModeHTML)
	case PageBlockCollage:
		return "<tg-collage>\n" + r.blocks(v.Blocks, ParseModeHTML) + richCaptionHTML(v.Caption) + "</tg-collage>\n"
	case PageBlockSlideshow:
		return "<tg-slideshow>\n" + r.blocks(v.Blocks, ParseModeHTML) + richCaptionHTML(v.Caption) + "</tg-slideshow>\n"
	case PageBlockMap:
		if v.Location == nil {
			return richCaptionHTML(v.Caption)
		}
		return richFigure(fmt.Sprintf(`<tg-map lat="%g" long="%g" zoom="%d" width="%d" height="%d"/>`, v.Location.Latitude, v.Location.Longitude, v.Zoom, v.Width, v.Height), v.Caption)
	case PageBlockPhoto:
		if v.Photo != nil {
			var largest *PhotoSize
			for i := range v.Photo.Sizes {
				size := &v.Photo.Sizes[i]
				if size.Photo != nil && (largest == nil || int64(size.Width)*int64(size.Height) > int64(largest.Width)*int64(largest.Height)) {
					largest = size
				}
			}
			if largest != nil {
				return r.mediaHTML("photo", "img", largest.Photo, &InputMessagePhoto{Photo: &InputPhoto{Photo: richInputFile(largest.Photo), Width: largest.Width, Height: largest.Height}}, v.Caption, spoilerAttr(v.HasSpoiler))
			}
		}
		return richCaptionHTML(v.Caption)
	case PageBlockAnimation:
		if v.Animation == nil {
			return richCaptionHTML(v.Caption)
		}
		a := v.Animation
		attrs := spoilerAttr(v.HasSpoiler)
		if v.NeedAutoplay {
			attrs += " autoplay"
		}
		return r.mediaHTML("video", "video", a.Animation, &InputMessageAnimation{Animation: &InputAnimation{Animation: richInputFile(a.Animation), Duration: a.Duration, Width: a.Width, Height: a.Height}}, v.Caption, attrs)
	case PageBlockVideo:
		if v.Video == nil {
			return richCaptionHTML(v.Caption)
		}
		a := v.Video
		attrs := spoilerAttr(v.HasSpoiler)
		if v.NeedAutoplay {
			attrs += " autoplay"
		}
		if v.IsLooped {
			attrs += " loop"
		}
		return r.mediaHTML("video", "video", a.Video, &InputMessageVideo{Video: &InputVideo{Video: richInputFile(a.Video), Duration: a.Duration, Width: a.Width, Height: a.Height, SupportsStreaming: a.SupportsStreaming}}, v.Caption, attrs)
	case PageBlockAudio:
		if v.Audio == nil {
			return richCaptionHTML(v.Caption)
		}
		a := v.Audio
		return r.mediaHTML("audio", "audio", a.Audio, &InputMessageAudio{Audio: &InputAudio{Audio: richInputFile(a.Audio), Duration: a.Duration, Title: a.Title, Performer: a.Performer}}, v.Caption, "")
	case PageBlockVoiceNote:
		if v.VoiceNote == nil {
			return richCaptionHTML(v.Caption)
		}
		a := v.VoiceNote
		return r.mediaHTML("audio", "audio", a.Voice, &InputMessageVoiceNote{VoiceNote: &InputVoiceNote{VoiceNote: richInputFile(a.Voice), Duration: a.Duration, Waveform: a.Waveform}}, v.Caption, "")
	case PageBlockDocument:
		if v.Document == nil {
			return richCaptionHTML(v.Caption)
		}
		a := v.Document
		return r.mediaHTML("document", "tg-document", a.Document, &InputMessageDocument{Document: &InputDocument{Document: richInputFile(a.Document)}}, v.Caption, "")
	case PageBlockButtonRow:
		align := richHorizontalAlignment(v.Align)
		attrs := ""
		if align != "" {
			attrs = ` align="` + align + `"`
		}
		var out strings.Builder
		out.WriteString("<tg-button-row" + attrs + ">")
		for i := range v.Buttons {
			out.WriteString(richButtonHTML(&v.Buttons[i]))
		}
		out.WriteString("</tg-button-row>\n")
		return out.String()
	case PageBlockChatLink:
		return textBlock("p", RichTextUrl{Text: RichTextPlain{Text: v.Title}, Url: "https://t.me/" + v.Username})
	case PageBlockEmbedded:
		return textBlock("p", RichTextUrl{Text: RichTextPlain{Text: v.Url}, Url: v.Url}) + richCaptionHTML(v.Caption)
	case PageBlockEmbeddedPost:
		return r.blocks(v.Blocks, ParseModeHTML) + richCaptionHTML(v.Caption)
	case PageBlockRelatedArticles:
		out := textBlock("h3", v.Header)
		for _, article := range v.Articles {
			out += textBlock("p", RichTextUrl{Text: RichTextPlain{Text: article.Title}, Url: article.Url})
		}
		return out
	}
	return ""
}

func spoilerAttr(spoiler bool) string {
	if spoiler {
		return " tg-spoiler"
	}
	return ""
}

func richInputFile(file *File) InputFile {
	if file == nil {
		return nil
	}
	return &InputFileId{Id: file.Id}
}

func (r *richRenderer) mediaHTML(kind, tag string, file *File, content InputMessageContent, caption *PageBlockCaption, attrs string) string {
	if file == nil || file.Id == 0 {
		return richCaptionHTML(caption)
	}
	id := strconv.Itoa(len(r.media) + 1)
	r.media = append(r.media, InputRichMessageMedia{Id: id, Media: content})
	markup := `<` + tag + ` src="tg://` + kind + `?id=` + id + `"` + attrs
	if tag == "img" {
		markup += "/>"
	} else {
		markup += "></" + tag + ">"
	}
	return richFigure(markup, caption)
}

func richCaptionHTML(caption *PageBlockCaption) string {
	if caption == nil {
		return ""
	}
	text := RichTextHTML(caption.Text)
	if formattingValue(caption.Credit) != nil {
		text += "<cite>" + RichTextHTML(caption.Credit) + "</cite>"
	}
	return "<figcaption>" + text + "</figcaption>\n"
}

func richFigure(content string, caption *PageBlockCaption) string {
	if caption == nil {
		return content + "\n"
	}
	return "<figure>" + content + richCaptionHTML(caption) + "</figure>\n"
}

func (r *richRenderer) listHTML(items []PageBlockListItem) string {
	if len(items) == 0 {
		return ""
	}
	var out strings.Builder
	// Preserve mixed list types and explicit numbering.
	for start := 0; start < len(items); {
		ordered := items[start].Value != 0
		end := start + 1
		for end < len(items) && (items[end].Value != 0) == ordered {
			end++
		}
		tag := "ul"
		if ordered {
			tag = "ol"
		}
		out.WriteString("<" + tag + ">\n")
		for _, item := range items[start:end] {
			attrs := ""
			if ordered {
				attrs = fmt.Sprintf(` value="%d"`, item.Value)
				if item.Type != "" {
					attrs += ` type="` + html.EscapeString(item.Type) + `"`
				}
			}
			out.WriteString("<li" + attrs + ">")
			if item.HasCheckbox {
				out.WriteString(`<input type="checkbox"`)
				if item.IsChecked {
					out.WriteString(" checked")
				}
				out.WriteString(">")
			}
			out.WriteString(r.blocks(item.Blocks, ParseModeHTML))
			out.WriteString("</li>\n")
		}
		out.WriteString("</" + tag + ">\n")
		start = end
	}
	return out.String()
}

func richHorizontalAlignment(align PageBlockHorizontalAlignment) string {
	switch formattingValue(align).(type) {
	case PageBlockHorizontalAlignmentLeft:
		return "left"
	case PageBlockHorizontalAlignmentCenter:
		return "center"
	case PageBlockHorizontalAlignmentRight:
		return "right"
	}
	return ""
}

func richTableHTML(table PageBlockTable) string {
	var out strings.Builder
	out.WriteString("<table")
	if table.IsBordered {
		out.WriteString(" bordered")
	}
	if table.IsStriped {
		out.WriteString(" striped")
	}
	if table.IsCompact {
		out.WriteString(" compact")
	}
	out.WriteString(">\n")
	if formattingValue(table.Caption) != nil {
		out.WriteString("<caption>" + RichTextHTML(table.Caption) + "</caption>\n")
	}
	for _, row := range table.Cells {
		out.WriteString("<tr>")
		for _, cell := range row {
			// Null cells are TDLib placeholders for spanning cells.
			if formattingValue(cell.Text) == nil {
				continue
			}
			tag := "td"
			if cell.IsHeader {
				tag = "th"
			}
			out.WriteString("<" + tag)
			if cell.Colspan > 1 {
				fmt.Fprintf(&out, ` colspan="%d"`, cell.Colspan)
			}
			if cell.Rowspan > 1 {
				fmt.Fprintf(&out, ` rowspan="%d"`, cell.Rowspan)
			}
			if align := richHorizontalAlignment(cell.Align); align != "" {
				out.WriteString(` align="` + align + `"`)
			}
			switch formattingValue(cell.Valign).(type) {
			case PageBlockVerticalAlignmentTop:
				out.WriteString(` valign="top"`)
			case PageBlockVerticalAlignmentMiddle:
				out.WriteString(` valign="middle"`)
			case PageBlockVerticalAlignmentBottom:
				out.WriteString(` valign="bottom"`)
			}
			out.WriteString(">" + RichTextHTML(cell.Text) + "</" + tag + ">")
		}
		out.WriteString("</tr>\n")
	}
	out.WriteString("</table>\n")
	return out.String()
}

func richButtonHTML(button *InlineButton) string {
	if button == nil {
		return ""
	}
	attr := func(name, value string) string { return " " + name + `="` + html.EscapeString(value) + `"` }
	attrs := ""
	switch v := formattingValue(button.Type).(type) {
	case InlineKeyboardButtonTypeUrl:
		attrs = attr("type", "url") + attr("url", v.Url)
	case InlineKeyboardButtonTypeUser:
		attrs = attr("type", "url") + attr("url", fmt.Sprintf("tg://user?id=%d", v.UserId))
	case InlineKeyboardButtonTypeCallback:
		attrs = attr("type", "callback_data") + attr("data", string(v.Data))
	case InlineKeyboardButtonTypeWebApp:
		attrs = attr("type", "web_app") + attr("url", v.Url)
	case InlineKeyboardButtonTypeCopyText:
		attrs = attr("type", "copy_text") + attr("text", v.Text)
	case InlineKeyboardButtonTypeSwitchInline:
		switch target := formattingValue(v.TargetChat).(type) {
		case TargetChatCurrent:
			attrs = attr("type", "switch_inline_query_current_chat")
		case TargetChatChosen:
			attrs = attr("type", "switch_inline_query_chosen_chat")
			if target.Types != nil {
				if target.Types.AllowUserChats {
					attrs += " allow-user-chats"
				}
				if target.Types.AllowBotChats {
					attrs += " allow-bot-chats"
				}
				if target.Types.AllowGroupChats {
					attrs += " allow-group-chats"
				}
				if target.Types.AllowChannelChats {
					attrs += " allow-channel-chats"
				}
			}
		default:
			return RichTextHTML(button.Text)
		}
		attrs += attr("query", v.Query)
	case InlineKeyboardButtonTypeDisabled:
		attrs = attr("type", "disabled")
	default:
		// Received login IDs cannot recreate authorization; preserve the URL.
		if v, ok := formattingValue(button.Type).(InlineKeyboardButtonTypeLoginUrl); ok {
			attrs = attr("type", "url") + attr("url", v.Url)
		} else {
			return RichTextHTML(button.Text)
		}
	}
	switch formattingValue(button.Style).(type) {
	case ButtonStylePrimary:
		attrs += attr("style", "primary")
	case ButtonStyleSuccess:
		attrs += attr("style", "success")
	case ButtonStyleDanger:
		attrs += attr("style", "danger")
	case ButtonStyleLink:
		attrs += attr("style", "link")
	}
	return "<tg-button" + attrs + ">" + RichTextHTML(button.Text) + "</tg-button>"
}

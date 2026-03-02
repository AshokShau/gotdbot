package gotdbot

import (
	"testing"
)

func TestUnparseEntities(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		entities []TextEntity
		mode     string
		expected string
	}{
		{
			name:     "Plain text HTML",
			text:     "Hello <World> & Everyone",
			entities: nil,
			mode:     "html",
			expected: "Hello &lt;World&gt; &amp; Everyone",
		},
		{
			name:     "Plain text MDV2",
			text:     "Hello. World!",
			entities: nil,
			mode:     "markdownv2",
			expected: "Hello\\. World\\!",
		},
		{
			name: "Bold and Italic HTML",
			text: "Bold and Italic",
			entities: []TextEntity{
				{Offset: 0, Length: 4, TypeField: &TextEntityTypeBold{}},
				{Offset: 9, Length: 6, TypeField: &TextEntityTypeItalic{}},
			},
			mode:     "html",
			expected: "<b>Bold</b> and <i>Italic</i>",
		},
		{
			name: "Bold and Italic MDV2",
			text: "Bold and Italic",
			entities: []TextEntity{
				{Offset: 0, Length: 4, TypeField: &TextEntityTypeBold{}},
				{Offset: 9, Length: 6, TypeField: &TextEntityTypeItalic{}},
			},
			mode:     "markdownv2",
			expected: "*Bold* and _Italic_",
		},
		{
			name: "Mention Name HTML",
			text: "User Name",
			entities: []TextEntity{
				{Offset: 0, Length: 9, TypeField: &TextEntityTypeMentionName{UserId: 123456789}},
			},
			mode:     "html",
			expected: "<a href=\"tg://user?id=123456789\">User Name</a>",
		},
		{
			name: "Mention Name MDV2",
			text: "User Name",
			entities: []TextEntity{
				{Offset: 0, Length: 9, TypeField: &TextEntityTypeMentionName{UserId: 123456789}},
			},
			mode:     "markdownv2",
			expected: "[User Name](tg://user?id=123456789)",
		},
		{
			name: "Text URL HTML",
			text: "Open Link",
			entities: []TextEntity{
				{Offset: 0, Length: 9, TypeField: &TextEntityTypeTextUrl{Url: "https://example.com"}},
			},
			mode:     "html",
			expected: "<a href=\"https://example.com\">Open Link</a>",
		},
		{
			name: "Text URL MDV2",
			text: "Open Link",
			entities: []TextEntity{
				{Offset: 0, Length: 9, TypeField: &TextEntityTypeTextUrl{Url: "https://example.com/test?a=1&b=2"}},
			},
			mode:     "markdownv2",
			expected: "[Open Link](https://example.com/test?a=1&b=2)",
		},
		{
			name: "Custom Emoji HTML",
			text: "Smile",
			entities: []TextEntity{
				{Offset: 0, Length: 5, TypeField: &TextEntityTypeCustomEmoji{CustomEmojiId: 123456}},
			},
			mode:     "html",
			expected: "<tg-emoji emoji-id=\"123456\">Smile</tg-emoji>",
		},
		{
			name: "Custom Emoji MDV2",
			text: "Smile",
			entities: []TextEntity{
				{Offset: 0, Length: 5, TypeField: &TextEntityTypeCustomEmoji{CustomEmojiId: 123456}},
			},
			mode:     "markdownv2",
			expected: "![Smile](tg://emoji?id=123456)",
		},
		{
			name: "Date Time HTML",
			text: "tomorrow",
			entities: []TextEntity{
				{Offset: 0, Length: 8, TypeField: &TextEntityTypeDateTime{UnixTime: 1234567890, FormattingType: &DateTimeFormattingTypeRelative{}}},
			},
			mode:     "html",
			expected: "<tg-time unix=\"1234567890\" format=\"r\">tomorrow</tg-time>",
		},
		{
			name: "Date Time MDV2",
			text: "tomorrow",
			entities: []TextEntity{
				{Offset: 0, Length: 8, TypeField: &TextEntityTypeDateTime{UnixTime: 1234567890, FormattingType: &DateTimeFormattingTypeRelative{}}},
			},
			mode:     "markdownv2",
			expected: "[tomorrow](tg://time?unix=1234567890&format=r)",
		},
		{
			name: "Implicit Entity Mention MDV2",
			text: "@username",
			entities: []TextEntity{
				{Offset: 0, Length: 9, TypeField: &TextEntityTypeMention{}},
			},
			mode:     "markdownv2",
			expected: "@username",
		},
		{
			name: "Implicit Entity URL MDV2",
			text: "https://example.com",
			entities: []TextEntity{
				{Offset: 0, Length: 19, TypeField: &TextEntityTypeUrl{}},
			},
			mode:     "markdownv2",
			expected: "https://example.com",
		},
		{
			name: "Pre Code HTML",
			text: "print('hello')",
			entities: []TextEntity{
				{Offset: 0, Length: 14, TypeField: &TextEntityTypePreCode{Language: "python"}},
			},
			mode:     "html",
			expected: "<pre><code class=\"language-python\">print('hello')</code></pre>",
		},
		{
			name: "Pre Code MDV2",
			text: "print('hello')",
			entities: []TextEntity{
				{Offset: 0, Length: 14, TypeField: &TextEntityTypePreCode{Language: "python"}},
			},
			mode:     "markdownv2",
			expected: "```python\nprint('hello')```\n",
		},
		{
			name: "Underline MDV2 Ambiguity Test",
			text: "under",
			entities: []TextEntity{
				{Offset: 0, Length: 5, TypeField: &TextEntityTypeUnderline{}},
			},
			mode:     "markdownv2",
			expected: "__under__",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UnparseEntities(tt.text, tt.entities, tt.mode)
			if result != tt.expected {
				t.Errorf("UnparseEntities() = %v, want %v", result, tt.expected)
			}
		})
	}
}

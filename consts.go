package gotdbot

const (
	Version      = "v0.9.5"
	TDLibVersion = "v1.8.68"
)

const (
	ParseModeHTML       = "HTML"
	ParseModeMarkdownV2 = "MarkdownV2"
	ParseModeMarkdown   = "Markdown"
	ParseModeNone       = ""
)

const (
	supergroupChatIDOffset int64 = -1000000000000
	secretChatIDOffset     int64 = -2000000000000
	minSecretChatID              = secretChatIDOffset - (1 << 31)
	maxSecretChatID              = secretChatIDOffset + (1 << 31) - 1
)

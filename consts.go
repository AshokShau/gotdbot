package gotdbot

const (
	Version      = "v0.9.6"
	TDLibVersion = "v1.8.68"
)

const (
	ParseModeHTML       = "html"
	ParseModeMarkdown   = "markdown"
	ParseModeMarkdownV2 = "markdownv2"
	ParseModeNone       = ""
)

const (
	supergroupChatIDOffset int64 = -1000000000000
	secretChatIDOffset     int64 = -2000000000000
	minSecretChatID              = secretChatIDOffset - (1 << 31)
	maxSecretChatID              = secretChatIDOffset + (1 << 31) - 1
)

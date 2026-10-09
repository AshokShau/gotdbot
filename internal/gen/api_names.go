package main

// Keep TDLib request/type names untouched. Only Go convenience methods that
// collide with handwritten, Bot API-style names need an explicit raw name.
func clientMethodName(tdName string) string {
	switch tdName {
	case "sendMessage":
		return "SendMessageContent"
	case "editMessageText":
		return "EditMessageTextContent"
	case "editMessageCaption":
		return "EditMessageCaptionText"
	default:
		return toCamelCase(tdName)
	}
}

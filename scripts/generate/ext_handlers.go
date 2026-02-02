package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func generateExtHandlers(types []TLType) []string {
	var generatedFiles []string
	skippedHandlers := map[string]bool{
		"updateNewMessage":            true,
		"updateNewCallbackQuery":      true,
		"updateNewInlineQuery":        true,
		"updateNewChosenInlineResult": true,
		"updateNewShippingQuery":      true,
		"updateNewPreCheckoutQuery":   true,
		"updatePoll":                  true,
		"updatePollAnswer":            true,
		"updateChatMember":            true,
		"updateNewChatJoinRequest":    true,
		"updateUser":                  true,
		"updateUserStatus":            true,
		"updateAuthorizationState":    true,
		"updateMessageEdited":         true,
	}

	skippedContext := map[string]bool{
		"updateNewMessage":                     true,
		"updateMessageContent":                 true,
		"updateMessageEdited":                  true,
		"updateMessageSendSucceeded":           true,
		"updateMessageSendFailed":              true,
		"updateMessageContentOpened":           true,
		"updateMessageIsPinned":                true,
		"updateMessageInteractionInfo":         true,
		"updateMessageLiveLocationViewed":      true,
		"updateMessageReaction":                true,
		"updateMessageReactions":               true,
		"updateMessageFactCheck":               true,
		"updateNewCallbackQuery":               true,
		"updateNewInlineCallbackQuery":         true,
		"updateNewBusinessCallbackQuery":       true,
		"updateNewInlineQuery":                 true,
		"updateNewPreCheckoutQuery":            true,
		"updateNewShippingQuery":               true,
		"updateNewCustomQuery":                 true,
		"updateNewChat":                        true,
		"updateChatTitle":                      true,
		"updateChatPhoto":                      true,
		"updateChatPermissions":                true,
		"updateChatLastMessage":                true,
		"updateChatPosition":                   true,
		"updateChatIsMarkedAsUnread":           true,
		"updateChatBlockList":                  true,
		"updateChatDraftMessage":               true,
		"updateChatNotificationSettings":       true,
		"updateChatAvailableReactions":         true,
		"updateChatBackground":                 true,
		"updateChatTheme":                      true,
		"updateChatUnreadMentionCount":         true,
		"updateChatUnreadReactionCount":        true,
		"updateChatVideoChat":                  true,
		"updateChatDefaultDisableNotification": true,
		"updateChatHasProtectedContent":        true,
		"updateChatIsTranslatable":             true,
		"updateChatMember":                     true,
		"updateChatAction":                     true,
		"updateNewChatJoinRequest":             true,
	}

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		if skippedHandlers[t.Name] {
			continue
		}

		structName := toCamelCase(t.Name)
		fileName := camelToSnake(t.Name) + ".go"
		filePath := filepath.Join("ext", "handlers", fileName)

		f, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("Failed to create handler file %s: %v", filePath, err)
		}

		fmt.Fprintln(f, "package handlers")
		fmt.Fprintln(f)
		fmt.Fprintln(f, "import (")
		fmt.Fprintln(f, "\t\"github.com/AshokShau/gotdbot/ext\"")
		fmt.Fprintln(f, "\t\"github.com/AshokShau/gotdbot/ext/handlers/filters\"")
		fmt.Fprintln(f, ")")
		fmt.Fprintln(f)

		fmt.Fprintf(f, "type %s struct {\n", structName)
		fmt.Fprintf(f, "\tFilter   filters.%s\n", structName)
		fmt.Fprintf(f, "\tResponse func(ctx *ext.Context) error\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func New%s(filter filters.%s, response func(ctx *ext.Context) error) *%s {\n", structName, structName, structName)
		fmt.Fprintf(f, "\treturn &%s{\n", structName)
		fmt.Fprintf(f, "\t\tFilter:   filter,\n")
		fmt.Fprintf(f, "\t\tResponse: response,\n")
		fmt.Fprintf(f, "\t}\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (h *%s) CheckUpdate(ctx *ext.Context) bool {\n", structName)
		fmt.Fprintf(f, "\tu := ctx.Update.%s\n", structName)
		fmt.Fprintf(f, "\tif u == nil {\n")
		fmt.Fprintf(f, "\t\treturn false\n")
		fmt.Fprintf(f, "\t}\n")
		fmt.Fprintf(f, "\tif h.Filter == nil {\n")
		fmt.Fprintf(f, "\t\treturn true\n")
		fmt.Fprintf(f, "\t}\n")
		fmt.Fprintf(f, "\treturn h.Filter(u)\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (h *%s) HandleUpdate(ctx *ext.Context) error {\n", structName)
		fmt.Fprintf(f, "\treturn h.Response(ctx)\n")
		fmt.Fprintf(f, "}\n")

		f.Close()
		generatedFiles = append(generatedFiles, filePath)
	}

	filtersPath := filepath.Join("ext", "handlers", "filters", "gen_filters.go")
	fFilters, err := os.Create(filtersPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(fFilters, "package filters")
	fmt.Fprintln(fFilters)
	fmt.Fprintln(fFilters, "import \"github.com/AshokShau/gotdbot\"")
	fmt.Fprintln(fFilters)
	fmt.Fprintln(fFilters, "type (")
	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		if skippedHandlers[t.Name] {
			continue
		}
		structName := toCamelCase(t.Name)
		fmt.Fprintf(fFilters, "\t%s func(u *gotdbot.%s) bool\n", structName, structName)
	}
	fmt.Fprintln(fFilters, ")")
	fFilters.Close()
	generatedFiles = append(generatedFiles, filtersPath)

	contextPath := filepath.Join("ext", "gen_context.go")
	fContext, err := os.Create(contextPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(fContext, "package ext")
	fmt.Fprintln(fContext)
	fmt.Fprintln(fContext, "import \"github.com/AshokShau/gotdbot\"")
	fmt.Fprintln(fContext)

	// Updates Struct
	fmt.Fprintln(fContext, "type Updates struct {")
	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)
		fmt.Fprintf(fContext, "\t%s *gotdbot.%s\n", structName, structName)
	}
	fmt.Fprintln(fContext, "}")
	fmt.Fprintln(fContext)

	// NewUpdates Function
	fmt.Fprintln(fContext, "func NewUpdates(u gotdbot.TlObject) *Updates {")
	fmt.Fprintln(fContext, "\tup := &Updates{}")
	fmt.Fprintln(fContext, "\tswitch u := u.(type) {")
	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)
		fmt.Fprintf(fContext, "\tcase *gotdbot.%s:\n", structName)
		fmt.Fprintf(fContext, "\t\tup.%s = u\n", structName)
	}
	fmt.Fprintln(fContext, "\t}")
	fmt.Fprintln(fContext, "\treturn up")
	fmt.Fprintln(fContext, "}")
	fmt.Fprintln(fContext)

	fmt.Fprintln(fContext, "func extractGeneratedEffectiveFields(u gotdbot.TlObject, c *Context) {")
	fmt.Fprintln(fContext, "\tswitch u := u.(type) {")
	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		if skippedContext[t.Name] {
			continue
		}

		structName := toCamelCase(t.Name)
		hasField := false

		for _, p := range t.Params {
			if p.Name == "chat_id" {
				hasField = true
				break
			}
			if p.Name == "message" && p.Type == "message" {
				hasField = true
				break
			}
			if p.Name == "sender_id" {
				hasField = true
				break
			}
		}
		if !hasField {
			continue
		}

		fmt.Fprintf(fContext, "\tcase *gotdbot.%s:\n", structName)
		for _, p := range t.Params {
			if p.Name == "chat_id" {
				fmt.Fprintf(fContext, "\t\tc.EffectiveChatId = u.ChatId\n")
			} else if p.Name == "message" && p.Type == "message" {
				fmt.Fprintf(fContext, "\t\tif u.Message != nil {\n")
				fmt.Fprintf(fContext, "\t\t\tc.EffectiveMessage = u.Message\n")
				fmt.Fprintf(fContext, "\t\t\tc.EffectiveChatId = u.Message.ChatId\n")
				fmt.Fprintf(fContext, "\t\t\tc.EffectiveSenderId = u.Message.SenderId\n")
				fmt.Fprintf(fContext, "\t\t}\n")
			} else if p.Name == "sender_id" {
				fmt.Fprintf(fContext, "\t\tc.EffectiveSenderId = u.SenderId\n")
			}
		}
	}
	fmt.Fprintln(fContext, "\t}")
	fmt.Fprintln(fContext, "}")
	fContext.Close()
	generatedFiles = append(generatedFiles, contextPath)

	testPath := filepath.Join("ext", "handlers", "gen_handlers_test.go")
	fTest, err := os.Create(testPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(fTest, "package handlers")
	fmt.Fprintln(fTest)
	fmt.Fprintln(fTest, "import (")
	fmt.Fprintln(fTest, "\t\"testing\"")
	fmt.Fprintln(fTest)
	fmt.Fprintln(fTest, "\t\"github.com/AshokShau/gotdbot\"")
	fmt.Fprintln(fTest, "\t\"github.com/AshokShau/gotdbot/ext\"")
	fmt.Fprintln(fTest, ")")
	fmt.Fprintln(fTest)

	fmt.Fprintln(fTest, "func TestGeneratedHandlers(t *testing.T) {")
	fmt.Fprintln(fTest, "\td := ext.NewDispatcher(&gotdbot.Client{})")
	fmt.Fprintln(fTest)

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		if skippedHandlers[t.Name] {
			continue
		}
		structName := toCamelCase(t.Name)

		fmt.Fprintf(fTest, "\tfunc() {\n")
		fmt.Fprintf(fTest, "\t\tcalled := false\n")
		fmt.Fprintf(fTest, "\t\th := New%s(nil, func(ctx *ext.Context) error {\n", structName)
		fmt.Fprintf(fTest, "\t\t\tcalled = true\n")
		fmt.Fprintf(fTest, "\t\t\treturn nil\n")
		fmt.Fprintf(fTest, "\t\t})\n")
		fmt.Fprintf(fTest, "\t\td.AddHandler(h)\n")
		fmt.Fprintf(fTest, "\t\td.ProcessUpdate(&gotdbot.%s{})\n", structName)
		fmt.Fprintf(fTest, "\t\tif !called {\n")
		fmt.Fprintf(fTest, "\t\t\tt.Errorf(\"Handler for %s not called\")\n", structName)
		fmt.Fprintf(fTest, "\t\t}\n")
		fmt.Fprintf(fTest, "\t}()\n\n")
	}
	fmt.Fprintln(fTest, "}")
	fTest.Close()
	generatedFiles = append(generatedFiles, testPath)

	return generatedFiles
}

func camelToSnake(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

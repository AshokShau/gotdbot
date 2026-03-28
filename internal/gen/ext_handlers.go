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

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}

		structName := toCamelCase(t.Name)
		fileName := camelToSnake(t.Name) + ".go"
		filePath := filepath.Join("handlers", fileName)

		var sb strings.Builder
		sb.WriteString(header)
		sb.WriteString("package handlers\n\n")
		sb.WriteString("import (\n")
		sb.WriteString("\t\"github.com/AshokShau/gotdbot\"\n")
		sb.WriteString("\t\"github.com/AshokShau/gotdbot/handlers/filters\"\n")
		sb.WriteString(")\n\n")

		sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
		sb.WriteString(fmt.Sprintf("\tFilter   filters.%s\n", structName))
		sb.WriteString("\tResponse func(b *gotdbot.Client, ctx *gotdbot.Context) error\n")
		sb.WriteString("}\n\n")

		sb.WriteString(fmt.Sprintf("func New%s(filter filters.%s, response func(b *gotdbot.Client, ctx *gotdbot.Context) error) *%s {\n", structName, structName, structName))
		sb.WriteString(fmt.Sprintf("\treturn &%s{\n", structName))
		sb.WriteString("\t\tFilter:   filter,\n")
		sb.WriteString("\t\tResponse: response,\n")
		sb.WriteString("\t}\n")
		sb.WriteString("}\n\n")

		sb.WriteString(fmt.Sprintf("func (h *%s) CheckUpdate(b *gotdbot.Client, ctx *gotdbot.Context) bool {\n", structName))
		sb.WriteString(fmt.Sprintf("\tu := ctx.Update.%s\n", structName))
		sb.WriteString("\tif u == nil {\n")
		sb.WriteString("\t\treturn false\n")
		sb.WriteString("\t}\n")
		sb.WriteString("\tif h.Filter == nil {\n")
		sb.WriteString("\t\treturn true\n")
		sb.WriteString("\t}\n")
		sb.WriteString("\treturn h.Filter(u)\n")
		sb.WriteString("}\n\n")

		sb.WriteString(fmt.Sprintf("func (h *%s) HandleUpdate(b *gotdbot.Client, ctx *gotdbot.Context) error {\n", structName))
		sb.WriteString("\treturn h.Response(b, ctx)\n")
		sb.WriteString("}\n")

		if err := os.WriteFile(filePath, []byte(sb.String()), 0644); err != nil {
			log.Fatalf("Failed to create handler file %s: %v", filePath, err)
		}
		generatedFiles = append(generatedFiles, filePath)
	}

	filtersPath := filepath.Join("handlers", "filters", "gen_filters.go")
	var sbFilters strings.Builder
	sbFilters.WriteString(header)
	sbFilters.WriteString("package filters\n\n")
	sbFilters.WriteString("import \"github.com/AshokShau/gotdbot\"\n\n")
	sbFilters.WriteString("type (\n")

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)
		sbFilters.WriteString(fmt.Sprintf("\t%s func(u *gotdbot.%s) bool\n", structName, structName))
	}

	sbFilters.WriteString("\tMessage func(msg *gotdbot.Message) bool\n")

	sbFilters.WriteString(")\n")
	if err := os.WriteFile(filtersPath, []byte(sbFilters.String()), 0644); err != nil {
		log.Fatal(err)
	}
	generatedFiles = append(generatedFiles, filtersPath)

	contextPath := filepath.Join("gen_context.go")
	var sbContext strings.Builder
	sbContext.WriteString(header)
	sbContext.WriteString("package gotdbot\n\n")

	// Updates Struct
	sbContext.WriteString("type ContextUpdates struct {\n")
	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)
		sbContext.WriteString(fmt.Sprintf("\t%s *%s\n", structName, structName))
	}
	sbContext.WriteString("}\n\n")

	// NewUpdates Function
	sbContext.WriteString("func NewContextUpdates(u TlObject) *ContextUpdates {\n")
	sbContext.WriteString("\tup := &ContextUpdates{}\n")
	sbContext.WriteString("\tswitch u := u.(type) {\n")
	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)
		sbContext.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
		sbContext.WriteString(fmt.Sprintf("\t\tup.%s = u\n", structName))
	}
	sbContext.WriteString("\t}\n")
	sbContext.WriteString("\treturn up\n")
	sbContext.WriteString("}\n\n")

	sbContext.WriteString("func extractGeneratedEffectiveFields(u TlObject, c *Context) {\n")
	sbContext.WriteString("\tswitch u := u.(type) {\n")
	for _, t := range types {
		if t.ResultType != "Update" {
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

		sbContext.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
		for _, p := range t.Params {
			if p.Name == "chat_id" {
				sbContext.WriteString("\t\tc.EffectiveChatId = u.ChatId\n")
			} else if p.Name == "message" && p.Type == "message" {
				sbContext.WriteString("\t\tif u.Message != nil {\n")
				sbContext.WriteString("\t\t\tc.EffectiveMessage = u.Message\n")
				sbContext.WriteString("\t\t\tc.EffectiveChatId = u.Message.ChatId\n")
				sbContext.WriteString("\t\t}\n")
			} else if p.Name == "sender_id" {
				sbContext.WriteString("\t\tif up, ok := u.SenderId.(*MessageSenderUser); ok {\n")
				sbContext.WriteString("\t\t\tc.EffectiveChatId = up.UserId\n")
				sbContext.WriteString("\t\t}\n")
				sbContext.WriteString("\t\tif up, ok := u.SenderId.(*MessageSenderChat); ok {\n")
				sbContext.WriteString("\t\t\tc.EffectiveChatId = up.ChatId\n")
				sbContext.WriteString("\t\t}\n")
			}
		}
	}
	sbContext.WriteString("\t}\n")
	sbContext.WriteString("}\n")
	if err := os.WriteFile(contextPath, []byte(sbContext.String()), 0644); err != nil {
		log.Fatal(err)
	}
	generatedFiles = append(generatedFiles, contextPath)

	testPath := filepath.Join("handlers", "gen_handlers_test.go")
	var sbTest strings.Builder
	sbTest.WriteString(header)
	sbTest.WriteString("package handlers_test\n\n")
	sbTest.WriteString("import (\n")
	sbTest.WriteString("\t\"testing\"\n")
	sbTest.WriteString("\t\"time\"\n\n")
	sbTest.WriteString("\t\"github.com/AshokShau/gotdbot\"\n")
	sbTest.WriteString("\t\"github.com/AshokShau/gotdbot/handlers\"\n")
	sbTest.WriteString(")\n\n")

	sbTest.WriteString("func TestGeneratedHandlers(t *testing.T) {\n")
	sbTest.WriteString("\td := gotdbot.NewDispatcher(&gotdbot.Client{}, nil)\n\n")

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)

		sbTest.WriteString("\tfunc() {\n")
		sbTest.WriteString("\t\tcalled := make(chan bool, 1)\n")
		sbTest.WriteString(fmt.Sprintf("\t\th := handlers.New%s(nil, func(b *gotdbot.Client, ctx *gotdbot.Context) error {\n", structName))
		sbTest.WriteString("\t\t\tcalled <- true\n")
		sbTest.WriteString("\t\t\treturn nil\n")
		sbTest.WriteString("\t\t})\n")
		sbTest.WriteString("\t\td.AddHandler(h)\n")
		sbTest.WriteString(fmt.Sprintf("\t\td.ProcessUpdate(&gotdbot.%s{})\n", structName))
		sbTest.WriteString("\t\tselect {\n")
		sbTest.WriteString("\t\tcase <-called:\n")
		sbTest.WriteString("\t\tcase <-time.After(100 * time.Millisecond):\n")
		sbTest.WriteString(fmt.Sprintf("\t\t\tt.Errorf(\"Handler for %s not called\")\n", structName))
		sbTest.WriteString("\t\t}\n")
		sbTest.WriteString("\t}()\n\n")
	}
	sbTest.WriteString("}\n")
	if err := os.WriteFile(testPath, []byte(sbTest.String()), 0644); err != nil {
		log.Fatal(err)
	}
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

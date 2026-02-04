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
		filePath := filepath.Join("ext", "handlers", fileName)

		var sb strings.Builder
		sb.WriteString(header)
		sb.WriteString("package handlers\n\n")
		sb.WriteString("import (\n")
		sb.WriteString("\t\"github.com/AshokShau/gotdbot/ext\"\n")
		sb.WriteString("\t\"github.com/AshokShau/gotdbot/ext/handlers/filters\"\n")
		sb.WriteString(")\n\n")

		fmt.Fprintf(&sb, "type %s struct {\n", structName)
		fmt.Fprintf(&sb, "\tFilter   filters.%s\n", structName)
		sb.WriteString("\tResponse func(ctx *ext.Context) error\n")
		sb.WriteString("}\n\n")

		fmt.Fprintf(&sb, "func New%s(filter filters.%s, response func(ctx *ext.Context) error) *%s {\n", structName, structName, structName)
		fmt.Fprintf(&sb, "\treturn &%s{\n", structName)
		sb.WriteString("\t\tFilter:   filter,\n")
		sb.WriteString("\t\tResponse: response,\n")
		sb.WriteString("\t}\n")
		sb.WriteString("}\n\n")

		fmt.Fprintf(&sb, "func (h *%s) CheckUpdate(ctx *ext.Context) bool {\n", structName)
		fmt.Fprintf(&sb, "\tu := ctx.Update.%s\n", structName)
		sb.WriteString("\tif u == nil {\n")
		sb.WriteString("\t\treturn false\n")
		sb.WriteString("\t}\n")
		sb.WriteString("\tif h.Filter == nil {\n")
		sb.WriteString("\t\treturn true\n")
		sb.WriteString("\t}\n")
		sb.WriteString("\treturn h.Filter(u)\n")
		sb.WriteString("}\n\n")

		fmt.Fprintf(&sb, "func (h *%s) HandleUpdate(ctx *ext.Context) error {\n", structName)
		sb.WriteString("\treturn h.Response(ctx)\n")
		sb.WriteString("}\n")

		if err := os.WriteFile(filePath, []byte(sb.String()), 0644); err != nil {
			log.Fatalf("Failed to create handler file %s: %v", filePath, err)
		}
		generatedFiles = append(generatedFiles, filePath)
	}

	filtersPath := filepath.Join("ext", "handlers", "filters", "gen_filters.go")
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
		fmt.Fprintf(&sbFilters, "\t%s func(u *gotdbot.%s) bool\n", structName, structName)
	}
	sbFilters.WriteString(")\n")
	if err := os.WriteFile(filtersPath, []byte(sbFilters.String()), 0644); err != nil {
		log.Fatal(err)
	}
	generatedFiles = append(generatedFiles, filtersPath)

	contextPath := filepath.Join("ext", "gen_context.go")
	var sbContext strings.Builder
	sbContext.WriteString(header)
	sbContext.WriteString("package ext\n\n")
	sbContext.WriteString("import \"github.com/AshokShau/gotdbot\"\n\n")

	// Updates Struct
	sbContext.WriteString("type Updates struct {\n")
	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)
		fmt.Fprintf(&sbContext, "\t%s *gotdbot.%s\n", structName, structName)
	}
	sbContext.WriteString("}\n\n")

	// NewUpdates Function
	sbContext.WriteString("func NewUpdates(u gotdbot.TlObject) *Updates {\n")
	sbContext.WriteString("\tup := &Updates{}\n")
	sbContext.WriteString("\tswitch u := u.(type) {\n")
	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)
		fmt.Fprintf(&sbContext, "\tcase *gotdbot.%s:\n", structName)
		fmt.Fprintf(&sbContext, "\t\tup.%s = u\n", structName)
	}
	sbContext.WriteString("\t}\n")
	sbContext.WriteString("\treturn up\n")
	sbContext.WriteString("}\n\n")

	sbContext.WriteString("func extractGeneratedEffectiveFields(u gotdbot.TlObject, c *Context) {\n")
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

		fmt.Fprintf(&sbContext, "\tcase *gotdbot.%s:\n", structName)
		for _, p := range t.Params {
			if p.Name == "chat_id" {
				sbContext.WriteString("\t\tc.EffectiveChatId = u.ChatId\n")
			} else if p.Name == "message" && p.Type == "message" {
				sbContext.WriteString("\t\tif u.Message != nil {\n")
				sbContext.WriteString("\t\t\tc.EffectiveMessage = u.Message\n")
				sbContext.WriteString("\t\t\tc.EffectiveChatId = u.Message.ChatId\n")
				sbContext.WriteString("\t\t}\n")
			} else if p.Name == "sender_id" {
				sbContext.WriteString("\t\tif up, ok := u.SenderId.(*gotdbot.MessageSenderUser); ok {\n")
				sbContext.WriteString("\t\t\tc.EffectiveChatId = up.UserId\n")
				sbContext.WriteString("\t\t}\n")
				sbContext.WriteString("\t\tif up, ok := u.SenderId.(*gotdbot.MessageSenderChat); ok {\n")
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

	testPath := filepath.Join("ext", "handlers", "gen_handlers_test.go")
	var sbTest strings.Builder
	sbTest.WriteString(header)
	sbTest.WriteString("package handlers\n\n")
	sbTest.WriteString("import (\n")
	sbTest.WriteString("\t\"testing\"\n\n")
	sbTest.WriteString("\t\"github.com/AshokShau/gotdbot\"\n")
	sbTest.WriteString("\t\"github.com/AshokShau/gotdbot/ext\"\n")
	sbTest.WriteString(")\n\n")

	sbTest.WriteString("func TestGeneratedHandlers(t *testing.T) {\n")
	sbTest.WriteString("\td := ext.NewDispatcher(&gotdbot.Client{})\n\n")

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)

		sbTest.WriteString("\tfunc() {\n")
		sbTest.WriteString("\t\tcalled := false\n")
		fmt.Fprintf(&sbTest, "\t\th := New%s(nil, func(ctx *ext.Context) error {\n", structName)
		sbTest.WriteString("\t\t\tcalled = true\n")
		sbTest.WriteString("\t\t\treturn nil\n")
		sbTest.WriteString("\t\t})\n")
		sbTest.WriteString("\t\td.AddHandler(h)\n")
		fmt.Fprintf(&sbTest, "\t\td.ProcessUpdate(&gotdbot.%s{})\n", structName)
		sbTest.WriteString("\t\tif !called {\n")
		fmt.Fprintf(&sbTest, "\t\t\tt.Errorf(\"Handler for %s not called\")\n", structName)
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

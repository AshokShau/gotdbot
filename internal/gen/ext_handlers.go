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

		sb.WriteString(fmt.Sprintf("// %s %s\n", structName, formatDesc(t.Description)))
		sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
		sb.WriteString(fmt.Sprintf("\tFilter   filters.%s\n", structName))
		sb.WriteString(fmt.Sprintf("\tResponse func(b *gotdbot.Client, u *gotdbot.%s) error\n", structName))
		sb.WriteString("}\n\n")

		sb.WriteString(fmt.Sprintf("// New%s creates a new %s\n", structName, structName))
		sb.WriteString(fmt.Sprintf("func New%s(filter filters.%s, response func(b *gotdbot.Client, u *gotdbot.%s) error) *%s {\n", structName, structName, structName, structName))
		sb.WriteString(fmt.Sprintf("\treturn &%s{\n", structName))
		sb.WriteString("\t\tFilter:   filter,\n")
		sb.WriteString("\t\tResponse: response,\n")
		sb.WriteString("\t}\n")
		sb.WriteString("}\n\n")

		sb.WriteString(fmt.Sprintf("func (h *%s) CheckUpdate(b *gotdbot.Client, update gotdbot.TlObject) bool {\n", structName))
		sb.WriteString(fmt.Sprintf("\tu, ok := update.(*gotdbot.%s)\n", structName))
		sb.WriteString("\tif !ok {\n")
		sb.WriteString("\t\treturn false\n")
		sb.WriteString("\t}\n")
		sb.WriteString("\tif h.Filter == nil {\n")
		sb.WriteString("\t\treturn true\n")
		sb.WriteString("\t}\n")
		sb.WriteString("\treturn h.Filter(u)\n")
		sb.WriteString("}\n\n")

		sb.WriteString(fmt.Sprintf("func (h *%s) HandleUpdate(b *gotdbot.Client, update gotdbot.TlObject) error {\n", structName))
		sb.WriteString(fmt.Sprintf("\treturn h.Response(b, update.(*gotdbot.%s))\n", structName))
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
		sbFilters.WriteString(fmt.Sprintf("\t// %s %s\n", structName, formatDesc(t.Description)))
		sbFilters.WriteString(fmt.Sprintf("\t%s func(u *gotdbot.%s) bool\n", structName, structName))
	}

	sbFilters.WriteString("\tMessage func(msg *gotdbot.Message) bool\n")

	sbFilters.WriteString(")\n")
	if err := os.WriteFile(filtersPath, []byte(sbFilters.String()), 0644); err != nil {
		log.Fatal(err)
	}
	generatedFiles = append(generatedFiles, filtersPath)

	/*
		    extractorsPath := filepath.Join("gen_extractors.go")
			var sbExtract strings.Builder
			sbExtract.WriteString(header)
			sbExtract.WriteString("package gotdbot\n\n")
			sbExtract.WriteString("// ExtractMessage extracts a Message from an update if possible.\n")
			sbExtract.WriteString("func ExtractMessage(u TlObject) *Message {\n")
			sbExtract.WriteString("\tswitch u := u.(type) {\n")
			for _, t := range types {
				if t.ResultType != "Update" {
					continue
				}
				for _, p := range t.Params {
					if p.Type == "message" {
						fieldName := toCamelCase(p.Name)
						sbExtract.WriteString(fmt.Sprintf("\tcase *%s:\n", toCamelCase(t.Name)))
						sbExtract.WriteString(fmt.Sprintf("\t\treturn u.%s\n", fieldName))
						break
					}
				}
			}
			sbExtract.WriteString("\tdefault:\n")
			sbExtract.WriteString("\t\treturn nil\n")
			sbExtract.WriteString("\t}\n")
			sbExtract.WriteString("}\n\n")

			sbExtract.WriteString("// ExtractChatID extracts a ChatID from an update if possible.\n")
			sbExtract.WriteString("func ExtractChatID(u TlObject) int64 {\n")
			sbExtract.WriteString("\tswitch u := u.(type) {\n")
			for _, t := range types {
				if t.ResultType != "Update" {
					continue
				}
				structName := toCamelCase(t.Name)
				found := false
				for _, p := range t.Params {
					if p.Name == "chat_id" {
						sbExtract.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
						sbExtract.WriteString("\t\treturn u.ChatId\n")
						found = true
						break
					}
					if p.Type == "chat" {
						fieldName := toCamelCase(p.Name)
						sbExtract.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
						sbExtract.WriteString(fmt.Sprintf("\t\tif u.%s != nil {\n", fieldName))
						sbExtract.WriteString(fmt.Sprintf("\t\t\treturn u.%s.Id\n", fieldName))
						sbExtract.WriteString("\t\t}\n")
						found = true
						break
					}
					if p.Type == "user" {
						fieldName := toCamelCase(p.Name)
						sbExtract.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
						sbExtract.WriteString(fmt.Sprintf("\t\tif u.%s != nil {\n", fieldName))
						sbExtract.WriteString(fmt.Sprintf("\t\t\treturn u.%s.Id\n", fieldName))
						sbExtract.WriteString("\t\t}\n")
						found = true
						break
					}
				}
				if found {
					continue
				}
				for _, p := range t.Params {
					if p.Type == "message" {
						fieldName := toCamelCase(p.Name)
						sbExtract.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
						sbExtract.WriteString(fmt.Sprintf("\t\tif u.%s != nil {\n", fieldName))
						sbExtract.WriteString(fmt.Sprintf("\t\t\treturn u.%s.ChatId\n", fieldName))
						sbExtract.WriteString("\t\t}\n")
						found = true
						break
					}
				}
				if found {
					continue
				}
				for _, p := range t.Params {
					if p.Name == "sender_user_id" {
						sbExtract.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
						sbExtract.WriteString("\t\treturn u.SenderUserId\n")
						found = true
						break
					}
					if p.Name == "user_id" {
						sbExtract.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
						sbExtract.WriteString("\t\treturn u.UserId\n")
						found = true
						break
					}
				}
				if found {
					continue
				}
				for _, p := range t.Params {
					if p.Name == "sender_id" {
						sbExtract.WriteString(fmt.Sprintf("\tcase *%s:\n", structName))
						sbExtract.WriteString("\t\tif u.SenderId != nil {\n")
						sbExtract.WriteString("\t\t\tif up, ok := u.SenderId.(*MessageSenderUser); ok {\n")
						sbExtract.WriteString("\t\t\t\treturn up.UserId\n")
						sbExtract.WriteString("\t\t\t}\n")
						sbExtract.WriteString("\t\t\tif up, ok := u.SenderId.(*MessageSenderChat); ok {\n")
						sbExtract.WriteString("\t\t\t\treturn up.ChatId\n")
						sbExtract.WriteString("\t\t\t}\n")
						sbExtract.WriteString("\t\t}\n")
						found = true
						break
					}
				}
			}
			sbExtract.WriteString("\tdefault:\n")
			sbExtract.WriteString("\t\treturn 0\n")
			sbExtract.WriteString("\t}\n")
			sbExtract.WriteString("}\n")

			if err := os.WriteFile(extractorsPath, []byte(sbExtract.String()), 0644); err != nil {
				log.Fatal(err)
			}

	*/
	generatedFiles = append(generatedFiles)
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
	sbTest.WriteString("\tc := &gotdbot.Client{}\n")
	sbTest.WriteString("\td := gotdbot.NewDispatcher(nil)\n\n")

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}
		structName := toCamelCase(t.Name)

		sbTest.WriteString("\tfunc() {\n")
		sbTest.WriteString("\t\tcalled := make(chan bool, 1)\n")
		sbTest.WriteString(fmt.Sprintf("\t\th := handlers.New%s(nil, func(b *gotdbot.Client, u *gotdbot.%s) error {\n", structName, structName))
		sbTest.WriteString("\t\t\tcalled <- true\n")
		sbTest.WriteString("\t\t\treturn nil\n")
		sbTest.WriteString("\t\t})\n")
		sbTest.WriteString("\t\td.AddHandler(h)\n")
		sbTest.WriteString(fmt.Sprintf("\t\td.ProcessUpdate(c, &gotdbot.%s{})\n", structName))
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

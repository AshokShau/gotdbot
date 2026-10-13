package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func generateEvents(types []TLType, classes map[string]*TLClass) {
	var updates []TLType
	for _, t := range types {
		if strings.HasPrefix(t.Name, "update") {
			updates = append(updates, t)
		}
	}
	sortTLTypesByName(updates)

	var fsb strings.Builder
	fsb.WriteString(header)
	fsb.WriteString("package gotdbot\n\n")
	for _, t := range updates {
		structName := toCamelCase(t.Name)
		fmt.Fprintf(&fsb, "// %sFilter filters %s updates.\n", structName, structName)
		if desc := formatDesc(t.Description); desc != "" {
			fmt.Fprintf(&fsb, "// %s\n", desc)
		}
		fmt.Fprintf(&fsb, "type %sFilter func(u *%s) bool\n\n", structName, structName)
	}

	if err := os.WriteFile("gen_filters.go", []byte(fsb.String()), 0644); err != nil {
		log.Fatal(err)
	}

	var sb strings.Builder
	sb.WriteString(header)
	sb.WriteString("package gotdbot\n\n")

	// Handler types
	for _, t := range updates {
		structName := toCamelCase(t.Name)
		name := structName
		unexportedHandlerName := strings.ToLower(name[:1]) + name[1:] + "Handler"

		handlerFunc := "func(client *Client, update *" + structName + ") error"
		filterType := structName + "Filter"

		fmt.Fprintf(&sb, "// %s handles %s updates\n", unexportedHandlerName, structName)
		fmt.Fprintf(&sb, "type %s struct {\n", unexportedHandlerName)
		sb.WriteString("\tfilter   " + filterType + "\n")
		sb.WriteString("\tresponse " + handlerFunc + "\n")
		sb.WriteString("}\n\n")

		fmt.Fprintf(&sb, "func (h *%s) CheckUpdate(client *Client, update TlObject) bool {\n", unexportedHandlerName)
		fmt.Fprintf(&sb, "\tu, ok := update.(*%s)\n", structName)
		sb.WriteString("\tif !ok || u == nil || h.response == nil {\n\t\treturn false\n\t}\n")
		sb.WriteString("\tif h.filter != nil && !h.filter(u) {\n\t\treturn false\n\t}\n")
		sb.WriteString("\treturn true\n")
		sb.WriteString("}\n\n")

		fmt.Fprintf(&sb, "func (h *%s) HandleUpdate(client *Client, update TlObject) error {\n", unexportedHandlerName)
		fmt.Fprintf(&sb, "\tu, ok := update.(*%s)\n\tif !ok || u == nil || h.response == nil { return nil }\n\treturn h.response(client, u)\n", structName)
		sb.WriteString("}\n\n")

		fmt.Fprintf(&sb, "// On%s registers a handler for %s updates with default group (0)\n", name, structName)
		fmt.Fprintf(&sb, "func (c *Client) On%s(handler %s, filter %s) {\n", name, handlerFunc, filterType)
		fmt.Fprintf(&sb, "\tc.Add%sHandlerGroup(handler, filter, 0)\n", name)
		sb.WriteString("}\n\n")

		fmt.Fprintf(&sb, "// Add%sHandlerGroup registers a handler for %s updates with a specific group\n", name, structName)
		fmt.Fprintf(&sb, "func (c *Client) Add%sHandlerGroup(handler %s, filter %s, group int) {\n", name, handlerFunc, filterType)
		fmt.Fprintf(&sb, "\tif handler == nil { return }\n\tc.AddHandlerGroup(&%s{filter: filter, response: handler}, group)\n", unexportedHandlerName)
		sb.WriteString("}\n\n")
	}

	if err := os.WriteFile("gen_events.go", []byte(sb.String()), 0644); err != nil {
		log.Fatal(err)
	}
}

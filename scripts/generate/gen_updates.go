package main

import (
	"fmt"
	"log"
	"os"
)

func generateUpdates(types []TLType) {
	f, err := os.Create("gen_updates.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package gotdbot")
	fmt.Fprintln(f)

	fmt.Fprintln(f)

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}

		structName := toCamelCase(t.Name)
		methodName := "On" + structName

		fmt.Fprintf(f, "// %s %s\n", methodName, formatDesc(t.Description))
		fmt.Fprintf(f, "func (c *Client) %s(fn func(client *Client, update *%s) error, filter FilterFunc, position int) {\n", methodName, structName)
		fmt.Fprintf(f, "\tc.AddHandler(\"%s\", func(cl *Client, u TlObject) error {\n", t.Name)
		fmt.Fprintf(f, "\t\tif up, ok := u.(*%s); ok {\n", structName)
		fmt.Fprintf(f, "\t\t\treturn fn(cl, up)\n")
		fmt.Fprintf(f, "\t\t}\n")
		fmt.Fprintf(f, "\t\treturn nil\n")
		fmt.Fprintf(f, "\t}, filter, position)\n")
		fmt.Fprintf(f, "}\n\n")
	}
}

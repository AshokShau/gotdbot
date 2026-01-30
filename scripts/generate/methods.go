package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func generateMethods(functions []TLType, classes map[string]*TLClass) {
	f, err := os.Create("gen_methods.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package gotdbot")
	fmt.Fprintln(f)

	fmt.Fprintln(f)

	for _, fn := range functions {
		methodName := toCamelCase(fn.Name)
		structName := toCamelCase(fn.Name)

		hasOptional := false
		for _, p := range fn.Params {
			if p.IsOptional {
				hasOptional = true
				break
			}
		}

		optsStructName := methodName + "Opts"

		resultType := toCamelCase(fn.ResultType)
		if fn.ResultType == "ok" {
			resultType = "Ok"
		}

		retTypeStr := "*" + resultType
		if _, ok := classes[fn.ResultType]; ok {
			retTypeStr = "*" + toCamelCase(fn.ResultType)
		}

		fmt.Fprintf(f, "// %s %s\n", methodName, formatDesc(fn.Description))
		fmt.Fprintf(f, "func (c *Client) %s(", methodName)

		// Args
		var args []string
		for _, p := range fn.Params {
			if p.IsOptional {
				continue
			}
			goType := toGoType(p.Type, classes)

			fieldName := toCamelCase(p.Name)
			argName := strings.ToLower(fieldName[:1]) + fieldName[1:]
			if argName == "type" {
				argName = "typeField"
			}
			if argName == "func" {
				argName = "funcArg"
			}

			args = append(args, fmt.Sprintf("%s %s", argName, goType))
		}

		if hasOptional {
			args = append(args, fmt.Sprintf("opts *%s", optsStructName))
		}

		fmt.Fprintf(f, "%s", strings.Join(args, ", "))
		fmt.Fprintf(f, ") (%s, error) {\n", retTypeStr)

		fmt.Fprintf(f, "\treq := &%s{\n", structName)
		for _, p := range fn.Params {
			if p.IsOptional {
				continue
			}
			fieldName := toCamelCase(p.Name)
			if fieldName == "Type" {
				fieldName = "TypeField"
			}
			argName := strings.ToLower(fieldName[:1]) + fieldName[1:]
			if argName == "type" {
				argName = "typeField"
			}
			if argName == "func" {
				argName = "funcArg"
			}

			fmt.Fprintf(f, "\t\t%s: %s,\n", fieldName, argName)
		}
		fmt.Fprintf(f, "\t}\n")

		if hasOptional {
			fmt.Fprintf(f, "\tif opts != nil {\n")
			for _, p := range fn.Params {
				if !p.IsOptional {
					continue
				}
				fieldName := toCamelCase(p.Name)
				if fieldName == "Type" {
					fieldName = "TypeField"
				}
				fmt.Fprintf(f, "\t\treq.%s = opts.%s\n", fieldName, fieldName)
			}
			fmt.Fprintf(f, "\t}\n")
		}

		fmt.Fprintf(f, "\tresp, err := c.Send(req)\n")
		fmt.Fprintf(f, "\tif err != nil {\n\t\treturn nil, err\n\t}\n")
		fmt.Fprintf(f, "\treturn resp.(%s), nil\n", retTypeStr)
		fmt.Fprintf(f, "}\n\n")
	}
}

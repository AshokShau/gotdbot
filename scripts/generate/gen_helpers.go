package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func generateHelpers(types []TLType, functions []TLType, classes map[string]*TLClass) {
	f, err := os.Create("gen_helpers.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package gotdbot")
	fmt.Fprintln(f)
	fmt.Fprintln(f)

	for _, t := range types {
		typeFields := make(map[string]TLParam)
		for _, p := range t.Params {
			typeFields[p.Name] = p
		}

		for _, fn := range functions {
			matches := make(map[string]string)
			matchedCount := 0
			requiredParamCount := 0

			for _, param := range fn.Params {
				if param.IsOptional {
					continue
				}
				requiredParamCount++

				if field, ok := typeFields[param.Name]; ok {
					if field.Type == param.Type {
						matches[param.Name] = toCamelCase(field.Name)
						matchedCount++
						continue
					}
				}

				typeNameSnake := toSnakeCase(t.Name)
				if param.Name == typeNameSnake+"_id" {
					if field, ok := typeFields["id"]; ok {
						if field.Type == param.Type {
							matches[param.Name] = "Id"
							matchedCount++
							continue
						}
					}
				}

				if param.Name == "id" {
					if field, ok := typeFields["id"]; ok {
						if field.Type == param.Type {
							matches[param.Name] = "Id"
							matchedCount++
							continue
						}
					}
				}
			}

			if matchedCount == 0 {
				continue
			}

			methodName := toCamelCase(fn.Name)
			helperName := methodName

			typeCamel := toCamelCase(t.Name)
			if strings.Contains(methodName, typeCamel) {
				helperName = strings.Replace(methodName, typeCamel, "", 1)
			}

			if helperName == "" {
				helperName = methodName
			}

			generateHelperMethod(f, t, fn, matches, helperName, classes)
		}
	}
}

func generateHelperMethod(f *os.File, t TLType, fn TLType, matches map[string]string, helperName string, classes map[string]*TLClass) {
	clientMethodName := toCamelCase(fn.Name)
	structName := toCamelCase(t.Name)
	for _, p := range t.Params {
		if toCamelCase(p.Name) == helperName {
			helperName = helperName + "Action"
		}
	}

	var funcArgs []string
	var callArgs []string

	hasOptional := false
	for _, p := range fn.Params {
		if p.IsOptional {
			hasOptional = true
			continue
		}

		fieldName := toCamelCase(p.Name)
		argName := strings.ToLower(fieldName[:1]) + fieldName[1:]
		if argName == "type" {
			argName = "typeField"
		}
		if argName == "func" {
			argName = "funcArg"
		}

		if receiverField, ok := matches[p.Name]; ok {
			callArgs = append(callArgs, fmt.Sprintf("%s.%s", strings.ToLower(structName[:1]), receiverField))
		} else {
			goType := toGoType(p.Type, classes)
			funcArgs = append(funcArgs, fmt.Sprintf("%s %s", argName, goType))
			callArgs = append(callArgs, argName)
		}
	}

	optsStructName := clientMethodName + "Opts"
	if hasOptional {
		funcArgs = append(funcArgs, fmt.Sprintf("opts *%s", optsStructName))
		callArgs = append(callArgs, "opts")
	}

	resultType := toCamelCase(fn.ResultType)
	if fn.ResultType == "ok" {
		resultType = "Ok"
	}
	retTypeStr := "*" + resultType
	if _, ok := classes[fn.ResultType]; ok {
		retTypeStr = "*" + toCamelCase(fn.ResultType)
	}

	receiverVar := strings.ToLower(structName[:1])

	fmt.Fprintf(f, "// %s is a helper method for Client.%s\n", helperName, clientMethodName)
	fmt.Fprintf(f, "func (%s %s) %s(c *Client, %s) (%s, error) {\n",
		receiverVar, structName, helperName, strings.Join(funcArgs, ", "), retTypeStr)

	fmt.Fprintf(f, "\treturn c.%s(%s)\n", clientMethodName, strings.Join(callArgs, ", "))
	fmt.Fprintf(f, "}\n\n")
}

func toSnakeCase(s string) string {
	var res []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			res = append(res, '_')
		}
		res = append(res, r)
	}
	return strings.ToLower(string(res))
}

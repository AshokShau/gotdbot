package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// generateHelpers splits helper methods into multiple files.
func generateHelpers(types []TLType, functions []TLType, classes map[string]*TLClass) []string {
	fileCount := 0
	helpersPerFile := 1500
	currentHelpers := 0

	f, err := createHelperFile(fileCount)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	generatedFiles := []string{f.Name()}

	allowedTypes := map[string]bool{
		"File":       true,
		"RemoteFile": true,
		"Message":    true,
		"Chat":       true,
		"User":       true,
	}

	manualMethods := map[string]bool{
		"Message.EditText": true,
	}

	generatedMethods := make(map[string]bool)

	for _, t := range types {
		if !allowedTypes[toCamelCase(t.Name)] {
			continue
		}

		typeFields := make(map[string]TLParam)
		for _, p := range t.Params {
			typeFields[p.Name] = p
		}

		for _, fn := range functions {
			matches := make(map[string]string)

			matchedCount := 0

			for _, param := range fn.Params {
				if param.IsOptional {
					continue
				}

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

			if _, hasId := typeFields["id"]; hasId {
				usedId := false
				for _, matchVal := range matches {
					if matchVal == "Id" {
						usedId = true
						break
					}
				}
				if !usedId {
					continue
				}
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

			fullHelperName := toCamelCase(t.Name) + "." + helperName

			if manualMethods[fullHelperName] || generatedMethods[fullHelperName] {
				continue
			}

			generatedMethods[fullHelperName] = true

			if currentHelpers >= helpersPerFile {
				f.Close()
				fileCount++
				currentHelpers = 0

				f, err = createHelperFile(fileCount)
				if err != nil {
					log.Fatal(err)
				}
				generatedFiles = append(generatedFiles, f.Name())
			}

			generateHelperMethod(f, t, fn, matches, helperName, classes)
			currentHelpers++
		}
	}

	return generatedFiles
}

func createHelperFile(index int) (*os.File, error) {
	filename := fmt.Sprintf("gen_helpers%d.go", index)
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	fmt.Fprintln(f, "package gotdbot")
	fmt.Fprintln(f)
	fmt.Fprintln(f)

	return f, nil
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

	fmt.Fprintf(f, "func (%s %s) %s(client *Client, %s) (%s, error) {\n",
		receiverVar, structName, helperName, strings.Join(funcArgs, ", "), retTypeStr)

	fmt.Fprintf(f, "\treturn client.%s(%s)\n", clientMethodName, strings.Join(callArgs, ", "))
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

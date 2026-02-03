package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func generateClasses(classes map[string]*TLClass) {
	f, err := os.Create("gen_classes.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package gotdbot")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import \"encoding/json\"")
	fmt.Fprintln(f, "import \"fmt\"")
	fmt.Fprintln(f)

	fmt.Fprintln(f, "// TlObject is the interface that all TDLib types satisfy")
	fmt.Fprintln(f, "type TlObject interface {")
	fmt.Fprintln(f, "\tType() string")
	fmt.Fprintln(f, "}")
	fmt.Fprintln(f)

	sortedNames := sortKeysAZ(classes)
	for _, name := range sortedNames {
		cls := classes[name]
		structName := toCamelCase(name)
		methodName := toLowerCamelCase(name)

		fmt.Fprintf(f, "// %s %s\n", structName, formatDesc(cls.Description))
		fmt.Fprintf(f, "// Implemented by:\n")
		for _, impl := range cls.Implementations {
			fmt.Fprintf(f, "// %s\n", toCamelCase(impl))
		}
		fmt.Fprintf(f, "type %s interface {\n", structName)
		fmt.Fprintf(f, "\tTlObject\n")
		fmt.Fprintf(f, "\t%s()\n", methodName)
		fmt.Fprintf(f, "}\n\n")

		// Unmarshal helper
		fmt.Fprintf(f, "// Unmarshal%s unmarshals the JSON into the correct concrete implementation of %s\n", structName, structName)
		fmt.Fprintf(f, "func Unmarshal%s(data []byte) (%s, error) {\n", structName, structName)
		fmt.Fprintf(f, "\tvar typeObj struct { Type string `json:\"@type\"` }\n")
		fmt.Fprintf(f, "\tif err := json.Unmarshal(data, &typeObj); err != nil { return nil, err }\n")
		fmt.Fprintf(f, "\tswitch typeObj.Type {\n")
		for _, impl := range cls.Implementations {
			implStructName := toCamelCase(impl)
			fmt.Fprintf(f, "\tcase \"%s\":\n", impl)
			fmt.Fprintf(f, "\t\tvar obj %s\n", implStructName)
			fmt.Fprintf(f, "\t\tif err := json.Unmarshal(data, &obj); err != nil { return nil, err }\n")
			fmt.Fprintf(f, "\t\treturn &obj, nil\n")
		}
		fmt.Fprintf(f, "\tdefault:\n")
		fmt.Fprintf(f, "\t\treturn nil, fmt.Errorf(\"unknown type: %%s\", typeObj.Type)\n")
		fmt.Fprintf(f, "\t}\n")
		fmt.Fprintf(f, "}\n\n")
	}
}

func generateObjects(types []TLType, classes map[string]*TLClass) {
	f, err := os.Create("gen_types.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package gotdbot")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import \"encoding/json\"")
	fmt.Fprintln(f, "import \"fmt\"")
	fmt.Fprintln(f)

	for _, t := range types {
		structName := toCamelCase(t.Name)
		fmt.Fprintf(f, "// %s %s\n", structName, formatDesc(t.Description))
		fmt.Fprintf(f, "type %s struct {\n", structName)

		// Check for fields that are classes to determine if we need custom UnmarshalJSON
		var classFields []TLParam

		for _, p := range t.Params {
			goType := toGoType(p.Type, classes)
			fieldName := toCamelCase(p.Name)
			if fieldName == "Type" {
				fieldName = "TypeField"
			}
			jsonTag := fmt.Sprintf("`json:\"%s\"`", p.Name)
			if p.IsOptional {
				jsonTag = fmt.Sprintf("`json:\"%s,omitempty\"`", p.Name)
			}
			fmt.Fprintf(f, "\t// %s\n", formatDesc(p.Description))
			fmt.Fprintf(f, "\t%s %s %s\n", fieldName, goType, jsonTag)

			// Check if type corresponds to a class
			isClass := false
			typeName := p.Type
			if strings.HasPrefix(typeName, "vector<") {
				typeName = strings.TrimSuffix(strings.TrimPrefix(typeName, "vector<"), ">")
			}
			if _, ok := classes[typeName]; ok {
				isClass = true
			}

			if isClass {
				classFields = append(classFields, p)
			}
		}
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) Type() string {\n", structName)
		fmt.Fprintf(f, "\treturn \"%s\"\n", t.Name)
		fmt.Fprintf(f, "}\n\n")

		// Implement interface dummy method if this type belongs to a class
		if cls, ok := classes[t.ResultType]; ok {
			methodName := toLowerCamelCase(cls.Name)
			fmt.Fprintf(f, "func (t *%s) %s() {}\n\n", structName, methodName)
		}

		// MarshalJSON (only @type)
		fmt.Fprintf(f, "func (t *%s) MarshalJSON() ([]byte, error) {\n", structName)
		fmt.Fprintf(f, "\ttype Alias %s\n", structName)
		fmt.Fprintf(f, "\treturn json.Marshal(&struct {\n")
		fmt.Fprintf(f, "\t\tTypeStr string `json:\"@type\"`\n")
		fmt.Fprintf(f, "\t\t*Alias\n")
		fmt.Fprintf(f, "\t}{\n")
		fmt.Fprintf(f, "\t\tTypeStr: \"%s\",\n", t.Name)
		fmt.Fprintf(f, "\t\tAlias:   (*Alias)(t),\n")
		fmt.Fprintf(f, "\t})\n")
		fmt.Fprintf(f, "}\n\n")

		// UnmarshalJSON if needed
		if len(classFields) > 0 {
			fmt.Fprintf(f, "func (t *%s) UnmarshalJSON(data []byte) error {\n", structName)
			fmt.Fprintf(f, "\ttype Alias %s\n", structName)
			fmt.Fprintf(f, "\taux := &struct {\n")
			for _, p := range classFields {
				fieldName := toCamelCase(p.Name)
				if fieldName == "Type" {
					fieldName = "TypeField"
				}
				if strings.HasPrefix(p.Type, "vector<") {
					fmt.Fprintf(f, "\t\t%s []json.RawMessage `json:\"%s\"`\n", fieldName, p.Name)
				} else {
					fmt.Fprintf(f, "\t\t%s json.RawMessage `json:\"%s\"`\n", fieldName, p.Name)
				}
			}
			fmt.Fprintf(f, "\t\t*Alias\n")
			fmt.Fprintf(f, "\t}{\n")
			fmt.Fprintf(f, "\t\tAlias: (*Alias)(t),\n")
			fmt.Fprintf(f, "\t}\n")
			fmt.Fprintf(f, "\n")
			fmt.Fprintf(f, "\tif err := json.Unmarshal(data, &aux); err != nil {\n")
			fmt.Fprintf(f, "\t\treturn err\n")
			fmt.Fprintf(f, "\t}\n")
			fmt.Fprintf(f, "\n")

			for _, p := range classFields {
				fieldName := toCamelCase(p.Name)
				if fieldName == "Type" {
					fieldName = "TypeField"
				}

				typeName := p.Type
				isVector := false
				if strings.HasPrefix(typeName, "vector<") {
					typeName = strings.TrimSuffix(strings.TrimPrefix(typeName, "vector<"), ">")
					isVector = true
				}
				className := toCamelCase(typeName)

				if isVector {
					fmt.Fprintf(f, "\tif aux.%s != nil {\n", fieldName)
					fmt.Fprintf(f, "\t\tt.%s = make([]%s, len(aux.%s))\n", fieldName, className, fieldName)
					fmt.Fprintf(f, "\t\tfor i, raw := range aux.%s {\n", fieldName)
					fmt.Fprintf(f, "\t\t\tv, err := Unmarshal%s(raw)\n", className)
					fmt.Fprintf(f, "\t\t\tif err != nil { return err }\n")
					fmt.Fprintf(f, "\t\t\tt.%s[i] = v\n", fieldName)
					fmt.Fprintf(f, "\t\t}\n")
					fmt.Fprintf(f, "\t}\n")
				} else {
					fmt.Fprintf(f, "\tif aux.%s != nil {\n", fieldName)
					fmt.Fprintf(f, "\t\tv, err := Unmarshal%s(aux.%s)\n", className, fieldName)
					fmt.Fprintf(f, "\t\tif err != nil { return err }\n")
					fmt.Fprintf(f, "\t\tt.%s = v\n", fieldName)
					fmt.Fprintf(f, "\t}\n")
				}
			}
			fmt.Fprintf(f, "\treturn nil\n")
			fmt.Fprintf(f, "}\n\n")
		}
	}

	// Unmarshal helper
	fmt.Fprintln(f, "func Unmarshal(data []byte) (TlObject, string, error) {")
	fmt.Fprintln(f, "\tvar typeObj struct {")
	fmt.Fprintln(f, "\t\tType string `json:\"@type\"`")
	fmt.Fprintln(f, "\t\tExtra string `json:\"@extra\"`")
	fmt.Fprintln(f, "\t}")
	fmt.Fprintln(f, "\tif err := json.Unmarshal(data, &typeObj); err != nil { return nil, \"\", err }")
	fmt.Fprintln(f, "\tswitch typeObj.Type {")
	for _, t := range types {
		structName := toCamelCase(t.Name)
		fmt.Fprintf(f, "\tcase \"%s\":\n", t.Name)
		fmt.Fprintf(f, "\t\tvar obj %s\n", structName)
		fmt.Fprintf(f, "\t\tif err := json.Unmarshal(data, &obj); err != nil { return nil, \"\", err }\n")
		fmt.Fprintf(f, "\t\treturn &obj, typeObj.Extra, nil\n")
	}
	fmt.Fprintln(f, "\tdefault:")
	fmt.Fprintf(f, "\t\treturn nil, \"\", fmt.Errorf(\"unknown type: %%s\", typeObj.Type)\n")
	fmt.Fprintln(f, "\t}")
	fmt.Fprintln(f, "}")
}

func generateFunctions(functions []TLType, classes map[string]*TLClass) {
	f, err := os.Create("gen_functions.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package gotdbot")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import \"encoding/json\"")
	fmt.Fprintln(f)

	for _, t := range functions {
		structName := toCamelCase(t.Name)
		fmt.Fprintf(f, "// %s %s\n", structName, formatDesc(t.Description))
		fmt.Fprintf(f, "type %s struct {\n", structName)

		for _, p := range t.Params {
			goType := toGoType(p.Type, classes)
			fieldName := toCamelCase(p.Name)
			if fieldName == "Type" {
				fieldName = "TypeField"
			}
			jsonTag := fmt.Sprintf("`json:\"%s\"`", p.Name)
			if p.IsOptional {
				jsonTag = fmt.Sprintf("`json:\"%s,omitempty\"`", p.Name)
			}
			fmt.Fprintf(f, "\t// %s\n", formatDesc(p.Description))
			fmt.Fprintf(f, "\t%s %s %s\n", fieldName, goType, jsonTag)
		}
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) Type() string {\n", structName)
		fmt.Fprintf(f, "\treturn \"%s\"\n", t.Name)
		fmt.Fprintf(f, "}\n\n")

		// Functions do not implement class interfaces usually, they return objects.

		// MarshalJSON (only @type)
		fmt.Fprintf(f, "func (t *%s) MarshalJSON() ([]byte, error) {\n", structName)
		fmt.Fprintf(f, "\ttype Alias %s\n", structName)
		fmt.Fprintf(f, "\treturn json.Marshal(&struct {\n")
		fmt.Fprintf(f, "\t\tTypeStr string `json:\"@type\"`\n")
		fmt.Fprintf(f, "\t\t*Alias\n")
		fmt.Fprintf(f, "\t}{\n")
		fmt.Fprintf(f, "\t\tTypeStr: \"%s\",\n", t.Name)
		fmt.Fprintf(f, "\t\tAlias:   (*Alias)(t),\n")
		fmt.Fprintf(f, "\t})\n")
		fmt.Fprintf(f, "}\n\n")
	}
}

func generateOptions(functions []TLType, classes map[string]*TLClass) {
	f, err := os.Create("gen_options.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package gotdbot")
	fmt.Fprintln(f)

	for _, fn := range functions {
		methodName := toCamelCase(fn.Name)
		hasOptional := false
		for _, p := range fn.Params {
			if p.IsOptional {
				hasOptional = true
				break
			}
		}

		if hasOptional {
			optsStructName := methodName + "Opts"
			fmt.Fprintf(f, "// %s contains optional parameters for %s\n", optsStructName, methodName)
			fmt.Fprintf(f, "type %s struct {\n", optsStructName)
			for _, p := range fn.Params {
				if p.IsOptional {
					goType := toGoType(p.Type, classes)
					fieldName := toCamelCase(p.Name)
					if fieldName == "Type" {
						fieldName = "TypeField"
					}
					fmt.Fprintf(f, "\t%s %s\n", fieldName, goType)
				}
			}
			fmt.Fprintf(f, "}\n\n")
		}
	}
}

func toLowerCamelCase(s string) string {
	if s == "" {
		return ""
	}
	camel := toCamelCase(s)
	runes := []rune(camel)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

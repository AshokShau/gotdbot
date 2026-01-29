package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	fmt.Fprintln(f, "\tSetExtra(extra string)")
	fmt.Fprintln(f, "\tGetExtra() string")
	fmt.Fprintln(f, "}")
	fmt.Fprintln(f)

	sortedNames := sortKeysAZ(classes)
	for _, name := range sortedNames {
		cls := classes[name]
		sort.Strings(cls.Implementations)
		structName := toCamelCase(name)
		fmt.Fprintf(f, "// %s %s\n", structName, formatDesc(cls.Description))
		fmt.Fprintf(f, "type %s struct {\n", structName)
		fmt.Fprintf(f, "\tTypeStr string `json:\"@type\"`\n")
		// Add fields for implementations
		for _, impl := range cls.Implementations {
			implStructName := toCamelCase(impl)
			fmt.Fprintf(f, "\t%s *%s `json:\"%s,omitempty\"`\n", implStructName, implStructName, impl)
		}
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) Type() string {\n", structName)
		fmt.Fprintf(f, "\treturn t.TypeStr\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) SetExtra(extra string) {\n", structName)
		fmt.Fprintf(f, "\t// Extra not supported on abstract class directly, or no op\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) GetExtra() string {\n", structName)
		fmt.Fprintf(f, "\treturn \"\"\n")
		fmt.Fprintf(f, "}\n\n")

		// UnmarshalJSON
		fmt.Fprintf(f, "func (t *%s) UnmarshalJSON(b []byte) error {\n", structName)
		fmt.Fprintf(f, "\tvar typeObj struct { Type string `json:\"@type\"` }\n")
		fmt.Fprintf(f, "\tif err := json.Unmarshal(b, &typeObj); err != nil { return err }\n")
		fmt.Fprintf(f, "\tt.TypeStr = typeObj.Type\n")
		fmt.Fprintf(f, "\tswitch typeObj.Type {\n")
		for _, impl := range cls.Implementations {
			implStructName := toCamelCase(impl)
			fmt.Fprintf(f, "\tcase \"%s\":\n", impl)
			fmt.Fprintf(f, "\t\tt.%s = new(%s)\n", implStructName, implStructName)
			fmt.Fprintf(f, "\t\treturn json.Unmarshal(b, t.%s)\n", implStructName)
		}
		fmt.Fprintf(f, "\t}\n")
		fmt.Fprintf(f, "\treturn nil\n")
		fmt.Fprintf(f, "}\n\n")

		// MarshalJSON for Class
		fmt.Fprintf(f, "func (t *%s) MarshalJSON() ([]byte, error) {\n", structName)
		for _, impl := range cls.Implementations {
			implStructName := toCamelCase(impl)
			fmt.Fprintf(f, "\tif t.%s != nil {\n", implStructName)
			fmt.Fprintf(f, "\t\treturn json.Marshal(t.%s)\n", implStructName)
			fmt.Fprintf(f, "\t}\n")
		}
		fmt.Fprintf(f, "\treturn nil, fmt.Errorf(\"no value set for %s\")\n", structName)
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
	fmt.Fprintln(f)

	for _, t := range types {
		structName := toCamelCase(t.Name)
		fmt.Fprintf(f, "// %s %s\n", structName, formatDesc(t.Description))
		fmt.Fprintf(f, "type %s struct {\n", structName)
		fmt.Fprintf(f, "\tTypeStr string `json:\"@type\"`\n")
		fmt.Fprintf(f, "\tExtra   string `json:\"@extra,omitempty\"`\n")
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

		fmt.Fprintf(f, "func (t *%s) SetExtra(extra string) {\n", structName)
		fmt.Fprintf(f, "\tt.Extra = extra\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) GetExtra() string {\n", structName)
		fmt.Fprintf(f, "\treturn t.Extra\n")
		fmt.Fprintf(f, "}\n\n")

		// MarshalJSON for Concrete Type
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

	// Generate unmarshal helper
	fmt.Fprintln(f, "func Unmarshal(data []byte) (TlObject, error) {")
	fmt.Fprintln(f, "\tvar typeObj struct { Type string `json:\"@type\"` }")
	fmt.Fprintln(f, "\tif err := json.Unmarshal(data, &typeObj); err != nil { return nil, err }")
	fmt.Fprintln(f, "\tswitch typeObj.Type {")
	for _, t := range types {
		structName := toCamelCase(t.Name)
		fmt.Fprintf(f, "\tcase \"%s\":\n", t.Name)
		fmt.Fprintf(f, "\t\tvar obj %s\n", structName)
		fmt.Fprintf(f, "\t\tif err := json.Unmarshal(data, &obj); err != nil { return nil, err }\n")
		fmt.Fprintf(f, "\t\treturn &obj, nil\n")
	}
	fmt.Fprintln(f, "\tdefault:")
	fmt.Fprintln(f, "\t\treturn nil, nil // Unknown type")
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
		fmt.Fprintf(f, "\tTypeStr string `json:\"@type\"`\n")
		fmt.Fprintf(f, "\tExtra   string `json:\"@extra,omitempty\"`\n")
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

		fmt.Fprintf(f, "func (t *%s) SetExtra(extra string) {\n", structName)
		fmt.Fprintf(f, "\tt.Extra = extra\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) GetExtra() string {\n", structName)
		fmt.Fprintf(f, "\treturn t.Extra\n")
		fmt.Fprintf(f, "}\n\n")

		// MarshalJSON for Function
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

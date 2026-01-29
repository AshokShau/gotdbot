package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"unicode"
)

type TLType struct {
	Name        string
	Description string
	Params      []TLParam
	ResultType  string
	IsFunction  bool
}

type TLParam struct {
	Name        string
	Type        string
	Description string
	IsOptional  bool
}

type TLClass struct {
	Name            string
	Description     string
	Implementations []string
}

var (
	classRegex       = regexp.MustCompile(`^//@class\s+(?P<name>\w+)\s+@description\s+(?P<description>.*)`)
	paramRegex       = regexp.MustCompile(`^//@(?P<name>\w+)\s+(?P<description>.*)`)
	tlDefRegex       = regexp.MustCompile(`^(?P<name>\w+)\s+(?P<params>.*)=\s+(?P<type>\w+);$`)
	paramDetailRegex = regexp.MustCompile(`(?P<name>\w+):(?P<type>[\w<>]+)`)
)

func main() {
	file, err := os.Open("td_api.tl")
	if err != nil {
		log.Fatalf("Failed to open td_api.tl: %v", err)
	}
	defer file.Close()

	var types []TLType
	var functions []TLType
	var classes = make(map[string]*TLClass)

	var currentDescription string
	var currentParams = make(map[string]string)
	var isFunctionsSection bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if line == "---functions---" {
			isFunctionsSection = true
			continue
		}

		if strings.HasPrefix(line, "//") {
			if strings.HasPrefix(line, "//@class") {
				matches := classRegex.FindStringSubmatch(line)
				if len(matches) > 0 {
					classes[matches[1]] = &TLClass{
						Name:        matches[1],
						Description: matches[2],
					}
				}
			} else if strings.HasPrefix(line, "//@description") {
				currentDescription = strings.TrimPrefix(line, "//@description ")
			} else if strings.HasPrefix(line, "//@") {
				matches := paramRegex.FindStringSubmatch(line)
				if len(matches) > 0 {
					currentParams[matches[1]] = matches[2]
				}
			}
			continue
		}

		// Parse TL definition
		matches := tlDefRegex.FindStringSubmatch(line)
		if len(matches) > 0 {
			name := matches[1]
			paramsStr := matches[2]
			resultType := matches[3]

			var params []TLParam
			paramMatches := paramDetailRegex.FindAllStringSubmatch(paramsStr, -1)
			for _, pm := range paramMatches {
				pName := pm[1]
				pType := pm[2]
				pDesc := currentParams[pName]
				isOptional := strings.Contains(pDesc, "may be null") || strings.Contains(pDesc, "pass null")

				params = append(params, TLParam{
					Name:        pName,
					Type:        pType,
					Description: pDesc,
					IsOptional:  isOptional,
				})
			}

			tlType := TLType{
				Name:        name,
				Description: currentDescription,
				Params:      params,
				ResultType:  resultType,
				IsFunction:  isFunctionsSection,
			}

			if isFunctionsSection {
				functions = append(functions, tlType)
			} else {
				types = append(types, tlType)
				// Add to class implementations
				if cls, ok := classes[resultType]; ok {
					cls.Implementations = append(cls.Implementations, name)
				}
			}

			// Reset current documentation
			currentDescription = ""
			currentParams = make(map[string]string)
		}
	}

	if err := os.MkdirAll("types", 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("client", 0755); err != nil {
		log.Fatal(err)
	}

	generateClasses(classes)
	generateObjects(types, classes)
	generateFunctions(functions, classes)
	generateOptions(functions, classes)
	generateMethods(functions, classes)
	generateUpdates(types)

	gofmt("types", "client")
}

func toGoType(tlType string, classes map[string]*TLClass) string {
	if strings.HasPrefix(tlType, "vector<") {
		inner := strings.TrimSuffix(strings.TrimPrefix(tlType, "vector<"), ">")
		return "[]" + toGoType(inner, classes)
	}
	switch tlType {
	case "int32":
		return "int32"
	case "int53":
		return "int64"
	case "int64":
		return "string"
	case "double":
		return "float64"
	case "string":
		return "string"
	case "bytes":
		return "string"
	case "Bool":
		return "bool"
	default:
		if _, ok := classes[tlType]; ok {
			return "*" + toCamelCase(tlType)
		}
		return "*" + toCamelCase(tlType)
	}
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	var sb strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		sb.WriteString(string(runes))
	}
	return sb.String()
}

func formatDesc(desc string) string {
	return strings.ReplaceAll(desc, "\n", " ")
}

func generateClasses(classes map[string]*TLClass) {
	f, err := os.Create("types/classes.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package types")
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

	for name, cls := range classes {
		structName := toCamelCase(name)
		fmt.Fprintf(f, "// %s %s\n", structName, formatDesc(cls.Description))
		fmt.Fprintf(f, "type %s struct {\n", structName)
		// Removed typeStr field
		// Add fields for implementations
		for _, impl := range cls.Implementations {
			implStructName := toCamelCase(impl)
			fmt.Fprintf(f, "\t%s *%s `json:\"%s,omitempty\"`\n", implStructName, implStructName, impl)
		}
		fmt.Fprintf(f, "}\n\n")

		// Type() returns empty for abstract class
		fmt.Fprintf(f, "func (t *%s) Type() string {\n", structName)
		fmt.Fprintf(f, "\treturn \"\"\n")
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
		// no assignment to t.typeStr since field removed
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
	f, err := os.Create("types/objects.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package types")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import \"encoding/json\"")
	fmt.Fprintln(f)

	for _, t := range types {
		structName := toCamelCase(t.Name)
		fmt.Fprintf(f, "// %s %s\n", structName, formatDesc(t.Description))
		fmt.Fprintf(f, "type %s struct {\n", structName)
		// Removed extra field
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
		fmt.Fprintf(f, "\t// Extra removed; no-op\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) GetExtra() string {\n", structName)
		fmt.Fprintf(f, "\treturn \"\"\n")
		fmt.Fprintf(f, "}\n\n")

		// MarshalJSON for Concrete Type (no @extra)
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

		// UnmarshalJSON (ignore @extra)
		fmt.Fprintf(f, "func (t *%s) UnmarshalJSON(b []byte) error {\n", structName)
		fmt.Fprintf(f, "\ttype Alias %s\n", structName)
		fmt.Fprintf(f, "\tvar aux struct {\n")
		fmt.Fprintf(f, "\t\t*Alias\n")
		fmt.Fprintf(f, "\t}\n")
		fmt.Fprintf(f, "\taux.Alias = (*Alias)(t)\n")
		fmt.Fprintf(f, "\tif err := json.Unmarshal(b, &aux); err != nil { return err }\n")
		fmt.Fprintf(f, "\treturn nil\n")
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
	f, err := os.Create("types/functions.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package types")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import \"encoding/json\"")
	fmt.Fprintln(f)

	for _, t := range functions {
		structName := toCamelCase(t.Name)
		fmt.Fprintf(f, "// %s %s\n", structName, formatDesc(t.Description))
		fmt.Fprintf(f, "type %s struct {\n", structName)
		// Removed extra field
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
		fmt.Fprintf(f, "\t// Extra removed; no-op\n")
		fmt.Fprintf(f, "}\n\n")

		fmt.Fprintf(f, "func (t *%s) GetExtra() string {\n", structName)
		fmt.Fprintf(f, "\treturn \"\"\n")
		fmt.Fprintf(f, "}\n\n")

		// MarshalJSON for Function (no @extra)
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

		// UnmarshalJSON to ignore @extra
		fmt.Fprintf(f, "func (t *%s) UnmarshalJSON(b []byte) error {\n", structName)
		fmt.Fprintf(f, "\ttype Alias %s\n", structName)
		fmt.Fprintf(f, "\tvar aux struct {\n")
		fmt.Fprintf(f, "\t\t*Alias\n")
		fmt.Fprintf(f, "\t}\n")
		fmt.Fprintf(f, "\taux.Alias = (*Alias)(t)\n")
		fmt.Fprintf(f, "\tif err := json.Unmarshal(b, &aux); err != nil { return err }\n")
		fmt.Fprintf(f, "\treturn nil\n")
		fmt.Fprintf(f, "}\n\n")
	}
}

func generateOptions(functions []TLType, classes map[string]*TLClass) {
	f, err := os.Create("types/options.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package types")
	fmt.Fprintln(f)

	for _, fn := range functions {
		methodName := toCamelCase(fn.Name)

		// Check for optional parameters
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

func generateMethods(functions []TLType, classes map[string]*TLClass) {
	f, err := os.Create("client/methods.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package client")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import \"github.com/AshokShau/gotdbot/types\"")
	fmt.Fprintln(f)

	for _, fn := range functions {
		methodName := toCamelCase(fn.Name)
		structName := toCamelCase(fn.Name)

		// Check for optional parameters
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

		retTypeStr := "*types." + resultType
		if _, ok := classes[fn.ResultType]; ok {
			retTypeStr = "*types." + toCamelCase(fn.ResultType)
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
			// Fix import path
			if strings.HasPrefix(goType, "*") {
				goType = "*types." + strings.TrimPrefix(goType, "*")
			} else if strings.HasPrefix(goType, "[]*") {
				goType = "[]*types." + strings.TrimPrefix(goType, "[]*")
			}

			fieldName := toCamelCase(p.Name)
			argName := strings.ToLower(fieldName[:1]) + fieldName[1:]
			if argName == "type" {
				argName = "typeField"
			}
			if argName == "func" {
				argName = "funcArg"
			}

			// Avoid collision with imported package `types`
			if argName == "types" {
				argName = "typesTd"
			}

			args = append(args, fmt.Sprintf("%s %s", argName, goType))
		}

		if hasOptional {
			args = append(args, fmt.Sprintf("opts *types.%s", optsStructName))
		}

		fmt.Fprintf(f, "%s", strings.Join(args, ", "))
		fmt.Fprintf(f, ") (%s, error) {\n", retTypeStr)

		fmt.Fprintf(f, "\treq := &types.%s{\n", structName)
		// No need to set TypeStr anymore
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
			if argName == "types" {
				argName = "typesTd"
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

func generateUpdates(types []TLType) {
	f, err := os.Create("client/updates.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package client")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import \"github.com/AshokShau/gotdbot/types\"")
	fmt.Fprintln(f)

	for _, t := range types {
		if t.ResultType != "Update" {
			continue
		}

		structName := toCamelCase(t.Name)
		methodName := "On" + structName

		fmt.Fprintf(f, "// %s %s\n", methodName, formatDesc(t.Description))
		fmt.Fprintf(f, "func (c *Client) %s(fn func(client *Client, update *types.%s) error, filter FilterFunc, position int) {\n", methodName, structName)
		fmt.Fprintf(f, "\tc.AddHandler(\"%s\", func(cl *Client, u types.TlObject) error {\n", t.Name)
		fmt.Fprintf(f, "\t\tif up, ok := u.(*types.%s); ok {\n", structName)
		fmt.Fprintf(f, "\t\t\treturn fn(cl, up)\n")
		fmt.Fprintf(f, "\t\t}\n")
		fmt.Fprintf(f, "\t\treturn nil\n")
		fmt.Fprintf(f, "\t}, filter, position)\n")
		fmt.Fprintf(f, "}\n\n")
	}
}

func gofmt(paths ...string) {
	args := append([]string{"-w"}, paths...)
	cmd := exec.Command("gofmt", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("gofmt failed: %v", err)
	}
}

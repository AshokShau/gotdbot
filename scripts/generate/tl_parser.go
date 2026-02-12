package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	classRegex       = regexp.MustCompile(`^//@class\s+(?P<name>\w+)\s+@description\s+(?P<description>.*)`)
	paramRegex       = regexp.MustCompile(`^//@(?P<name>\w+)\s+(?P<description>.*)`)
	tlDefRegex       = regexp.MustCompile(`^(?P<name>\w+)\s+(?P<params>.*)=\s+(?P<type>\w+);$`)
	paramDetailRegex = regexp.MustCompile(`(?P<name>\w+):(?P<type>[\w<>]+)`)
)

// FetchAndParseTL fetches the TL schema from the given URL and parses it into TDLibJSON structure.
func FetchAndParseTL(url string) (*TDLibJSON, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to GET %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch %s: status %d", url, resp.StatusCode)
	}

	return parseTLFromReader(resp.Body)
}

func parseTLFromReader(r io.Reader) (*TDLibJSON, error) {
	data := &TDLibJSON{
		Name:      "Auto-generated JSON TDLib API",
		Classes:   make(map[string]*ClassDef),
		Types:     make(map[string]*TypeDef),
		Updates:   make(map[string]*TypeDef),
		Functions: make(map[string]*TypeDef),
	}

	var currentDescription string
	var currentParams = make(map[string]string)
	var isFunctionsSection bool
	var start bool

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "---functions---") {
			isFunctionsSection = true
			continue
		}

		if strings.HasPrefix(line, "//") {
			start = true
			if strings.HasPrefix(line, "//@class") {
				matches := classRegex.FindStringSubmatch(line)
				if len(matches) > 0 {
					name := matches[1]
					desc := matches[2]
					data.Classes[name] = &ClassDef{
						Description: desc,
						Types:       []string{},
						Functions:   []string{},
					}
				}
			} else if strings.HasPrefix(line, "//@description") {
				currentDescription = strings.TrimPrefix(line, "//@description ")
			} else if strings.HasPrefix(line, "//@") {
				matches := paramRegex.FindStringSubmatch(line)
				if len(matches) > 0 {
					name := matches[1]
					desc := matches[2]
					cleanName := strings.TrimPrefix(name, "param_")
					currentParams[cleanName] = desc
				}
			}
			continue
		}

		if line != "" && start {
			// Parse TL definition
			matches := tlDefRegex.FindStringSubmatch(line)
			if len(matches) > 0 {
				name := matches[1]
				paramsStr := matches[2]
				resultType := matches[3]

				typeDef := &TypeDef{
					Description: currentDescription,
					Args:        make(map[string]*ArgDef),
					Type:        resultType,
				}

				paramMatches := paramDetailRegex.FindAllStringSubmatch(paramsStr, -1)
				for _, pm := range paramMatches {
					pName := pm[1]
					pType := pm[2]
					pDesc := currentParams[pName]
					isOptional := strings.Contains(pDesc, "may be null") ||
						strings.Contains(pDesc, "pass null") ||
						strings.Contains(pDesc, "may be empty") ||
						strings.Contains(pDesc, "If non-empty,")

					typeDef.Args[pName] = &ArgDef{
						Description: pDesc,
						IsOptional:  isOptional,
						Type:        pType,
					}
				}

				if isFunctionsSection {
					data.Functions[name] = typeDef
					if cls, ok := data.Classes[resultType]; ok {
						cls.Functions = append(cls.Functions, name)
					}
				} else if strings.HasPrefix(name, "update") {
					data.Updates[name] = typeDef
					if cls, ok := data.Classes[resultType]; ok {
						cls.Types = append(cls.Types, name)
					}
				} else {
					data.Types[name] = typeDef
					if cls, ok := data.Classes[resultType]; ok {
						cls.Types = append(cls.Types, name)
					}
				}

				// Reset
				currentDescription = ""
				currentParams = make(map[string]string)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %v", err)
	}

	return data, nil
}

func SaveTDLibJSON(data *TDLibJSON, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(data)
}

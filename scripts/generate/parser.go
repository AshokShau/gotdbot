package main

import (
	"bufio"
	"fmt"
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

func ParseTL(filename string) (types []TLType, functions []TLType, classes map[string]*TLClass, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to open %s: %v", filename, err)
	}
	defer file.Close()

	classes = make(map[string]*TLClass)
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

	return types, functions, classes, nil
}

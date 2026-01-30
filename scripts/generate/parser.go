package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var (
	classRegex       = regexp.MustCompile(`^//@class\s+(?P<name>\w+)\s+@description\s+(?P<description>.*)`)
	paramRegex       = regexp.MustCompile(`^//@(?P<name>\w+)\s+(?P<description>.*)`)
	tlDefRegex       = regexp.MustCompile(`^(?P<name>\w+)\s+(?P<params>.*)=\s+(?P<type>\w+);$`)
	paramDetailRegex = regexp.MustCompile(`(?P<name>\w+):(?P<type>[\w<>]+)`)
)

// ParseTL fetches a TL file over HTTP and parses.
func ParseTL(url string) (types []TLType, functions []TLType, classes map[string]*TLClass, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to GET %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, nil, fmt.Errorf("failed to fetch %s: status %d", url, resp.StatusCode)
	}

	return parseFromReader(resp.Body)
}

// parseFromReader contains the core parsing logic and works with any io.Reader.
func parseFromReader(r io.Reader) (types []TLType, functions []TLType, classes map[string]*TLClass, err error) {
	classes = make(map[string]*TLClass)
	var currentDescription string
	var currentParams = make(map[string]string)
	var isFunctionsSection bool

	scanner := bufio.NewScanner(r)
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

	if serr := scanner.Err(); serr != nil {
		return nil, nil, nil, fmt.Errorf("scanner error: %v", serr)
	}

	return types, functions, classes, nil
}

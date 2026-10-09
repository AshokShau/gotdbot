package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// LoadTDLibJSON reads the JSON file and returns the TDLibJSON structure.
func LoadTDLibJSON(source string) (*TDLibJSON, error) {
	var body io.ReadCloser
	if strings.HasPrefix(source, "https://") || strings.HasPrefix(source, "http://") {
		client := &http.Client{Timeout: 30 * time.Second}
		resp, err := client.Get(source)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to fetch json: %s", resp.Status)
		}
		body = resp.Body
	} else {
		f, err := os.Open(source)
		if err != nil {
			return nil, err
		}
		body = f
	}
	defer body.Close()

	var data TDLibJSON
	decoder := json.NewDecoder(io.LimitReader(body, 32<<20))
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	if data.Version == "" || len(data.Functions) == 0 || len(data.Types) == 0 {
		return nil, fmt.Errorf("invalid TDLib schema")
	}

	return &data, nil
}

// ConvertJSONToGen converts the TDLibJSON structure into the format expected by the generators.
func ConvertJSONToGen(data *TDLibJSON) ([]TLType, []TLType, map[string]*TLClass, map[string]*OptionDef) {
	var types []TLType
	var functions []TLType
	classes := make(map[string]*TLClass)

	// Convert Classes
	for name, def := range data.Classes {
		classes[name] = &TLClass{
			Name:            name,
			Description:     def.Description,
			Implementations: def.Types,
		}
	}

	// Helper to convert TypeDef to TLType
	convert := func(name string, def *TypeDef) TLType {
		var params []TLParam
		var argNames []string
		for k := range def.Args {
			argNames = append(argNames, k)
		}
		sort.Strings(argNames)

		for _, k := range argNames {
			arg := def.Args[k]
			params = append(params, TLParam{
				Name:        k,
				Type:        arg.Type,
				Description: arg.Description,
				IsOptional:  arg.IsOptional,
			})
		}

		return TLType{
			Name:        name,
			Description: def.Description,
			Params:      params,
			ResultType:  def.Type,
		}
	}

	// Convert Types
	for name, def := range data.Types {
		types = append(types, convert(name, def))
	}

	// Convert Updates
	for name, def := range data.Updates {
		types = append(types, convert(name, def))
	}

	// Convert Functions
	for name, def := range data.Functions {
		functions = append(functions, convert(name, def))
	}

	return types, functions, classes, data.Options
}

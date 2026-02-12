package main

import (
	"encoding/json"
	"os"
	"sort"
)

// LoadTDLibJSON reads the JSON file and returns the TDLibJSON structure.
func LoadTDLibJSON(filename string) (*TDLibJSON, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data TDLibJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

// ConvertJSONToGen converts the TDLibJSON structure into the format expected by the generators.
func ConvertJSONToGen(data *TDLibJSON) ([]TLType, []TLType, map[string]*TLClass) {
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
	convert := func(name string, def *TypeDef, isFunction bool) TLType {
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
			IsFunction:  isFunction,
		}
	}

	// Convert Types
	for name, def := range data.Types {
		types = append(types, convert(name, def, false))
	}

	// Convert Updates
	for name, def := range data.Updates {
		types = append(types, convert(name, def, false))
	}

	// Convert Functions
	for name, def := range data.Functions {
		functions = append(functions, convert(name, def, true))
	}

	return types, functions, classes
}

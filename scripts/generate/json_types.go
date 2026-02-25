package main

type TDLibJSON struct {
	Name      string               `json:"name"`
	Version   string               `json:"version"`
	Commit    string               `json:"commit"`
	Classes   map[string]*ClassDef `json:"classes"`
	Types     map[string]*TypeDef  `json:"types"`
	Updates   map[string]*TypeDef  `json:"updates"`
	Functions map[string]*TypeDef  `json:"functions"`
}

type ClassDef struct {
	Description string   `json:"description"`
	Types       []string `json:"types"`
	Functions   []string `json:"functions"`
}

type TypeDef struct {
	Description string             `json:"description"`
	Args        map[string]*ArgDef `json:"args"`
	Type        string             `json:"type"`
}

type ArgDef struct {
	Description string `json:"description"`
	IsOptional  bool   `json:"is_optional"`
	Type        string `json:"type"`
}

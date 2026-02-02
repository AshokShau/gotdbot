package main

import (
	"log"
	"sort"
)

func main() {
	types, functions, classes, err := ParseTL("https://raw.githubusercontent.com/tdlib/td/refs/heads/master/td/generate/scheme/td_api.tl")
	if err != nil {
		log.Fatalf("Failed to parse TL: %v", err)
	}

	sortTLTypesByName(types)
	sortTLTypesByName(functions)
	for _, cls := range classes {
		sort.Strings(cls.Implementations)
	}

	generateClasses(classes)
	generateObjects(types, classes)
	generateFunctions(functions, classes)
	generateOptions(functions, classes)
	generateMethods(functions, classes)
	// generateUpdates(types)
	extFiles := generateExtHandlers(types)
	helperFiles := generateHelpers(types, functions, classes)
	filesToFmt := append([]string{"gen_classes.go", "gen_types.go", "gen_functions.go", "gen_options.go", "gen_methods.go"}, helperFiles...)
	filesToFmt = append(filesToFmt, extFiles...)
	gofmt(filesToFmt...)
}

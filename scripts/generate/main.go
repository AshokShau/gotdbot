package main

import (
	"log"
	"sort"
)

func main() {
	types, functions, classes, err := ParseTL("td_api.tl")
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
	generateUpdates(types)
	//helperFiles := generateHelpers(types, functions, classes)
	//helperFiles...
	filesToFmt := append([]string{"gen_classes.go", "gen_types.go", "gen_functions.go", "gen_options.go", "gen_methods.go", "gen_updates.go"})
	gofmt(filesToFmt...)
}

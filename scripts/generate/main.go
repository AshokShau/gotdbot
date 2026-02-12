package main

import (
	"flag"
	"log"
	"os"
	"sort"
)

func main() {
	download := flag.Bool("download", false, "Download TL schema and generate tdlib.json")
	jsonPath := flag.String("json", "scripts/generate/tdlib.json", "Path to tdlib.json")
	tlUrl := flag.String("url", "https://raw.githubusercontent.com/tdlib/td/refs/heads/master/td/generate/scheme/td_api.tl", "URL to fetch TL schema from")
	flag.Parse()

	if *download {
		log.Printf("Fetching TL schema from %s...", *tlUrl)
		data, err := FetchAndParseTL(*tlUrl)
		if err != nil {
			log.Fatalf("Failed to fetch/parse TL: %v", err)
		}

		log.Printf("Saving JSON to %s...", *jsonPath)
		if err := SaveTDLibJSON(data, *jsonPath); err != nil {
			log.Fatalf("Failed to save JSON: %v", err)
		}
		log.Println("Done.")
		return
	}

	log.Printf("Loading JSON from %s...", *jsonPath)
	if _, err := os.Stat(*jsonPath); os.IsNotExist(err) {
		log.Fatalf("JSON file not found at %s. Run with --download to generate it.", *jsonPath)
	}

	data, err := LoadTDLibJSON(*jsonPath)
	if err != nil {
		log.Fatalf("Failed to load JSON: %v", err)
	}

	types, functions, classes := ConvertJSONToGen(data)

	sortTLTypesByName(types)
	sortTLTypesByName(functions)
	for _, cls := range classes {
		sort.Strings(cls.Implementations)
	}

	log.Println("Generating Go code...")
	generateClasses(classes)
	generateObjects(types, classes)
	generateFunctions(functions, classes)
	generateOptions(functions, classes)
	generateMethods(functions, classes)
	extFiles := generateExtHandlers(types)
	helperFiles := generateHelpers(types, functions, classes)
	filesToFmt := append([]string{"gen_classes.go", "gen_types.go", "gen_functions.go", "gen_options.go", "gen_methods.go"}, helperFiles...)
	filesToFmt = append(filesToFmt, extFiles...)
	gofmt(filesToFmt...)
	log.Println("Done.")
}

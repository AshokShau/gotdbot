package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var tdlibVersionPattern = regexp.MustCompile(`^v?\d+\.\d+\.\d+$`)
var tdlibConstPattern = regexp.MustCompile(`TDLibVersion\s*=\s*"v\d+\.\d+\.\d+"`)
var downloaderVersionPattern = regexp.MustCompile(`version\s*=\s*"v\d+\.\d+\.\d+"`)
var tdlibLibraryPattern = regexp.MustCompile(`libtdjson\.so\.\d+\.\d+\.\d+`)
var markdownVersionPatterns = []*regexp.Regexp{
	regexp.MustCompile("(\\bTDLib[ \\t]+[`*]*v?)\\d+\\.\\d+\\.\\d+"),
	regexp.MustCompile(`(\bdownloader[^\n]*\btargets[ \t]+v?)\d+\.\d+\.\d+`),
}

func updateTDLibVersionReferences(root, schemaVersion string) error {
	if !tdlibVersionPattern.MatchString(schemaVersion) {
		return fmt.Errorf("invalid TDLib schema version %q", schemaVersion)
	}
	version := strings.TrimPrefix(schemaVersion, "v")
	return filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		markdown := strings.HasSuffix(relativePath, ".md")
		example := strings.HasPrefix(relativePath, "examples"+string(os.PathSeparator)) && strings.HasSuffix(relativePath, ".go")
		if !markdown && !example && relativePath != "consts.go" && relativePath != "scripts/tools/get_tdjson.go" && relativePath != "internal/tdjson/tdjson_test.go" {
			return nil
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		updated := body
		switch {
		case relativePath == "consts.go":
			updated = tdlibConstPattern.ReplaceAllLiteral(body, []byte(`TDLibVersion = "v`+version+`"`))
		case relativePath == "scripts/tools/get_tdjson.go":
			updated = downloaderVersionPattern.ReplaceAllLiteral(body, []byte(`version = "v`+version+`"`))
		case markdown:
			for _, pattern := range markdownVersionPatterns {
				updated = pattern.ReplaceAll(updated, []byte("${1}"+version))
			}
		}
		updated = tdlibLibraryPattern.ReplaceAllLiteral(updated, []byte("libtdjson.so."+version))
		if bytes.Equal(updated, body) {
			return nil
		}
		return os.WriteFile(path, updated, 0644)
	})
}

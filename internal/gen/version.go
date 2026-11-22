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

var tdlibVersionReferences = []struct {
	path        string
	pattern     *regexp.Regexp
	replacement func(string) string
}{
	{"consts.go", regexp.MustCompile(`TDLibVersion\s*=\s*"v\d+\.\d+\.\d+"`), func(v string) string { return `TDLibVersion = "v` + v + `"` }},
	{"scripts/tools/get_tdjson.go", regexp.MustCompile(`version\s*=\s*"v\d+\.\d+\.\d+"`), func(v string) string { return `version = "v` + v + `"` }},
	{"internal/tdjson/tdjson_test.go", regexp.MustCompile(`libtdjson\.so\.\d+\.\d+\.\d+`), func(v string) string { return `libtdjson.so.` + v }},
}

var exampleTDLibLibraryPattern = regexp.MustCompile(`libtdjson\.so\.\d+\.\d+\.\d+`)

func updateTDLibVersionReferences(root, schemaVersion string) error {
	if !tdlibVersionPattern.MatchString(schemaVersion) {
		return fmt.Errorf("invalid TDLib schema version %q", schemaVersion)
	}
	version := strings.TrimPrefix(schemaVersion, "v")
	type update struct {
		path string
		body []byte
	}
	var updates []update
	queueUpdate := func(relativePath string, pattern *regexp.Regexp, replacement string, required bool) error {
		path := filepath.Join(root, relativePath)
		body, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", relativePath, err)
		}
		if !pattern.Match(body) {
			if required {
				return fmt.Errorf("TDLib version reference missing from %s", relativePath)
			}
			return nil
		}
		updated := pattern.ReplaceAllLiteral(body, []byte(replacement))
		if bytes.Equal(updated, body) {
			return nil
		}
		updates = append(updates, update{path: path, body: updated})
		return nil
	}
	for _, ref := range tdlibVersionReferences {
		if err := queueUpdate(ref.path, ref.pattern, ref.replacement(version), true); err != nil {
			return err
		}
	}
	if err := filepath.WalkDir(filepath.Join(root, "examples"), func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || entry.Type()&os.ModeSymlink != 0 || !strings.HasSuffix(entry.Name(), ".go") {
			return nil
		}
		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		return queueUpdate(relativePath, exampleTDLibLibraryPattern, "libtdjson.so."+version, false)
	}); err != nil {
		return fmt.Errorf("scan examples: %w", err)
	}
	for _, item := range updates {
		if err := os.WriteFile(item.path, item.body, 0644); err != nil {
			return fmt.Errorf("write %s: %w", item.path, err)
		}
	}
	return nil
}

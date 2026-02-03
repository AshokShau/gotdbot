package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const downloadURL = "https://github.com/AshokShau/gotdbot/releases/latest/download/TDLib-tdjson-linux-x86_64.tar.gz"

func main() {
	fmt.Println("Downloading tdjson release...")
	resp, err := http.Get(downloadURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("failed to download: " + resp.Status)
	}

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		panic(err)
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	found := false

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		name := filepath.Base(hdr.Name)
		if strings.HasPrefix(name, "libtdjson.so") {
			out, err := os.Create(name)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(out, tr)
			out.Close()
			if err != nil {
				panic(err)
			}

			fmt.Println("Extracted:", name)
			found = true
			break
		}
	}

	if !found {
		panic("libtdjson.so not found in archive")
	}
	fmt.Println("Done.")
}

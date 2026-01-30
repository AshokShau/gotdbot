package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	downloadURL = "https://github.com/AshokShau/gotdbot/releases/latest/download/tdjson-linux-x86_64.tar.gz"
	targetLib   = "libtdjson.so.1.8.60"
)

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

		if name == targetLib {
			out, err := os.Create(targetLib)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(out, tr)
			out.Close()
			if err != nil {
				panic(err)
			}

			fmt.Println("Extracted:", targetLib)
			found = true
			break
		}
	}

	if !found {
		panic("library not found in archive: " + targetLib)
	}
}

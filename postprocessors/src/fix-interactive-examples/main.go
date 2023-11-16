package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var TARGET_DIR = "book/interactive-examples"

func getHtmlFilesInDir(dir string) []string {
	var html_files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if path == dir {
			return nil
		}
		if info.IsDir() {
			html_files = append(html_files, getHtmlFilesInDir(path)...)
		} else if strings.HasSuffix(path, ".html") {
			html_files = append(html_files, path)
		}
		return nil
	})
	return html_files
}

func fixHtmlFile(htmlFilePath string) {
	defer wg.Done()
	// log.Println("Fixing " + htmlFilePath)
	html, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(html), "src=&#34;/media/") {
		html = []byte(strings.ReplaceAll(string(html), "src=&#34;/media/", "src=&#34;/interactive-examples/media/"))
	}
	err = os.WriteFile(htmlFilePath, html, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

var wg sync.WaitGroup

func main() {
	// log.Println("Fixing interactive examples")
	htmlFilePathsInDir := getHtmlFilesInDir(TARGET_DIR)
	log.Println("Found " + fmt.Sprint(len(htmlFilePathsInDir)) + " html files in " + TARGET_DIR)
	for _, htmlFilePath := range htmlFilePathsInDir {
		wg.Add(1)
		go fixHtmlFile((htmlFilePath))
	}
	wg.Wait()
}

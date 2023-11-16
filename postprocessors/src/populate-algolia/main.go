package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var NO_GO_TAGS []string = []string{"#Формальний-синтаксис", "#Специфікації", "script", "style", "noscript", "table", "pre", "ul", "ol", "h1", "blockquote"}

var CUT_MARKERS []string = []string{"#Дивіться-також", "#Автори-статті", "#Автори-перекладу"}

// Removes all html tags from the given html string
func extractTextFromHtml(htmlCode []byte) (string, string, error) {
	reader := strings.NewReader(string(htmlCode))
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", "", err
	}
	main := document.Find("main")
	h1 := main.Find("h1")
	title := h1.Text()
	for _, noGoTag := range NO_GO_TAGS {
		main.Find(noGoTag).Each(func(i int, selection *goquery.Selection) {
			selection.Remove()
		})
	}
	// Remove everything after CUT_MARKERS
	for _, cutMarker := range CUT_MARKERS {
		main.Find(cutMarker).Each(func(i int, selection *goquery.Selection) {
			selection.NextAll().Remove()
			selection.Remove()
		})
	}
	text := main.Text()
	textWithoutRedundantSpaces := strings.Join(strings.Fields(text), " ")
	return title, textWithoutRedundantSpaces, nil
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

var hashes map[string]uint32

func initHashes() {
	jsonData, err := os.ReadFile("hashes.json")
	if err != nil {
		if os.IsNotExist(err) {
			hashes = make(map[string]uint32)
			return
		}
		panic(err)
	}
	err = json.Unmarshal(jsonData, &hashes)
	if err != nil {
		panic(err)
	}
}

func main() {
	initHashes()
	algolia := NewAlgolia()
	err := filepath.Walk("book/uk/docs", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() != "index.html" {
			return nil
		}
		if path == "book/uk/docs/index.html" {
			return nil
		}
		htmlCode, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		title, text, err := extractTextFromHtml(htmlCode)
		if err != nil {
			return err
		}
		slug := path[strings.Index(path, "book/uk/docs/")+len("book/uk/docs/") : strings.LastIndex(path, "/index.html")]

		oldHash, ok := hashes[slug]
		newHash := hash(text)
		if ok && oldHash == hash(text) {
			return nil
		}
		fmt.Println("Saving to Algolia: " + slug)
		result, err := algolia.Index.SaveObject(Record{ObjectID: slug, Slug: slug, Text: text, Title: title})
		if err != nil {
			os.WriteFile("error.txt", []byte(text), 0644)
			if strings.Contains(err.Error(), "Record is too big") {
				result, err = algolia.Index.SaveObject(Record{ObjectID: slug, Slug: slug, Text: string(([]byte(text))[:9000]), Title: title})
			} else {
				os.WriteFile("error.txt", []byte(text), 0644)
				return err
			}
		}
		if err != nil {
			os.WriteFile("error.txt", []byte(text), 0644)
			return err
		}
		result.Wait()
		err = saveHash(slug, newHash)
		if err != nil {
			return err
		}
		// panic("stop")
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func saveHash(slug string, hashValue uint32) error {
	// textHash := hash(text)
	hashes[slug] = hashValue
	json, err := json.Marshal(hashes)
	if err != nil {
		return err
	}
	return os.WriteFile("hashes.json", json, 0644)
}

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	err := filepath.Walk("book/uk/docs", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() != "index.html" {
			return nil
		}
		log.Println(path)
		htmlCode, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		page, err := goquery.NewDocumentFromReader(strings.NewReader(string(htmlCode)))
		if err != nil {
			return err
		}
		err = fixDefinitions(page)
		if err != nil {
			return err
		}
		// if strings.Contains(path, "nth-child") {
		// 	return errors.New("stop")
		// }
		err = fixCyrillicIds(page)
		if err != nil {
			return err
		}
		err = fixCodeLinks(page)
		if err != nil {
			return err
		}
		err = fixHljs(page)
		if err != nil {
			return err
		}
		// Write page back into file
		html, err := page.Html()
		if err != nil {
			return err
		}
		err = os.WriteFile(path, []byte(html), 0644)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

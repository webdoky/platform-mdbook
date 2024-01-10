package main

import (
	"html"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fixCodeLinks(doc *goquery.Document) error {
	doc.Find("a > code").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		// log.Printf("text: %s", text)
		if strings.Contains(text, "&lt;") {
			selection.SetHtml(text)
			// log.Printf("text after: %s", selection.Text())
			title, titleExists := selection.Attr("title")
			if titleExists {
				selection.SetAttr("title", html.UnescapeString(title))
			}
		}
	})
	return nil
}

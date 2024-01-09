package main

import (
	"strings"
	"webdoky3/revamp/helpers"

	"github.com/PuerkitoBio/goquery"
)

func fixCyrillicIds(doc *goquery.Document) error {
	doc.Find("[id]").Each(func(i int, selection *goquery.Selection) {
		id, _ := selection.Attr("id")
		// if id has a cyrillic letter
		if strings.ContainsAny(id, "абвгдеєжзиіїйклмнопрстуфхцчшщьюя") {
			fixedId := helpers.GetSectionId(id)
			selection.SetAttr("id", fixedId)
		}
	})
	// Also fix Cyrillic anchor links
	doc.Find("a[href^=\"#\"]").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		id := strings.TrimPrefix(href, "#")
		// if id has a cyrillic letter
		if strings.ContainsAny(id, "абвгдеєжзиіїйклмнопрстуфхцчшщьюя") {
			fixedId := helpers.GetSectionId(id)
			selection.SetAttr("href", "#"+fixedId)
		}
	})

	return nil
}

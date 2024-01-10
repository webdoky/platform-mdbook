package main

import "github.com/PuerkitoBio/goquery"

func fixHljs(doc *goquery.Document) error {
	doc.Find(".hljs").Each(func(i int, selection *goquery.Selection) {
		if selection.HasClass("hidden") {
			selection.RemoveClass("hljs")
		} else if selection.HasClass("language-plain") {
			selection.RemoveClass("language-plain")
			selection.RemoveClass("hljs")
		} else if selection.HasClass("language-plain-nolint") {
			selection.RemoveClass("language-plain-nolint")
			selection.RemoveClass("hljs")
		} else if selection.HasClass("language-js-nolint") {
			selection.RemoveClass("language-js-nolint")
			selection.AddClass("language-js")
		} else if selection.HasClass("language-css-nolint") {
			selection.RemoveClass("language-css-nolint")
			selection.AddClass("language-css")
		} else if selection.HasClass("language-html-nolint") {
			selection.RemoveClass("language-html-nolint")
			selection.AddClass("language-html")
		} else if selection.HasClass("language-md-nolint") {
			selection.RemoveClass("language-md-nolint")
			selection.AddClass("language-md")
		} else if selection.HasClass("language-hbs-nolint") {
			selection.RemoveClass("language-hbs-nolint")
			selection.AddClass("language-hbs")
		} else if selection.HasClass("language-jsx-nolint") {
			selection.RemoveClass("language-jsx-nolint")
			selection.AddClass("language-jsx")
		}
	})
	return nil
}

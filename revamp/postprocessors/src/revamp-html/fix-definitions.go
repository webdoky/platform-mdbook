package main

import (
	"log"
	"strings"
	"webdoky3/revamp/helpers"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
)

func fixDefinitions(doc *goquery.Document) error {
	lis := doc.Find("li")
	pseudoDls := []*goquery.Selection{}
	lis.Each(func(i int, li *goquery.Selection) {
		// if li starts with ": "
		if !strings.HasPrefix(li.Text(), ": ") {
			return
		}
		pseudoDl := li.Closest("ul").Parent().Closest("ul")
		if !slices.Contains(pseudoDls, pseudoDl) {
			pseudoDls = append(pseudoDls, pseudoDl)
		}
	})
	for _, pseudoDl := range pseudoDls {
		dlHtml := ""
		pseudoDefinitions := pseudoDl.ChildrenFiltered("li")
		successMarker := true
		pseudoDefinitions.Each(func(i int, pseudoDefinition *goquery.Selection) {
			term := ""
			definition := ""
			htmlCode, err := pseudoDefinition.Html()
			if err != nil {
				log.Fatal(err)
			}
			pseudoDefinition.Contents().Each(func(i int, selection *goquery.Selection) {
				nodeHtml, err := selection.Html()
				if err != nil {
					log.Fatal(err)
				}
				if nodeHtml == "" {
					nodeHtml = selection.Text()
				}
				// log.Printf("selection, %d: %s", selection.Nodes[0].Type, nodeHtml)
				if selection.Nodes[0].Type != html.ElementNode || !selection.Is("ul") {
					term += nodeHtml
				} else {
					selection.Find("ul > li").Each(func(i int, li *goquery.Selection) {
						li.AddSelection(li.Find("p")).Each(func(i int, textContainer *goquery.Selection) {
							if strings.HasPrefix(textContainer.Text(), ": ") {
								textHtml, err := textContainer.Html()
								if err != nil {
									log.Fatal(err)
								}
								textContainer.SetHtml(strings.TrimPrefix(textHtml, ": "))
							}
						})

						liHtml, err := li.Html()
						if err != nil {
							log.Fatal(err)
						}
						definition += liHtml
					})
				}
			})
			if term == "" {
				log.Println("term is empty. " + htmlCode)
				successMarker = false
			}
			if definition == "" {
				log.Println("definition is empty. " + htmlCode)
				successMarker = false
			}
			dlHtml += "<dt id=\"" + helpers.GetSectionId(term) + "\">" + term + "</dt><dd>" + definition + "</dd>"
		})
		dlHtml = "<dl>" + dlHtml + "</dl>"
		if !successMarker {
			continue
		}
		// Replace pseudoDl with dlHtml
		pseudoDl.ReplaceWithHtml(dlHtml)
	}
	return nil
}

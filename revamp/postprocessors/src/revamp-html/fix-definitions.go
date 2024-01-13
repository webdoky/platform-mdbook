package main

import (
	"errors"
	"log"
	"strings"
	"webdoky3/revamp/helpers"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func removeColonPrefix(s *goquery.Selection) {
	s.Each(func(i int, selection *goquery.Selection) {
		htmlCode, err := selection.Html()
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(htmlCode, ": ") {
			selection.SetHtml(strings.TrimPrefix(htmlCode, ": "))
		}
	})
}

func fixDefinitions(doc *goquery.Document) error {
	lis := doc.Find("li")
	pseudoDls := []*goquery.Selection{}
	lis.Each(func(i int, li *goquery.Selection) {
		// if li starts with ": "
		if !strings.HasPrefix(li.Text(), ": ") {
			return
		}
		pseudoDl := li.Closest("ul").Parent().Closest("ul")
		for _, knownPseudoDl := range pseudoDls {
			if knownPseudoDl.IsSelection(pseudoDl) {
				return
			}
		}
		pseudoDls = append(pseudoDls, pseudoDl)
	})
	// log.Printf("number of pseudoDls: %d", len(pseudoDls))
	var mainError error
	for _, pseudoDl := range pseudoDls {
		if pseudoDl.Length() != 1 {
			return errors.New("pseudoDl.Length() != 1")
		}
		dlHtml := ""
		pseudoDefinitions := pseudoDl.ChildrenFiltered("li")
		// log.Printf("pseudoDefinitions.Length(): %d", pseudoDefinitions.Length())
		pseudoDefinitions.Each(func(i int, pseudoDefinition *goquery.Selection) {
			// if !pseudoDefinition.Is("li") {
			// 	mainError = errors.New("pseudoDefinition is not li")
			// 	return
			// }
			term := ""
			definition := ""
			htmlCode, err := pseudoDefinition.Html()
			if err != nil {
				log.Fatal(err)
			}
			processedSelections := []*goquery.Selection{}
			contents := pseudoDefinition.Contents()
			// log.Printf("contents.Length(): %d", contents.Length())
			contents.Each(func(i int, selection *goquery.Selection) {
				if !selection.Parent().IsSelection(pseudoDefinition) || selection.IsSelection(pseudoDefinition) {
					return
				}
				// log.Printf("selection, %d: %s", selection.Get(0).Type, selection.Text())
				for _, processedSelection := range processedSelections {
					if processedSelection.Contains(selection.Get(0)) {
						// log.Printf("skipping #%d", i)
						return
					}
				}
				// log.Printf("selection.Get(0).Type: %d", selection.Get(0).Type)
				// if selection.Get(0).Type == html.ElementNode {
				// 	log.Printf("selection.Get(0).Data: %s", selection.Get(0).Data)
				// }
				// log.Printf("selection, %d: %s", selection.Nodes[0].Type, nodeHtml)
				if selection.Get(0).Type != html.ElementNode || !selection.Is("ul") {
					nodeHtml, err := goquery.OuterHtml(selection)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("nodeHtml: '%s'", nodeHtml)
					if definition != "" {
						// log.Printf("definition: '%s'", definition)
						// log.Printf("nodeHtml: '%s'", nodeHtml)
						if strings.TrimSpace(nodeHtml) != "" {
							mainError = errors.New("definition is not empty, term: " + term + ", definition: " + definition + ", nodeHtml: " + nodeHtml)
						}
						return
					}
					if nodeHtml == "" {
						nodeHtml = selection.Text()
					}
					// nodeHtml = strings.TrimSpace(nodeHtml)
					if nodeHtml == "" {
						return
					}
					term += nodeHtml
				} else {
					selection.Find("ul > li").Each(func(i int, li *goquery.Selection) {
						li.Each(func(i int, textContainer *goquery.Selection) {
							// textHtml, err := textContainer.Html()
							// if err != nil {
							// 	log.Fatal(err)
							// }
							if strings.HasPrefix(textContainer.Text(), ": ") {
								removeColonPrefix(li)
							} else {
								li.Find("p").Each(func(i int, p *goquery.Selection) {
									removeColonPrefix(p)
								})
							}
						})

						liHtml, err := li.Html()
						if err != nil {
							log.Fatal(err)
						}
						if strings.Contains(definition, liHtml) {
							return
						}
						// log.Printf("liHtml #%d: %s", i, liHtml)
						definition += liHtml
					})
				}
				processedSelections = append(processedSelections, selection)
			})
			term = strings.TrimSpace(term)
			if term == "" {
				log.Println(htmlCode)
				mainError = errors.New("term is empty")
				return
			}
			if definition == "" {
				log.Println(htmlCode)
				mainError = errors.New("definition is empty")
				return
			}
			definitionWithParagraphs := ""
			log.Printf("term: '%s'", term)
			log.Printf("definition: '%s'", definition)
			for _, defLine := range strings.Split(definition, "\n\n") {
				defParagraph := strings.TrimSpace(defLine)
				if defParagraph != "" && !strings.HasPrefix(defParagraph, "<p>") {
					defParagraph = "<p>" + defParagraph + "</p>"
				}
				definitionWithParagraphs += defParagraph
			}
			// firstTerm := term
			// endIndex := strings.Index(term[1:], "<") + 1
			// if endIndex != 0 {
			// 	firstTerm = strings.TrimSpace(term[:endIndex])
			// }

			dlHtml += "<dt id=\"" + helpers.GetSectionId(term) + "\">" + term + "</dt><dd>" + definitionWithParagraphs + "</dd>"
		})
		dlHtml = "<dl>" + dlHtml + "</dl>"
		if mainError != nil {
			// return mainError
			log.Println(mainError)
			continue
		}
		// Replace pseudoDl with dlHtml
		pseudoDl.ReplaceWithHtml(dlHtml)
	}
	return nil
}

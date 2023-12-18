package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"webdoky3/revamp/helpers"
	"webdoky3/revamp/preprocessors/src/helpers/l10n"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/preprocessor"
)

var UKRAINIAN_MARKERS = []string{"//uk.", ".ua", "hl=uk"}

var MARKDOWN_LINK_REGEX = regexp.MustCompile(`([^!])\[([^\[\]]+)]\(([^\s\(\)\[\]]+)\)`)
var HTML_LINK_REGEX = regexp.MustCompile(`<a[^>]*\shref="([^>"]+)"[^>]*>`)

func checkLinkForMissing(href string) {
	if strings.HasPrefix(href, "http") {
		return
	}
	href = strings.TrimPrefix(href, "/")
	hashIndex := strings.Index(href, "#")
	if hashIndex != -1 {
		href = href[:hashIndex]
	}

	if href == "" {
		return
	}
	if !bookHasPath(href + "/index.md") {
		missingLinksCounter[href]++
	}
}

func getClassForLink(href string) string {
	href = strings.TrimPrefix(href, "/")
	if href == "" {
		return ""
	}
	classes := []string{}
	hashIndex := strings.Index(href, "#")
	if hashIndex != -1 {
		href = href[:hashIndex]
	}
	if strings.HasPrefix(href, "http") {
		classes = append(classes, "external-link")
		isUkrainian := false
		for _, marker := range UKRAINIAN_MARKERS {
			if strings.Contains(href, marker) {
				isUkrainian = true
				break
			}
		}
		if isUkrainian {
			classes = append(classes, "ukrainian-link")
		} else {
			classes = append(classes, "foreign-link")
		}
	} else {
		classes = append(classes, "internal-link")
		if !bookHasPath(href + "/index.md") {
			log.Println("Missing link: " + href)
			classes = append(classes, "missing-link")
		}
	}
	return strings.Join(classes, " ")
}

func getLinkTooltip(href string) string {

	if strings.HasPrefix(href, "http") {
		for _, marker := range UKRAINIAN_MARKERS {
			if strings.Contains(href, marker) {
				return "🇺🇦 Зовнішнє посилання українською мовою"
			}
		}
		return "🏴󠁧󠁢󠁥󠁮󠁧󠁿 Зовнішнє посилання англійською мовою"
	}
	href = strings.TrimPrefix(href, "/")
	hashIndex := strings.Index(href, "#")
	if hashIndex != -1 {
		href = href[:hashIndex]
	}
	tooltip := pathsSet[href+"/index.md"]
	if tooltip == "" {
		return l10n.L10nCommon("uk", "TranslationCTA")
	}
	return helpers.EscapeForMarkdown(tooltip)
}

func formatLinksInMarkdown(markdown string) string {
	result := markdown
	matches := HTML_LINK_REGEX.FindAllSubmatch([]byte(result), -1)
	for _, match := range matches {
		href := string(match[1])
		checkLinkForMissing(href)
		aOpeningHtml, err := renderhtml.RenderAOpening(&renderhtml.AParams{
			Class: getClassForLink(href),
			Href:  href,
			Title: getLinkTooltip(href),
		})
		if err == nil {
			// result = strings.ReplaceAll(result, string(match[0]), "<a class=\""+getClassForLink(href)+"\" href=\""+href+"\" title=\""+getLinkTooltip(href)+"\">")
			result = strings.ReplaceAll(result, string(match[0]), aOpeningHtml)
		} else {
			log.Println(err)
		}
	}
	matches = MARKDOWN_LINK_REGEX.FindAllSubmatch([]byte(result), -1)
	for _, match := range matches {
		href := string(match[3])
		text := string(match[2])
		checkLinkForMissing(href)

		aHtml, err := renderhtml.RenderA(&renderhtml.AParams{
			Class: getClassForLink(href),
			Href:  href,
			Text:  text,
			Title: getLinkTooltip(href),
		})
		if err == nil {
			// result = strings.ReplaceAll(result, string(match[0]), string(match[1])+"<a class=\""+getClassForLink(href)+"\" href=\""+href+"\" title=\""+getLinkTooltip(href)+"\">"+text+"</a>")
			result = strings.ReplaceAll(result, string(match[0]), string(match[1])+aHtml)
		} else {
			log.Println(err)
		}
	}
	return result
}

func formatLinksInSection(section *preprocessor.Section) {
	section.Chapter.Content = formatLinksInMarkdown(section.Chapter.Content)
	for _, subItem := range section.Chapter.SubItems {
		if subItem.IsSeparator {
			continue
		}
		formatLinksInSection(&subItem)
	}
}

var pathsSet = make(map[string]string)
var missingLinksCounter = make(map[string]int)

func bookHasPath(path string) bool {
	return pathsSet[path] != ""
}

func registerPath(section *preprocessor.Section) {
	pathsSet[section.Chapter.Path] = section.Chapter.Name
	for _, subItem := range section.Chapter.SubItems {
		if subItem.IsSeparator {
			continue
		}
		registerPath(&subItem)
	}
}

func formatLinks(book *preprocessor.Book, context *preprocessor.Context) (*preprocessor.Book, error) {
	// globalBook = &book
	for _, section := range book.Sections {
		if section.IsSeparator {
			continue
		}
		registerPath(&section)
	}
	// panic("stop")
	for _, section := range book.Sections {
		if section.IsSeparator {
			continue
		}
		formatLinksInSection(&section)
	}

	return book, nil
}

func main() {
	log.Println("Starting format-links")
	p := preprocessor.NewPreprocessor(formatLinks)
	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = saveMissingLinksReport()
	if err != nil {
		log.Fatal(err)
	}
	err = saveLinksReport()
	if err != nil {
		log.Fatal(err)
	}
}

func saveMissingLinksReport() error {
	markdownReport := "# Missing links report\n\n"
	markdownReport += "| Link | Count |\n"
	markdownReport += "| ---- | ----- |\n"
	linksByCountDescending := make([]string, 0, len(missingLinksCounter))
	for k := range missingLinksCounter {
		linksByCountDescending = append(linksByCountDescending, k)
	}
	sort.SliceStable(linksByCountDescending, func(i, j int) bool {
		return missingLinksCounter[linksByCountDescending[i]] > missingLinksCounter[linksByCountDescending[j]]
	})
	for _, link := range linksByCountDescending {
		markdownReport += fmt.Sprintf("| %s | %d |\n", link, missingLinksCounter[link])
	}
	markdownReport += "\n"
	err := os.WriteFile("missing-links-report.md", []byte(markdownReport), 0644)
	return err
}

func saveLinksReport() error {
	markdownReport := "# Articles report\n\n"
	linksAscending := make([]string, 0, len(pathsSet))
	for link := range pathsSet {
		linksAscending = append(linksAscending, link)
	}
	sort.Strings((linksAscending))
	for _, link := range linksAscending {
		markdownReport += fmt.Sprintf("- [%s](%s)\n", pathsSet[link], link)
	}
	markdownReport += "\n"
	err := os.WriteFile("links-report.md", []byte(markdownReport), 0644)
	return err
}

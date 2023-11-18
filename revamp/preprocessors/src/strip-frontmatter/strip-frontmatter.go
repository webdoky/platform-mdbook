package main

import (
	"log"
	"webdoky3/revamp/preprocessors/src/helpers"
	"webdoky3/revamp/preprocessors/src/preprocessor"
)

func removeFrontmatterFromSection(section *preprocessor.Section) {
	section.Chapter.Content = helpers.RemoveFrontmatter(section.Chapter.Content)
	for _, subItem := range section.Chapter.SubItems {
		if subItem.IsSeparator {
			continue
		}
		removeFrontmatterFromSection(&subItem)
	}
}

func stripFrontmatter(book *preprocessor.Book, context *preprocessor.Context) (*preprocessor.Book, error) {
	for _, section := range book.Sections {
		if section.IsSeparator {
			continue
		}
		removeFrontmatterFromSection(&section)
	}
	return book, nil
}

func main() {
	p := preprocessor.NewPreprocessor(stripFrontmatter)
	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

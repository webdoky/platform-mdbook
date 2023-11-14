package main

import (
	"log"
	"webdoky3/preprocessors/src/preprocessor"
)

const NUMBER_OF_WORKERS = 10

func injectAuthorsIntoSection(section preprocessor.Section, returnChannel chan preprocessor.Section) {
	log.Println("injectAuthorsIntoSection: " + section.Chapter.Path)
	if section.IsSeparator {
		return
	}
	originalAuthors, err := determineOriginalAuthors(section.Chapter.Path)
	if err != nil {
		log.Fatal(err)
	}
	translationAuthors, err := determineTranslationAuthors(section.Chapter.Path)
	if err != nil {
		log.Fatal(err)
	}
	if len(originalAuthors) > 0 || len(translationAuthors) > 0 {
		appendix := "\n## Автори статті\n\n"
		for _, author := range originalAuthors {
			appendix += "- " + author + "\n"
		}
		if len(translationAuthors) > 0 {
			appendix += "\n### Автори перекладу\n\n"
			for _, author := range translationAuthors {
				appendix += "- " + author + "\n"
			}
		}
		section.Chapter.Content += appendix
	}
	subItemsReturnChannel := make(chan preprocessor.Section)
	// subItemsWaitGroup.Add(len(section.Chapter.SubItems))
	for _, subItem := range section.Chapter.SubItems {
		go injectAuthorsIntoSection(subItem, subItemsReturnChannel)
	}
	for index := range section.Chapter.SubItems {
		section.Chapter.SubItems[index] = <-subItemsReturnChannel
	}
	returnChannel <- section
}

func injectAuthors(book *preprocessor.Book, context *preprocessor.Context) (*preprocessor.Book, error) {
	// log.Println("injectAuthors")
	sectionsReturnChannel := make(chan preprocessor.Section)
	// var waitGroup sync.WaitGroup
	// waitGroup.Add(len(book.Sections))
	for _, section := range book.Sections {
		go injectAuthorsIntoSection(section, sectionsReturnChannel)
	}
	for index := range book.Sections {
		book.Sections[index] = <-sectionsReturnChannel
	}
	// waitGroup.Wait()
	return book, nil
}

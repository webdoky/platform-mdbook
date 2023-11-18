package main

import (
	"log"
	"os"
	"path/filepath"
	"webdoky3/revamp/preprocessors/src/helpers"
	"webdoky3/revamp/preprocessors/src/preprocessor"
)

func createSummary(book *preprocessor.Book, context *preprocessor.Context) (*preprocessor.Book, error) {
	root := context.Root
	srcFolder := context.Config.Book.Source
	absoluteSrcFolder, err := filepath.Abs(root + "/" + srcFolder)
	if err != nil {
		return nil, err
	}
	summaryPath := absoluteSrcFolder + "/SUMMARY.md"
	summary, err := GetSummaryItemFromFolder(srcFolder, srcFolder)
	if err != nil {
		return nil, err
	}
	summaryContent := "# Зміст\n\n"
	var traverseSummaryItem func(summaryItem *SummaryItem, depth int)
	traverseSummaryItem = func(summaryItem *SummaryItem, depth int) {
		summaryContent += helpers.GetIndentation(depth) + "- [" + helpers.EscapeForMarkdown(summaryItem.title) + "](./" + summaryItem.path + ")\n"
		log.Println("Traversing summary item: " + summaryItem.path)
		for _, subItem := range summaryItem.sub_items {
			traverseSummaryItem(subItem, depth+1)
		}
	}
	traverseSummaryItem(summary, 0)
	// write summaryContent by summaryPath
	err = os.WriteFile(summaryPath, []byte(summaryContent), 0644)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func main() {
	log.Println("Starting summary")
	p := preprocessor.NewPreprocessor(createSummary)
	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	// os.Exit(1)
}

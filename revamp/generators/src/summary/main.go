package main

import (
	"log"
	"os"
	"path/filepath"
	"webdoky3/revamp/helpers"
	preprocessor_helpers "webdoky3/revamp/preprocessors/src/helpers"
)

var sourceFolder = "./content/files/uk"

func createSummary() error {
	var rootFolder, err = os.Getwd()
	if err != nil {
		return err
	}
	absoluteSrcFolder, err := filepath.Abs(rootFolder + "/" + sourceFolder)
	if err != nil {
		return err
	}
	summaryPath := absoluteSrcFolder + "/SUMMARY.md"
	summary, err := GetSummaryItemFromFolder(sourceFolder, sourceFolder)
	if err != nil {
		return err
	}
	summaryContent := "# Зміст\n\n"
	var traverseSummaryItem func(summaryItem *SummaryItem, depth int)
	traverseSummaryItem = func(summaryItem *SummaryItem, depth int) {
		summaryContent += preprocessor_helpers.GetIndentation(depth) + "- [" + helpers.EscapeForMarkdown(summaryItem.title) + "](./" + summaryItem.path + ")\n"
		log.Println("Traversing summary item: " + summaryItem.path)
		for _, subItem := range summaryItem.sub_items {
			traverseSummaryItem(subItem, depth+1)
		}
	}
	traverseSummaryItem(summary, 0)
	// write summaryContent by summaryPath
	return os.WriteFile(summaryPath, []byte(summaryContent), 0644)
}

func main() {
	err := createSummary()
	if err != nil {
		log.Fatal(err)
	}
}

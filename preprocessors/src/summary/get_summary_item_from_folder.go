package main

import (
	"log"
	"path/filepath"
	"sort"
	"webdoky3/preprocessors/src/helpers"
)

type SummaryItem struct {
	path      string
	title     string
	sub_items []*SummaryItem
}

func GetSummaryItemFromFolder(folderPath string, sourceFolderPath string) (*SummaryItem, error) {
	log.Println("Getting summary item from folder: " + folderPath)
	summaryItem := SummaryItem{sub_items: []*SummaryItem{}}
	indexMdPath := folderPath + "/index.md"
	doesIndexMdExist := helpers.DoesFileExist(indexMdPath)
	var err error
	if doesIndexMdExist {
		summaryItem.path, err = filepath.Rel(sourceFolderPath, indexMdPath)
		if err != nil {
			log.Printf("Error getting relative path for %s: %s", indexMdPath, err)
			return nil, err
		}
		frontmatterData, err := helpers.GetFrontmatterData(indexMdPath)
		if err != nil {
			log.Printf("Error getting frontmatter data for %s: %s", indexMdPath, err)
			return nil, err
		}
		summaryItem.title = frontmatterData.Title
	} else {
		log.Println("Index.md does not exist in folder: " + folderPath)
		return nil, nil
	}
	files := helpers.GetFilesInFolder(folderPath)
	log.Printf("Files in folder: %d", len(files))
	for _, file := range files {
		if file.IsDir() {
			subItem, err := GetSummaryItemFromFolder(folderPath+"/"+file.Name(), sourceFolderPath)
			if err != nil {
				return nil, err
			}
			if subItem != nil {
				summaryItem.sub_items = append(summaryItem.sub_items, subItem)
			}
		} else {
			log.Printf("File %s is not a directory", file.Name())
		}
	}
	if len(files)-1 != len(summaryItem.sub_items) {
		log.Printf("Files in folder %s: %d, subitems: %d", folderPath, len(files), len(summaryItem.sub_items))
	}
	// sort summaryItem.sub_items by title
	sort.SliceStable(summaryItem.sub_items, func(i, j int) bool {
		return summaryItem.sub_items[i].title < summaryItem.sub_items[j].title
	})
	return &summaryItem, nil
}

package main

import (
	"webdoky3/preprocessors/src/helpers"
	"webdoky3/preprocessors/src/preprocessor"
)

func alterSection(section *preprocessor.Section, sourcePath string) error {
	fullSectionPath := sourcePath + "/" + section.Chapter.Path
	data, err := helpers.GetFrontmatterData(fullSectionPath)
	if err != nil {
		return err
	}
	if data.Slug != "" {
		section.Chapter.Path = "uk/docs/" + data.Slug + "/index.md"
	}
	if data.Title != "" {
		section.Chapter.Content = "# " + helpers.EscapeForMarkdown(data.Title) + "\n\n" + section.Chapter.Content
	} else {
		panic("Title is empty")
	}
	for _, subItem := range section.Chapter.SubItems {
		if subItem.IsSeparator {
			continue
		}
		err = alterSection(&subItem, sourcePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func rewirePaths(book *preprocessor.Book, context *preprocessor.Context) (*preprocessor.Book, error) {
	sourceFolder := context.Config.Book.Source
	var err error
	for _, section := range book.Sections {
		if section.IsSeparator {
			continue
		}
		err = alterSection(&section, sourceFolder)
		if err != nil {
			return nil, err
		}
	}
	return book, nil
}

func main() {
	p := preprocessor.NewPreprocessor(rewirePaths)
	error := p.Run()
	if error != nil {
		panic(error)
	}
}

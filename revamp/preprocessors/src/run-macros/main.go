package main

import (
	"log"
	"webdoky3/revamp/preprocessors/src/preprocessor"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
	"webdoky3/revamp/preprocessors/src/run-macros/runner"
)

var locale = "uk"

func runMacrosInSection(registryInstance *registry.Registry, section *preprocessor.Section) error {
	log.Printf("Running macros in section %s", section.Chapter.Path)
	frontmatterData, err := get_section_frontmatter(section)
	if err != nil {
		return err
	}
	macrosRunner := runner.NewMacrosRunner(&environment.Environment{
		Content:     &section.Chapter.Content,
		Frontmatter: frontmatterData,
		Locale:      locale,
		Path:        section.Chapter.Path,
	}, registryInstance)
	section.Chapter.Content = macrosRunner.Run(section.Chapter.Content)
	for _, subItem := range section.Chapter.SubItems {
		if subItem.IsSeparator {
			continue
		}
		err := runMacrosInSection(registryInstance, &subItem)
		if err != nil {
			return err
		}
	}
	return nil
}

func runMacros(book *preprocessor.Book, context *preprocessor.Context) (*preprocessor.Book, error) {
	registryInstance := registry.NewRegistry(book)
	for _, section := range book.Sections {
		if section.IsSeparator {
			continue
		}
		err := runMacrosInSection(registryInstance, &section)
		if err != nil {
			return nil, err
		}
	}
	return book, nil
}

func main() {
	p := preprocessor.NewPreprocessor(runMacros)
	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

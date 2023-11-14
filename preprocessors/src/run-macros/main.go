package main

import (
	"log"
	"webdoky3/preprocessors/src/preprocessor"
	"webdoky3/preprocessors/src/run-macros/macros"
	"webdoky3/preprocessors/src/run-macros/registry"
	"webdoky3/preprocessors/src/run-macros/runner"
)

func runMacrosInSection(macrosRunner *runner.MacrosRunner, section *preprocessor.Section) error {
	section.Chapter.Content = macrosRunner.Run(section.Chapter.Content)
	for _, subItem := range section.Chapter.SubItems {
		if subItem.IsSeparator {
			continue
		}
		err := runMacrosInSection(macrosRunner, &subItem)
		if err != nil {
			return err
		}
	}
	return nil
}

func runMacros(book *preprocessor.Book, context *preprocessor.Context) (*preprocessor.Book, error) {
	macrosRunner := runner.NewMacrosRunner(&macros.Environment{
		Locale: "uk",
	}, registry.NewRegistry(book))
	for _, section := range book.Sections {
		if section.IsSeparator {
			continue
		}
		err := runMacrosInSection(macrosRunner, &section)
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

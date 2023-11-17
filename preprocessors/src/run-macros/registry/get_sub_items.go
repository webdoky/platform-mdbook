package registry

import (
	"strings"
	"webdoky3/preprocessors/src/preprocessor"
)

func findSubItemsInSections(section *preprocessor.Section, path string) []preprocessor.Section {
	// log.Printf("section.Chapter.Path: %q, path: %q", section.Chapter.Path, path)
	if section.Chapter.Path == path {
		return section.Chapter.SubItems
	}
	for _, subItem := range section.Chapter.SubItems {
		if subItem.IsSeparator {
			continue
		}
		subItems := findSubItemsInSections(&subItem, path)
		if subItems != nil {
			return subItems
		}
	}
	return nil
}

func (r *Registry) GetSubItems(path string) []preprocessor.Section {
	if !strings.HasSuffix(path, "/index.md") {
		path += "/index.md"
	}
	for _, section := range r.book.Sections {
		if section.IsSeparator {
			continue
		}
		subItems := findSubItemsInSections(&section, path)
		if subItems != nil {
			return subItems
		}
	}
	return []preprocessor.Section{}
}

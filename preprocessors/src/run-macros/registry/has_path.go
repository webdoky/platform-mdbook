package registry

import (
	"strings"
	"webdoky3/preprocessors/src/preprocessor"
)

func sectionHasPath(section *preprocessor.Section, path string) bool {
	if section.Chapter.Path == path {
		return true
	}
	for _, subItem := range section.Chapter.SubItems {
		if subItem.IsSeparator {
			continue
		}
		value := sectionHasPath(&subItem, path)
		if value {
			return true
		}
	}
	return false
}

func (r *Registry) HasPath(path string) bool {
	if !strings.HasSuffix(path, "/index.md") {
		path += "/index.md"
	}
	for _, section := range r.book.Sections {
		if section.IsSeparator {
			continue
		}
		value := sectionHasPath(&section, path)
		if value {
			return true
		}
	}
	return false
}

package main

import (
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
	"webdoky3/revamp/preprocessors/src/preprocessor"
)

func get_section_frontmatter(section *preprocessor.Section) (*helpers.FrontmatterData, error) {
	if section.IsSeparator {
		return nil, nil
	}
	filePath := strings.ToLower(section.Chapter.Path)
	filePath = strings.Replace(filePath, locale+"/docs", "content/files/"+locale, 1)
	filePath = strings.Replace(filePath, "::", "_doublecolon_", -1)
	filePath = strings.Replace(filePath, ":", "_colon_", -1)
	filePath = strings.Replace(filePath, "*", "_star_", -1)
	return helpers.GetFrontmatterData(filePath)
}

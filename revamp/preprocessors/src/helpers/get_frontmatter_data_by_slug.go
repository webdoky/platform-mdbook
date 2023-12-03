package helpers

import (
	"os"
	"strings"
)

func GetFrontmatterDataBySlug(slug string, locale string) (*FrontmatterData, error) {
	// try to read markdown file from "content" directory
	// if file does not exist, try to read markdown file from "original-content" directory
	// if file does not exist, return error
	// if file exists, extract frontmatter data
	var markdown string
	slug = strings.Replace(slug, "::", "_doublecolon_", -1)
	slug = strings.Replace(slug, ":", "_colon_", -1)
	slug = strings.Replace(slug, "*", "_star_", -1)
	file, err := os.ReadFile("./content/files/" + locale + "/" + slug + "/index.md")
	if err == nil {
		markdown = string(file)
	} else {
		// if error is not "no such file or directory", return error
		if !os.IsNotExist(err) {
			return nil, err
		}
		file, err = os.ReadFile("./original-content/files/en-us/" + slug + "/index.md")
		if err == nil {
			markdown = string(file)
		} else {
			if !os.IsNotExist(err) {
				return nil, err
			}
			return nil, nil
		}
	}
	return ExtractFrontmatterData(markdown)
}

package main

import (
	"regexp"
	"strings"
	"webdoky3/revamp/helpers"
)

var UPDATE_REGEX = regexp.MustCompile(`\*\*(?:Переклад|Оновлення перекладу)\(\w+\):\*\* ([^\s]+) `)

func populate_links(md string) (string, error) {
	lines := strings.Split(md, "\n")
	result := ""
	for _, line := range lines {
		// try to find UPDATE_REGEX match
		match := UPDATE_REGEX.FindStringSubmatch(line)
		if match != nil {
			path := match[1]
			filePath := "content/files/uk/" + path + "/index.md"
			frontmatterData, err := helpers.GetFrontmatterData(filePath)
			if err != nil {
				if !strings.Contains(err.Error(), "no such file or directory") {
					return "", err
				}
			} else {
				line = strings.ReplaceAll(line, match[1], "["+frontmatterData.Title+"](/uk/docs/"+frontmatterData.Slug+")")
			}
		}
		result += line + "\n"
	}
	return result, nil
}

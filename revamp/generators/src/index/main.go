package main

import (
	"log"
	"os"
	"strings"
)

func getTemplate() (string, error) {
	// read src/index-template.md
	// and return it
	indexTemplate, err := os.ReadFile("src/index-template.md")
	if err != nil {
		return "", err
	}
	return string(indexTemplate), nil
}

const MAX_SECTIONS = 2

func main() {
	// Read content/CHANGELOG.md file
	fileContent, err := os.ReadFile("content/CHANGELOG.md")
	if err != nil {
		log.Fatal(err)
	}
	fileContentString := string(fileContent)
	// split into lines
	// find first line that starts with "## "
	// find next line that starts with "## "
	// and find line that starts with ## after that
	// and replace "## " with "### "
	lines := strings.Split(fileContentString, "\n")
	result := ""
	sectionsFound := 0
	for i, line := range lines {
		log.Println(line)
		if strings.HasPrefix(line, "## ") {
			sectionsFound++
			if sectionsFound > MAX_SECTIONS {
				break
			}
			result += strings.ReplaceAll(line, "## ", "### ")
		} else {
			if sectionsFound > 0 {
				result += line
			}
		}
		if i != 0 {
			result += "\n"
		}
	}

	template, err := getTemplate()
	if err != nil {
		log.Fatal(err)
	}
	md := strings.Replace(template, "{{CHANGELOG}}", result, 1)
	// Write md to content/files/uk/index.md file
	err = os.WriteFile("content/files/uk/index.md", []byte(md), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

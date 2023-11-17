package helpers

import (
	"log"
	"os"
)

func GetFrontmatterData(markdownFilePath string) (*FrontmatterData, error) {
	markdown, err := os.ReadFile(markdownFilePath)
	if err != nil {
		return nil, err
	}

	frontmatterData, err := ExtractFrontmatterData(string(markdown))
	if err != nil {
		log.Printf("Error parsing frontmatter for %s: %s", markdownFilePath, err)
		return nil, err
	}

	return frontmatterData, err
}

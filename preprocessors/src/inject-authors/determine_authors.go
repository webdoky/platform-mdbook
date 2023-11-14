package main

import (
	"log"
	"os/exec"
)

// var COMMAND_PATTERN = "cd %s && git log --pretty=full --follow -- %s"

func determineAuthors(repositoryPath string, filePath string, locale string) ([]string, error) {
	log.Println("determineAuthors: " + filePath)
	// filePath = strings.TrimSuffix(filePath, "/index.md")
	if filePath == "index.md" {
		// Root index.md is not a part of content
		return []string{}, nil
	}
	// command := fmt.Sprintf(COMMAND_PATTERN, repositoryPath, filePath)
	// log.Println(command)
	// output, err := exec.Command(fmt.Sprintf(COMMAND_PATTERN, repositoryPath, filePath)).Output()
	output, err := exec.Command("git", "-C", repositoryPath, "log", "--pretty=full", "--follow", "--", "files/"+locale+"/"+filePath).Output()
	if err != nil {
		return nil, err
	}
	return getAuthorsFromOutput(string(output))
}

func determineOriginalAuthors(filePath string) ([]string, error) {
	return determineAuthors("original-content", filePath, "en-us")
}

func determineTranslationAuthors(filePath string) ([]string, error) {
	return determineAuthors("content", filePath, "uk")
}

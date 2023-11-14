package main

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

var MAIN_AUTHOR_REGEX = regexp.MustCompile(`Author: (([^<>]+)(?: <(.+)>)?)\n`)
var CO_AUTHOR_REGEX = regexp.MustCompile(`Co-authored-by: (([^<>]+)(?: <(.+)>)?)\n`)

func getMainAuthor(commit string) (string, error) {
	match := MAIN_AUTHOR_REGEX.FindStringSubmatch(commit)
	if match == nil || len(match) < 2 {
		log.Println(commit)
		return "", errors.New("main author not found")
	}
	if len(match) > 3 && strings.Contains(match[3], "noreply.github.com") {
		return match[2], nil
		// Otherwise – use name and email
	}
	return match[1], nil
}

func getAuthorsFromCommit(commit string) ([]string, error) {
	mainAuthor, err := getMainAuthor(commit)
	if err != nil {
		return nil, err
	}

	authors := []string{mainAuthor}
	coAuthorMatches := CO_AUTHOR_REGEX.FindAllStringSubmatch(commit, -1)
	for _, coAuthorMatch := range coAuthorMatches {
		// If the co-author has a noreply email, use only name instead
		if strings.Contains(coAuthorMatch[3], "noreply.github.com") {
			authors = append(authors, coAuthorMatch[2])
			// Otherwise – use name and email
		} else {
			authors = append(authors, coAuthorMatch[1])
		}
	}
	// Deduplicate authors
	uniqueAuthors := []string{}
	for _, author := range authors {
		if !slices.Contains(uniqueAuthors, author) {
			uniqueAuthors = append(uniqueAuthors, author)
		}
	}
	return uniqueAuthors, nil
}

package main

import (
	"sort"
)

func getAuthorsFromOutput(output string) ([]string, error) {
	// log.Println("getAuthorsFromOutput")
	commits := splitOutputIntoCommits(output)
	authorsCounter := map[string]int{}
	for _, commit := range commits {
		commitAuthors, err := getAuthorsFromCommit(commit)
		if err != nil {
			return nil, err
		}
		for _, author := range commitAuthors {
			authorsCounter[author]++
		}
	}
	authors := []string{}
	for author := range authorsCounter {
		authors = append(authors, author)
	}
	sort.SliceStable(authors, func(i, j int) bool {
		return authorsCounter[authors[i]] > authorsCounter[authors[j]]
	})
	return authors, nil
}

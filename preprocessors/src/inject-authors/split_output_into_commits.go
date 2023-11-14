package main

import (
	"strings"
)

func splitOutputIntoCommits(output string) []string {
	lines := strings.Split(output, "\n")
	commits := []string{}
	var lastCommit string
	inCommit := false
	for _, line := range lines {
		if inCommit {
			if strings.HasPrefix(line, "commit") {
				commits = append(commits, lastCommit)
				lastCommit = strings.TrimSpace(line)
			} else {
				lastCommit += "\n" + strings.TrimSpace(line)
			}
		} else {
			if strings.HasPrefix(line, "commit") {
				lastCommit = strings.TrimSpace(line)
				inCommit = true
			}
		}
	}
	if inCommit && lastCommit != "" {
		// log.Println(lastCommit)
		commits = append(commits, lastCommit)
	}
	return commits
}

package embedlivesample

import (
	"errors"
	"log"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
)

func extractTargetSection(content string, parentId string) (string, string, error) {
	var i, j int
	macroStartIndex := strings.Index(strings.ToLower(content), "{{embedlivesample")
	beforeMacro := content[:macroStartIndex]
	macroAndAfter := content[macroStartIndex:]
	lines := strings.Split(beforeMacro, "\n")
	var depth int
	var headerText string
	for index := len(lines) - 1; index >= 0; index-- {
		line := lines[index]
		if strings.HasPrefix(strings.ToLower(line), "{{embedlivesample") {
			break
		}
		if strings.HasPrefix(line, "#") {
			firstSpaceIndex := strings.Index(line, " ")
			depth = firstSpaceIndex
			headerText = strings.TrimSpace(line[firstSpaceIndex+1:])
			if parentId == "" || helpers.GetSectionId(headerText) == parentId {
				i = strings.Index(content, line) + len(line)
				break
			}
		}
	}

	if depth == 0 {
		log.Println("ERROR")
		return "", "", errors.New("macro argument is wrong")
	}
	j = strings.Index(macroAndAfter, strings.Repeat("#", depth))
	if j == -1 {
		j = len(content)
	} else {
		j += macroStartIndex
	}
	return (content)[i:j], headerText, nil
}

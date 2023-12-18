package embedlivesample

import (
	"errors"
	"log"
	"strings"
	preprocessor_helpers "webdoky3/revamp/preprocessors/src/helpers"
)

func extractTargetSection(content string, parentId string) (string, string, error) {
	log.Printf("extractTargetSection: %s\n", parentId)
	var i, j int
	// macroStartIndex := strings.Index(strings.ToLower(content), "{{embedlivesample")
	macroStartIndex := -1
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(strings.ToLower(line), "{{embedlivesample") && strings.Contains(line, parentId) {
			macroStartIndex = strings.Index(content, line)
		}
	}
	if macroStartIndex == -1 {
		return "", "", errors.New("macro argument is wrong")
	}
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
			if parentId == "" || preprocessor_helpers.GetSectionId(headerText) == parentId {
				i = strings.Index(content, line) + len(line)
				break
			}
		}
	}

	if depth == 0 || i == 0 {
		return "", "", errors.New("macro argument is wrong")
	}
	j = strings.Index(macroAndAfter, strings.Repeat("#", depth))
	if j == -1 {
		j = len(content)
	} else {
		j += macroStartIndex
	}
	log.Printf("extractTargetSection indexes: %d %d\n", i, j)
	return (content)[i:j], headerText, nil
}

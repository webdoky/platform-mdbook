package embedlivesample

import (
	"strings"
)

func extractCodeItems(text string) []CodeItem {
	var result []CodeItem
	var i, j int
	for {
		i = strings.Index(text, "```")
		if i == -1 {
			break
		}
		j = strings.Index(text[i+3:], "```")
		if j == -1 {
			break
		}
		j += i + 3
		codeBlock := text[i+3 : j]
		codeBlockLines := strings.Split(codeBlock, "\n")
		language := codeBlockLines[0]
		language = strings.Split(language, " ")[0]
		language = strings.TrimSuffix(language, "-nolint")
		// log.Println("language:", language)
		code := strings.Join(codeBlockLines[1:], "\n")
		result = append(result, CodeItem{
			Code:     code,
			Language: language,
		})
		text = text[j+3:]
	}
	return result
}

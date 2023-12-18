package preprocessor_helpers

import (
	"regexp"
	"strings"
	"webdoky3/revamp/preprocessors/src/preprocessor"
)

var MARKDOWN_LINK_REGEX = regexp.MustCompile(`\[(.*?)\]\(.*?\)`)
var MACRO_REGEX = regexp.MustCompile(`\{\{.*?\}\}`)
var MARKDOWN_STRONG_REGEX = regexp.MustCompile(`\*\*(.*?)\*\*`)

func GetSummary(section *preprocessor.Section, startIndex int) string {
	if startIndex >= len(section.Chapter.Content) {
		return ""
	}
	summary := section.Chapter.Content[startIndex:]
	i := strings.Index(summary, "\n\n")
	j := strings.Index(summary[i+2:], "\n\n") + i + 2
	summary = summary[i+2 : j]
	// Remove all markdown links
	summary = MARKDOWN_LINK_REGEX.ReplaceAllString(summary, "$1")
	// Remove macros
	summary = MACRO_REGEX.ReplaceAllString(summary, "")
	summary = strings.TrimSpace(summary)
	if summary == "" || strings.HasPrefix(summary, "---") || strings.HasPrefix(summary, "#") {
		return GetSummary(section, j+2)
	}
	summary = MARKDOWN_STRONG_REGEX.ReplaceAllString(summary, "<strong>$1</strong>")
	return summary
}

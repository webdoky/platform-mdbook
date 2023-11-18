package helpers

import (
	"regexp"
)

var AMPERSAND_RE = regexp.MustCompile(`&`)
var LESS_THAN_RE = regexp.MustCompile(`<`)
var MORE_THAN_RE = regexp.MustCompile(`>`)
var OPEN_BRACKET_RE = regexp.MustCompile(`\[`)
var CLOSE_BRACKET_RE = regexp.MustCompile(`\]`)

func EscapeForMarkdown(text string) string {
	text = AMPERSAND_RE.ReplaceAllString(text, "&amp;")

	text = LESS_THAN_RE.ReplaceAllString(text, "&lt;")

	text = MORE_THAN_RE.ReplaceAllString(text, "&gt;")

	text = OPEN_BRACKET_RE.ReplaceAllString(text, "&#91;")

	text = CLOSE_BRACKET_RE.ReplaceAllString(text, "&#93;")

	return text
}

package helpers

import (
	"html"
	"regexp"
	"strings"

	"github.com/fre5h/transliteration-go"
	strip "github.com/grokify/html-strip-tags-go"
)

var CHARACTERS_TO_REMOVE = regexp.MustCompile(`[^a-z0-9-]`)


func GetSectionId(text string) string {
	text = strip.StripTags(text)
	text = html.UnescapeString(text)
	text = transliteration.UkrToLat(text)
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "-")
	// Remove all characters except letters, numbers, and hyphens
	text = CHARACTERS_TO_REMOVE.ReplaceAllString(text, "")
	return text
}

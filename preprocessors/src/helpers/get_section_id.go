package helpers

import (
	"regexp"
	"strings"

	"github.com/fre5h/transliteration-go"
)

var CHARACTERS_TO_REMOVE = regexp.MustCompile(`[^a-z0-9-]`)

func GetSectionId(text string) string {
	text = transliteration.UkrToLat(text)
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "-")
	// Remove all characters except letters, numbers, and hyphens
	text = CHARACTERS_TO_REMOVE.ReplaceAllString(text, "")
	return text
}

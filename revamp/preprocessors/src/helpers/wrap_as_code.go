package preprocessor_helpers

import (
	"html"
	"log"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
)

func WrapAsCode(text string) string {
	// Unescape HTML entities
	text = html.UnescapeString(text)
	result, err := renderhtml.RenderCode(&renderhtml.CodeParams{Text: text})
	if err != nil {
		log.Fatal(err)
	}
	return result
}

package helpers

import (
	"log"
	renderhtml "webdoky3/preprocessors/src/helpers/render_html"
)

func WrapAsCode(text string) string {
	result, err := renderhtml.RenderCode(&renderhtml.WrapperParams{Text: text})
	if err != nil {
		log.Fatal(err)
	}
	return result
}

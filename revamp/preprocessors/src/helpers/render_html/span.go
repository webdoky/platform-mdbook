package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tSpan *template.Template

type SpanParams struct {
	Class template.HTMLAttr
	Style template.HTMLAttr
	Text  string
}

func RenderSpan(params *SpanParams) (string, error) {
	var b bytes.Buffer
	err := tSpan.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tSpan, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/span.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tSpan, err = template.ParseFiles("../../helpers/render_html/templates/span.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

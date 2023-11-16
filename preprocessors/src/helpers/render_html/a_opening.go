package renderhtml

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

var tAO *template.Template

func RenderAOpening(params *AParams) (string, error) {
	var b bytes.Buffer
	err := tAO.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tAO, err = template.ParseFiles("./preprocessors/src/helpers/render_html/templates/a_opening.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tAO, err = template.ParseFiles("../../helpers/render_html/templates/a_opening.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

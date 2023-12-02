package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tBr *template.Template

func RenderBr() (string, error) {
	var b bytes.Buffer
	err := tBr.Execute(&b, nil)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tBr, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/br.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tBr, err = template.ParseFiles("../../helpers/render_html/templates/br.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tP *template.Template

type PParams struct {
	InnerHtml template.HTML
	Text      string
}

func RenderP(params *PParams) (string, error) {
	var b bytes.Buffer
	err := tP.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tP, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/p.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tP, err = template.ParseFiles("../../helpers/render_html/templates/p.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

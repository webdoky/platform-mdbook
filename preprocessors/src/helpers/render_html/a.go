package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tA *template.Template

type AParams struct {
	Class     string
	Href      string
	InnerHtml template.HTML
	Text      string
	Title     string
}

func RenderA(params *AParams) (string, error) {
	var b bytes.Buffer
	if (params.InnerHtml == "") == (params.Text == "") {
		return "", errors.New("either InnerHtml or Text must be set")
	}
	err := tA.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tA, err = template.ParseFiles("./preprocessors/src/helpers/render_html/templates/a.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tA, err = template.ParseFiles("../../helpers/render_html/templates/a.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

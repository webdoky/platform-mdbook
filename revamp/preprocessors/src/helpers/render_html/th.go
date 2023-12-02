package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tTh *template.Template

type ThParams struct {
	InnerHtml template.HTML
	Scope     string
	Text      string
}

func RenderTh(params *ThParams) (string, error) {
	var b bytes.Buffer
	if (params.InnerHtml == "") == (params.Text == "") {
		return "", errors.New("either InnerHtml or Text must be set")
	}
	err := tTh.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tTh, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/th.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tTh, err = template.ParseFiles("../../helpers/render_html/templates/th.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

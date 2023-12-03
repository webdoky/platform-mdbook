package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tTd *template.Template

type TdParams struct {
	InnerHtml template.HTML
	Text      string
}

func RenderTd(params *TdParams) (string, error) {
	var b bytes.Buffer
	if (params.InnerHtml == "") == (params.Text == "") {
		return "", errors.New("in a td tag, either InnerHtml or Text must be set")
	}
	err := tTd.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tTd, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/td.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tTd, err = template.ParseFiles("../../helpers/render_html/templates/td.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

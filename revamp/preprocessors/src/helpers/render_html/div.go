package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tDiv *template.Template

type DivParams struct {
	Class     string
	Data      map[string]template.HTMLAttr
	InnerHtml template.HTML
	Text      string
}

func RenderDiv(params *DivParams) (string, error) {
	var b bytes.Buffer
	if (params.InnerHtml == "") == (params.Text == "") {
		return "", errors.New("in an div tag, either InnerHtml or Text must be set")
	}
	err := tDiv.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tDiv, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/div.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tDiv, err = template.ParseFiles("../../helpers/render_html/templates/div.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

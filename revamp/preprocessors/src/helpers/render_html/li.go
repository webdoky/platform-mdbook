package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tLi *template.Template

type LiParams struct {
	InnerHtml template.HTML
	Text      string
}

func RenderLi(params *LiParams) (string, error) {
	var b bytes.Buffer
	if (params.InnerHtml == "") == (params.Text == "") {
		return "", errors.New("in a li tag, either InnerHtml or Text must be set")
	}
	err := tLi.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tLi, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/li.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tLi, err = template.ParseFiles("../../helpers/render_html/templates/li.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

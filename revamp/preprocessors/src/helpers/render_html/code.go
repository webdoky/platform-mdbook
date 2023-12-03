package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tCode *template.Template

type CodeParams struct {
	InnerHtml template.HTML
	Text      string
}

func RenderCode(params *CodeParams) (string, error) {
	var b bytes.Buffer
	if (params.InnerHtml != "") && (params.Text != "") {
		return "", errors.New("in a code tag, either InnerHtml or Text must be set")
	}
	err := tCode.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tCode, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/code.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tCode, err = template.ParseFiles("../../helpers/render_html/templates/code.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

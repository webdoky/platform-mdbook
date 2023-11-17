package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tOl *template.Template

type OlParams struct {
	InnerHtml template.HTML
	Text      string
}

func RenderOl(params *OlParams) (string, error) {
	if (params.InnerHtml == "") == (params.Text == "") {
		return "", errors.New("either InnerHtml or Text must be set")
	}
	var b bytes.Buffer
	err := tOl.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tOl, err = template.ParseFiles("./preprocessors/src/helpers/render_html/templates/ol.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tOl, err = template.ParseFiles("../../helpers/render_html/templates/ol.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

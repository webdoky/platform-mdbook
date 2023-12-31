package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tUl *template.Template

type UlParams struct {
	InnerHtml template.HTML
}

func RenderUl(params *UlParams) (string, error) {
	if params.InnerHtml == "" {
		return "", errors.New("InnerHtml must be set")
	}
	var b bytes.Buffer
	err := tUl.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tUl, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/ul.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tUl, err = template.ParseFiles("../../helpers/render_html/templates/ul.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

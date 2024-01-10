package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tAbbr *template.Template

type AbbrParams struct {
	Class     template.HTMLAttr
	InnerHtml template.HTML
	Text      template.HTML
	Title     template.HTMLAttr
}

func RenderAbbr(params *AbbrParams) (string, error) {
	if params.InnerHtml != "" && params.Text != "" {
		return "", errors.New("in an abbr tag, either InnerHtml or Text must be set")
	}
	var b bytes.Buffer
	err := tAbbr.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tAbbr, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/abbr.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tAbbr, err = template.ParseFiles("../../helpers/render_html/templates/abbr.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

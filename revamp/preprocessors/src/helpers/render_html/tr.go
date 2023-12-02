package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tTr *template.Template

type TrParams struct {
	InnerHtml template.HTML
}

func RenderTr(params *TrParams) (string, error) {
	var b bytes.Buffer
	err := tTr.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tTr, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/tr.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tTr, err = template.ParseFiles("../../helpers/render_html/templates/tr.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

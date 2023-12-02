package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tTbody *template.Template

type TbodyParams struct {
	InnerHtml template.HTML
}

func RenderTbody(params *TbodyParams) (string, error) {
	var b bytes.Buffer
	err := tTbody.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tTbody, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/tbody.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tTbody, err = template.ParseFiles("../../helpers/render_html/templates/tbody.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

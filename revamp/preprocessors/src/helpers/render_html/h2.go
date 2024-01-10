package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tH2 *template.Template

func RenderH2(params *WrapperParams) (string, error) {
	var b bytes.Buffer
	err := tH2.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tH2, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/h2.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tH2, err = template.ParseFiles("../../helpers/render_html/templates/h2.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

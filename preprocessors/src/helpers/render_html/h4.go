package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tH4 *template.Template

func RenderH4(params *WrapperParams) (string, error) {
	var b bytes.Buffer
	err := tH4.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tH4, err = template.ParseFiles("./preprocessors/src/helpers/render_html/templates/h4.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tH4, err = template.ParseFiles("../../helpers/render_html/templates/h4.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

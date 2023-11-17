package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tCode *template.Template

func RenderCode(params *WrapperParams) (string, error) {
	var b bytes.Buffer
	err := tCode.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tCode, err = template.ParseFiles("./preprocessors/src/helpers/render_html/templates/code.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tCode, err = template.ParseFiles("../../helpers/render_html/templates/code.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

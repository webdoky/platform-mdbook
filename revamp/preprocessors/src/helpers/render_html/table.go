package renderhtml

import (
	"bytes"
	"html/template"
	"log"
	"strings"
)

var tTable *template.Template

type TableParams struct {
	Class     template.HTMLAttr
	InnerHtml template.HTML
}

func RenderTable(params *TableParams) (string, error) {
	var b bytes.Buffer
	err := tTable.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tTable, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/table.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tTable, err = template.ParseFiles("../../helpers/render_html/templates/table.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

package renderhtml

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"strings"
)

var tIframe *template.Template

type IframeParams struct {
	Class  template.HTMLAttr
	Height template.HTMLAttr
	Src    template.HTMLAttr
	Title  template.HTMLAttr
}

func RenderIframe(params *IframeParams) (string, error) {
	var b bytes.Buffer
	if params.Src == "" {
		return "", errors.New("in an iframe tag, src must be set")
	}
	err := tIframe.Execute(&b, params)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func init() {
	var err error
	tIframe, err = template.ParseFiles("./revamp/preprocessors/src/helpers/render_html/templates/iframe.tmpl")
	if err != nil && strings.Contains(err.Error(), "no such file") {
		tIframe, err = template.ParseFiles("../../helpers/render_html/templates/iframe.tmpl")
	}
	if err != nil {
		log.Fatal(err)
	}
}

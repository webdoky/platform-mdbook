package xulelem

import (
	"errors"
	"html/template"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func parseXulelemArgs(args string) (string, error) {
	argSlice := strings.Split(args, ",")
	if len(argSlice) < 1 {
		return "", errors.New("no arguments")
	}
	elementName := helpers.UnwrapString(argSlice[0])
	return elementName, nil
}

func Xulelem(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	elementName, err := parseXulelemArgs(args)
	if err != nil {
		return "", err
	}
	url := "/" + env.Locale + "/docs/Mozilla/Tech/XUL/" + elementName
	link, err := renderhtml.RenderA(&renderhtml.AParams{
		Href: url,
		Text: "<xul:" + elementName + ">",
	})
	if err != nil {
		return "", err
	}
	return renderhtml.RenderCode(&renderhtml.CodeParams{
		InnerHtml: template.HTML(link),
	})
}

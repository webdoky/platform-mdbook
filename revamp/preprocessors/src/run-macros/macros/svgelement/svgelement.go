package svgelement

import (
	"errors"
	"html/template"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func parseSvgelementArgs(args string) (string, error) {
	argsSlice := strings.Split(args, ",")
	if len(argsSlice) < 1 {
		return "", errors.New("not enough arguments")
	}
	return helpers.UnwrapString(argsSlice[0]), nil
}

func Svgelement(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	term, err := parseSvgelementArgs(args)
	if err != nil {
		return "", err
	}
	basePath := "/" + env.Locale + "/docs/Web/SVG/Element/"
	url := basePath + term
	title := "<" + term + ">"
	wrappedTitle, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: title,
	})
	if err != nil {
		return "", err
	}
	return renderhtml.RenderA(&renderhtml.AParams{
		Href:      url,
		InnerHtml: template.HTML(wrappedTitle),
	})
}

package compat

import (
	"errors"
	"html/template"
	"strings"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func Compat(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	compatItems := []string{}
	queries := env.Frontmatter.BrowserCompat
	if len(queries) == 0 {
		return "", errors.New("no browser compatibility queries found in frontmatter")
	}
	for _, query := range env.Frontmatter.BrowserCompat {
		isMultiple := len(queries) > 1
		multiple := ""
		if isMultiple {
			multiple = "multiple"
		}
		queryElement, err := renderhtml.RenderDiv(&renderhtml.DivParams{
			Class: "bc-data",
			Data: map[string]template.HTMLAttr{
				"multiple": template.HTMLAttr(multiple),
				"query":    template.HTMLAttr(query),
				"depth":    template.HTMLAttr("1"),
			},
			Text: "If you're able to see this, something went wrong on this page.",
		})
		if err != nil {
			return "", err
		}
		compatItems = append(compatItems, queryElement)
	}
	return strings.Join(compatItems, "\n"), nil
}

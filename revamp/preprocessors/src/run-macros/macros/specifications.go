package macros

import (
	"errors"
	"html/template"
	"strings"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func specifications(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	compatItems := []string{}
	queries := env.Frontmatter.BrowserCompat
	specUrls := env.Frontmatter.SpecUrls
	if len(queries) == 0 {
		return "", errors.New("no browser compatibility queries found in frontmatter")
	}
	if len(specUrls) == 0 {
		return "", errors.New("no spec urls found in frontmatter")
	}
	query := strings.Join(queries, ",")
	specUrl := strings.Join(specUrls, ",")

	queryElement, err := renderhtml.RenderDiv(&renderhtml.DivParams{
		Class: "bc-specs",
		Data: map[template.HTMLAttr]template.HTMLAttr{
			template.HTMLAttr("bcd-query"): template.HTMLAttr(query),
			template.HTMLAttr("spec-urls"): template.HTMLAttr(specUrl),
		},
		Text: "If you're able to see this, something went wrong on this page.",
	})
	if err != nil {
		return "", err
	}
	compatItems = append(compatItems, queryElement)
	return strings.Join(compatItems, "\n"), nil
}

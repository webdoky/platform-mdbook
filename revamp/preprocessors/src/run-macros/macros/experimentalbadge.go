package macros

import (
	"html/template"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func experimentalbadge(env *environment.Environment, reg *registry.Registry, _ string) (string, error) {
	// span, err := renderhtml.RenderSpan(&renderhtml.SpanParams{
	// 	Class: template.HTMLAttr("visually-hidden"),
	// 	Text:  "Експериментальне",
	// })
	// if err != nil {
	// 	return "", err
	// }
	return renderhtml.RenderAbbr(&renderhtml.AbbrParams{
		Class: template.HTMLAttr("icon icon-experimental"),
		// InnerHtml: template.HTML(span),
		Text:  "Експериментальне",
		Title: template.HTMLAttr("Експериментальне. Поведінка цієї можливості в майбутньому може змінитися."),
	})
}

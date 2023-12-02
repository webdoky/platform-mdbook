package cssinfo

import (
	"html/template"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
)

func related_at_rule(env *environment.Environment, atRuleName string) (string, error) {
	if atRuleName == "" {
		return "", nil
	}
	code, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: atRuleName,
	})
	if err != nil {
		return "", err
	}
	link, err := renderhtml.RenderA(&renderhtml.AParams{
		Href:      "/" + env.Locale + "/docs/Web/CSS/" + atRuleName,
		InnerHtml: template.HTML(code),
	})
	if err != nil {
		return "", err
	}
	return link, nil
}

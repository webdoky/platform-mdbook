package css_l10n

import (
	"html/template"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func angleRoundedToNextQuarter(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	angleWrap, err := renderhtml.RenderCode(&renderhtml.CodeParams{
		Text: "<angle>",
	})
	if err != nil {
		return "", err
	}
	xref_cssangle, err := renderhtml.RenderA(&renderhtml.AParams{
		Href:      "/" + env.Locale + "/docs/Web/CSS/angle",
		InnerHtml: template.HTML(angleWrap),
	})
	if err != nil {
		return "", err
	}
	zeroWrap, err := renderhtml.RenderCode(&renderhtml.CodeParams{
		Text: "0deg",
	})
	if err != nil {
		return "", err
	}
	oneTurnWrap, err := renderhtml.RenderCode(&renderhtml.CodeParams{
		Text: "1turn",
	})
	if err != nil {
		return "", err
	}
	return xref_cssangle + ", округлений до наступної від " + zeroWrap + " чверті та нормалізований, тобто остача від ділення на " + oneTurnWrap + " відсутня", nil
}

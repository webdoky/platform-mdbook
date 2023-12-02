package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func allElementsAndPseudos(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	beforeRef, err := cssxref.Cssxref(env, reg, "::before")
	if err != nil {
		return "", err
	}
	afterRef, err := cssxref.Cssxref(env, reg, "::after")
	if err != nil {
		return "", err
	}
	pseudoElementsLink, err := renderhtml.RenderA(&renderhtml.AParams{
		Text: "псевдоелементи",
		Href: "/" + env.Locale + "/docs/Web/CSS/Pseudo-elements",
	})
	if err != nil {
		return "", err
	}
	return "всі елементи, а також " + pseudoElementsLink + " " + beforeRef + " й " + afterRef, nil
}

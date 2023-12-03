package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func allElementsNoEffectIfDisplayNone(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	displayRef, err := cssxref.Cssxref(env, reg, "display")
	if err != nil {
		return "", err
	}
	noneWrap, err := renderhtml.RenderCode(&renderhtml.CodeParams{
		Text: "none",
	})
	if err != nil {
		return "", err
	}
	return "всі елементи, але не діє, якщо значенням " + displayRef + " є " + noneWrap, nil
}

package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func absoluteLengthOrNormal(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	normalWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "normal",
	})
	if err != nil {
		return "", err
	}
	return "абсолютна довжина або ключове слово " + normalWrap, nil
}

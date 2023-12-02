package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func absoluteLengthOrNone(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	noneWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "none",
	})
	if err != nil {
		return "", err
	}
	return "абсолютна довжина або " + noneWrap, nil
}

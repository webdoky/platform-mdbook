package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func absoluteLength0ForNone(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	noneWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "none",
	})
	if err != nil {
		return "", err
	}
	zeroWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "0",
	})
	if err != nil {
		return "", err
	}
	return "абсолютна довжина; якщо задано ключове слово " + noneWrap + ", то обчисленим значенням є " + zeroWrap, nil
}

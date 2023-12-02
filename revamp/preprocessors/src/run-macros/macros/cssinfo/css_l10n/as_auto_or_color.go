package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func asAutoOrColor(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	autoWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "auto",
	})
	if err != nil {
		return "", err
	}
	colorWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "<color>",
	})
	if err != nil {
		return "", err
	}
	colorRef, err := cssxref.Cssxref(env, reg, "color")
	if err != nil {
		return "", err
	}
	return autoWrap + " обчислюється як задано, а " + colorWrap + " обчислюється як визначено для властивості " + colorRef + ".", nil
}

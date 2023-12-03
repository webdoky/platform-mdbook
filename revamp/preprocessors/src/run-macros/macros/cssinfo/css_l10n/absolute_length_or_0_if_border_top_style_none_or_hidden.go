package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func absoluteLengthOr0IfBorderTopStyleNoneOrHidden(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	zeroWrap, err := renderhtml.RenderCode(&renderhtml.CodeParams{
		Text: "0",
	})
	if err != nil {
		return "", err
	}
	ref, err := cssxref.Cssxref(env, reg, "border-top-style")
	if err != nil {
		return "", err
	}
	noneWrap, err := renderhtml.RenderCode(&renderhtml.CodeParams{
		Text: "none",
	})
	if err != nil {
		return "", err
	}
	hiddenWrap, err := renderhtml.RenderCode(&renderhtml.CodeParams{
		Text: "hidden",
	})
	if err != nil {
		return "", err
	}
	return "абсолютна довжина; " + zeroWrap + ", якщо значення властивості " + ref + " – " + noneWrap + " або " + hiddenWrap, nil
}

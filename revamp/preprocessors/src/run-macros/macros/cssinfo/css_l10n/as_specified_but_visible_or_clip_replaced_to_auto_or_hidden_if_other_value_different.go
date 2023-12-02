package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func asSpecifiedButVisibleOrClipReplacedToAutoOrHiddenIfOtherValueDifferent(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	visibleWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "visible",
	})
	if err != nil {
		return "", err
	}
	clipWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "clip",
	})
	if err != nil {
		return "", err
	}
	autoWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "auto",
	})
	if err != nil {
		return "", err
	}
	hiddenWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "hidden",
	})
	if err != nil {
		return "", err
	}
	overflowXRef, err := cssxref.Cssxref(env, reg, "overflow-x")
	if err != nil {
		return "", err
	}
	overflowYRef, err := cssxref.Cssxref(env, reg, "overflow-y")
	if err != nil {
		return "", err
	}
	return "як задано, крім випадків, коли " + visibleWrap + " обчислюється в " + autoWrap + " або " + clipWrap + " обчислюється в " + hiddenWrap + ", якщо або " + overflowXRef + ", або " + overflowYRef + " має значення, відмінне від " + visibleWrap + " і " + clipWrap, nil
}

package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func allElementsExceptTableElementsWhenCollapse(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	borderCollapseWrap, err := cssxref.Cssxref(env, reg, "border-collapse")
	if err != nil {
		return "", err
	}
	collapseWrap, err := renderhtml.RenderCode(&renderhtml.CodeParams{
		Text: "collapse",
	})
	if err != nil {
		return "", err
	}
	return "всі елементи, крім внутрішніх елементів таблиць, коли " + borderCollapseWrap + " має значення " + collapseWrap, nil
}

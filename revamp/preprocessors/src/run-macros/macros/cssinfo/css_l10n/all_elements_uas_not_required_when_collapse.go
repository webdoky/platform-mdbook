package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func allElementsUAsNotRequiredWhenCollapse(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	tableWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "table",
	})
	if err != nil {
		return "", err
	}
	inlineTableWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "inline-table",
	})
	if err != nil {
		return "", err
	}
	borderCollapseRef, err := cssxref.Cssxref(env, reg, "border-collapse")
	if err != nil {
		return "", err
	}
	collapseWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "collapse",
	})
	if err != nil {
		return "", err
	}
	return "усі елементи; проте користувацькі агенти не зобов'язані застосовувати це до " + tableWrap + " і " + inlineTableWrap + ", якщо " + borderCollapseRef + " має значення " + collapseWrap + ". Поведінка внутрішніх табличних елементів наразі невизначена.", nil
}

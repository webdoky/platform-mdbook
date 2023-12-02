package css_l10n

import (
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func allElementsExceptTableDisplayTypes(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	displayRef, err := cssxref.Cssxref(env, reg, "display")
	if err != nil {
		return "", err
	}
	tableCaptionWrap, err := renderhtml.RenderCode(&renderhtml.WrapperParams{
		Text: "table-caption",
	})
	if err != nil {
		return "", err
	}
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
	return "всі елементи, крім елементів з табличними типами " + displayRef + ", крім " + tableCaptionWrap + ", " + tableWrap + " та " + inlineTableWrap, nil
}

package cssinfo

import (
	"html/template"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssinfo/css_l10n"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func as_longhands(env *environment.Environment, reg *registry.Registry, values []string) (string, error) {
	asLonghands, err := css_l10n.Localize(env, reg, "asLonghands", "", "")
	if err != nil {
		return "", err
	}
	br, err := renderhtml.RenderBr()
	if err != nil {
		return "", err
	}
	result := asLonghands + br
	listContent := ""
	for _, singleInitial := range values {
		ref, err := cssxref.Cssxref(env, reg, singleInitial)
		if err != nil {
			return "", err
		}
		listItem, err := renderhtml.RenderLi(&renderhtml.LiParams{
			InnerHtml: template.HTML(ref + ": TODO"),
			// (Object.prototype.hasOwnProperty.call(data.properties, longhand) ?
			// await getValueOutput(data.properties[longhand], property) :
			// replacePlaceholders(formattedError,
			// 	[localize(cssLocalStrings, "missing")]))
		})
		if err != nil {
			return "", err
		}
		listContent += listItem
	}
	list, err := renderhtml.RenderUl(&renderhtml.UlParams{
		InnerHtml: template.HTML(listContent),
	})
	if err != nil {
		return "", err
	}
	result += list
	return result, nil
}

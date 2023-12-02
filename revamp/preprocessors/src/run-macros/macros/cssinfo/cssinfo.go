package cssinfo

import (
	"html/template"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers/l10n"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssinfo/css_l10n"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

type propertyItem struct {
	content   template.HTML
	labelHtml template.HTML
	labelText string
}

func Cssinfo(env *environment.Environment, reg *registry.Registry, _ string) (string, error) {
	slug := env.Frontmatter.Slug
	name := slug[strings.LastIndex(slug, "/")+1:]
	var atRuleName string
	cssAtRuleMarkerIndex := strings.Index(slug, "/CSS/@")
	var data *CssData
	var err error
	if cssAtRuleMarkerIndex != -1 {
		atRuleName = slug[cssAtRuleMarkerIndex+len("/CSS/"):]
	}
	if atRuleName == "" {
		data, err = get_mdn_data("properties", name)
		if err != nil {
			return "", err
		}
	} else {
		data, err := get_mdn_data("at-rules", atRuleName)
		if err != nil {
			return "", err
		}
		data = data.Descriptors[name]
	}
	if data == nil {
		return renderhtml.RenderSpan(&renderhtml.SpanParams{
			Style: "color: red;",
			Text:  l10n.L10nCss(env.Locale, "missing"),
		})
	}
	properties := []propertyItem{}
	if atRuleName != "" {
		localizedRelatedAtRule, err := css_l10n.Localize(env, reg, "relatedAtRule", "", "")
		if err != nil {
			return "", err
		}
		if err != nil {
			return "", err
		}
		atRuleContent, err := related_at_rule(env, atRuleName)
		if err != nil {
			return "", err
		}
		properties = append(properties, propertyItem{
			content:   template.HTML(atRuleContent),
			labelHtml: template.HTML(localizedRelatedAtRule),
		})
	}
	initialLink, err := renderhtml.RenderA(&renderhtml.AParams{
		Href: "/" + env.Locale + "/docs/Web/CSS/initial_value",
		Text: "Початкове значення",
	})
	if err != nil {
		return "", err
	}
	initialContent, err := initial(env, reg, data)
	if err != nil {
		return "", err
	}
	properties = append(properties, propertyItem{
		content:   template.HTML(initialContent),
		labelHtml: template.HTML(initialLink),
	})
	appliestoContent, err := appliesto(env, reg, data)
	if err != nil {
		return "", err
	}
	if atRuleName == "" {
		appliesToLabel, err := css_l10n.Localize(env, reg, "appliesTo", "", "")
		if err != nil {
			return "", err
		}
		properties = append(properties, propertyItem{
			content:   template.HTML(appliestoContent),
			labelText: appliesToLabel,
		})
	}
	if data.Inherited {
		inheritedLink, err := renderhtml.RenderA(&renderhtml.AParams{
			Href: "/" + env.Locale + "/docs/Web/CSS/Inheritance",
			Text: "Успадковується",
		})
		if err != nil {
			return "", err
		}
		inheritedContent, err := inherited(env, reg, data)
		if err != nil {
			return "", err
		}
		properties = append(properties, propertyItem{
			content:   template.HTML(inheritedContent),
			labelHtml: template.HTML(inheritedLink),
		})
	}
	if data.Percentages == nil || data.Percentages[0] != "no" {
		percentagesContent, err := percentages(env, reg, data)
		if err != nil {
			return "", err
		}
		percentagesLabel, err := css_l10n.Localize(env, reg, "percentages", "", "")
		if err != nil {
			return "", err
		}
		properties = append(properties, propertyItem{
			content:   template.HTML(percentagesContent),
			labelText: percentagesLabel,
		})
	}
	computedLink, err := renderhtml.RenderA(&renderhtml.AParams{
		Href: "/" + env.Locale + "/docs/Web/CSS/computed_value",
		Text: "Обчислене значення",
	})
	if err != nil {
		return "", err
	}
	computedContent, err := computed(env, reg, data)
	if err != nil {
		return "", err
	}
	properties = append(properties, propertyItem{
		content:   template.HTML(computedContent),
		labelHtml: template.HTML(computedLink),
	})
	if atRuleName == "" {
		animationTypeContent, err := animationType(env, reg, data)
		if err != nil {
			return "", err
		}
		properties = append(properties, propertyItem{
			content:   template.HTML(animationTypeContent),
			labelText: "Тип анімування",
		})
	}
	if data.Stacking {
		stackingContent, err := stacking(env, reg, data)
		if err != nil {
			return "", err
		}
		stackingLabel, err := css_l10n.Localize(env, reg, "createsStackingContext", "", "")
		if err != nil {
			return "", err
		}
		properties = append(properties, propertyItem{
			content:   template.HTML(stackingContent),
			labelHtml: template.HTML(stackingLabel),
		})
	}
	rows := []string{}
	for _, property := range properties {
		th, err := renderhtml.RenderTh(&renderhtml.ThParams{
			InnerHtml: property.labelHtml,
			Scope:     "row",
			Text:      property.labelText,
		})
		if err != nil {
			return "", err
		}
		td, err := renderhtml.RenderTd(&renderhtml.TdParams{
			InnerHtml: property.content,
		})
		if err != nil {
			return "", err
		}
		tr, err := renderhtml.RenderTr(&renderhtml.TrParams{
			InnerHtml: template.HTML(th + td),
		})
		if err != nil {
			return "", err
		}
		rows = append(rows, tr)
	}
	tbody, err := renderhtml.RenderTbody(&renderhtml.TbodyParams{
		InnerHtml: template.HTML(strings.Join(rows, "")),
	})
	if err != nil {
		return "", err
	}
	table, err := renderhtml.RenderTable(&renderhtml.TableParams{
		Class:     template.HTMLAttr("properties"),
		InnerHtml: template.HTML(tbody),
	})
	if err != nil {
		return "", err
	}
	return table, nil
}

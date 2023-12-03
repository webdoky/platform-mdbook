package macros

import (
	"html/template"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func parseDeprecated_headerArgs(args string) (string, error) {
	argSlice := strings.Split(args, ",")
	var version string
	if len(argSlice) >= 1 {
		version = helpers.UnwrapString(argSlice[0])
	}
	// if version starts with a digit, prepend "gecko " to it
	if len(version) > 0 && version[0] >= '0' && version[0] <= '9' {
		version = "gecko " + version
	}
	return version, nil
}

func deprecated_header(environment *environment.Environment, registry *registry.Registry, args string) (string, error) {
	note, err := parseDeprecated_headerArgs(args)
	if err != nil {
		return "", err
	}
	if note == "" {
		note = "Нерекомендоване"
	}
	h4Html, err := renderhtml.RenderH4(&renderhtml.WrapperParams{Text: note})
	if err != nil {
		return "", err
	}
	compatibilityAnchorHtml, err := renderhtml.RenderA(&renderhtml.AParams{Href: "#sumisnist-iz-brauzeramy", Text: "таблицю сумісності"})
	if err != nil {
		return "", err
	}
	pHtml, err := renderhtml.RenderP(&renderhtml.PParams{InnerHtml: template.HTML("Ця можливість більше не є рекомендованою. Попри те, що деякі браузери досі можуть її підтримувати, можливо, вона вже вилучена з відповідних вебстандартів, або перебуває в процесі викидання, або ж збережена суто з міркувань сумісності. Слід уникати її використання та прибирати з наявного коду, коли це можливо; у прийнятті свого рішення керуйтесь " + compatibilityAnchorHtml + " внизу цієї сторінки. Майте на увазі, що ця можливість може в будь-який час припинити працювати.")})
	if err != nil {
		return "", err
	}
	html, err := renderhtml.RenderDiv(&renderhtml.DivParams{Class: "notecard deprecated", InnerHtml: template.HTML(h4Html + pHtml)})
	if err != nil {
		return "", err
	}
	return html, nil
}

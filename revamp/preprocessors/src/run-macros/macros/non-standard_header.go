package macros

import (
	"html/template"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func non_standard_header(environment *environment.Environment, registry *registry.Registry, args string) (string, error) {
	h4Html, err := renderhtml.RenderH4(&renderhtml.WrapperParams{Text: "Нестандартне"})
	if err != nil {
		return "", err
	}
	pHtml, err := renderhtml.RenderP(&renderhtml.PParams{Text: "Ця можливість не є стандартною та не планується для додавання в жодний стандарт. Не використовуйте її на публічних вебсайтах: вона працюватиме не для всіх користувачів. Крім цього, між реалізаціями можуть бути суттєві відмінності, а ще можуть бути зміни в майбутньому."})
	if err != nil {
		return "", err
	}
	html, err := renderhtml.RenderDiv(&renderhtml.DivParams{Class: "notecard nonstandard", InnerHtml: template.HTML(h4Html + pHtml)})
	if err != nil {
		return "", err
	}
	return html, nil
}

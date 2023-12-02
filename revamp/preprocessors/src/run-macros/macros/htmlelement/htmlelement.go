package htmlelement

import (
	"errors"
	"html/template"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

var SECTION_NAME = "Element"

func parseHtmlelementArgs(args string) (string, string, string, error) {
	// Split the args string into a slice of strings
	// using the comma as the separator
	// (e.g., "termName,displayName" -> ["termName", "displayName"])
	argSlice := strings.Split(args, ",")
	var elementName, displayName, anchor string
	if len(argSlice) == 0 {
		return "", "", "", errors.New("no arguments")
	}
	switch len(argSlice) {
	case 0:
		return "", "", "", errors.New("no arguments")
	case 3:
		anchor = strings.TrimPrefix(helpers.UnwrapString(argSlice[2]), "#")
		fallthrough
	case 2:
		displayName = helpers.UnwrapString(argSlice[1])
		fallthrough
	case 1:
		elementName = strings.ToLower(helpers.UnwrapString(argSlice[0]))
	default:
		return "", "", "", errors.New("too many arguments")
	}
	if displayName == "" {
		displayName = strings.ToLower(elementName)
	}
	return elementName, displayName, anchor, nil
}

func Htmlelement(env *environment.Environment, _ *registry.Registry, args string) (string, error) {
	elementName, displayName, anchor, err := parseHtmlelementArgs(args)
	if err != nil {
		return "", err
	}
	baseUrl := "/" + env.Locale + "/docs/Web/HTML"
	urlPath := "/" + SECTION_NAME + "/" + elementName
	href := baseUrl + urlPath
	if anchor != "" {
		href += "#" + anchor
	}
	aParams := renderhtml.AParams{
		Href: href,
	}
	if displayName == elementName && !strings.Contains(elementName, " ") {
		aParams.InnerHtml = template.HTML(helpers.WrapAsCode("<" + displayName + ">"))
	} else {
		aParams.Text = displayName
	}
	return renderhtml.RenderA(&aParams)
}

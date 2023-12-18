package macros

import (
	"errors"
	"html/template"
	"strings"
	preprocessor_helpers "webdoky3/revamp/preprocessors/src/helpers"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func parseHttpheaderArgs(args string) (string, string, string, bool, error) {
	// Split the args string into a slice of strings
	// using the comma as the separator
	// (e.g., "termName,displayName" -> ["termName", "displayName"])
	argSlice := strings.Split(args, ",")
	var header, displayName, anchor string
	var ignoreWrap bool
	if len(argSlice) == 0 {
		return "", "", "", false, errors.New("no arguments")
	}
	header = preprocessor_helpers.UnwrapString(argSlice[0])
	if len(argSlice) >= 2 {
		displayName = preprocessor_helpers.UnwrapString(argSlice[1])
	}
	if displayName == "" {
		displayName = header
	}
	if len(argSlice) >= 3 {
		anchor = preprocessor_helpers.UnwrapString(argSlice[2])
		anchor = strings.TrimPrefix(anchor, "#")
	}
	if len(argSlice) >= 4 {
		ignoreWrap = preprocessor_helpers.UnwrapBoolean(argSlice[3])
	}
	return header, displayName, anchor, ignoreWrap, nil
}

func httpheader(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	header, displayName, anchor, ignoreWrap, err := parseHttpheaderArgs(args)
	if err != nil {
		return "", err
	}
	url := "/" + env.Locale + "/docs/Web/HTTP/Headers/" + header
	if anchor != "" {
		displayName += "." + anchor
		url += "#" + anchor
	}
	aParams := renderhtml.AParams{
		Href: url,
	}
	if ignoreWrap {
		aParams.Text = displayName
	} else {
		aParams.InnerHtml = template.HTML(preprocessor_helpers.WrapAsCode(displayName))
	}
	aHtml, err := renderhtml.RenderA(&aParams)
	if err != nil {
		return "", err
	}
	return aHtml, err
}

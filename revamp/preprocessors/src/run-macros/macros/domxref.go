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

func parseDomxrefArgs(env *environment.Environment, args string) (string, string, string, bool, error) {
	// Split the args string into a slice of strings
	// using the comma as the separator
	// (e.g., "termName,displayName" -> ["termName", "displayName"])
	argSlice := strings.Split(args, ",")
	var apiName, displayName, anchor string
	var ignoreWrap bool
	if len(argSlice) == 0 {
		return "", "", "", false, errors.New("no arguments")
	}
	switch len(argSlice) {
	case 0:
		return "", "", "", false, errors.New("no arguments")
	case 4:
		ignoreWrap = preprocessor_helpers.UnwrapBoolean(argSlice[3])
		fallthrough
	case 3:
		anchor = strings.TrimPrefix(preprocessor_helpers.UnwrapString(argSlice[2]), "#")
		fallthrough
	case 2:
		displayName = preprocessor_helpers.UnwrapString(argSlice[1])
		fallthrough
	case 1:
		apiName = preprocessor_helpers.UnwrapString(argSlice[0])
	default:
		return "", "", "", false, errors.New("too many arguments")
	}
	if displayName == "" {
		displayName = apiName
	}
	if anchor != "" {
		displayName += "." + anchor
	}
	apiName = strings.ReplaceAll(apiName, " ", "_")
	apiName = strings.ReplaceAll(apiName, "()", "")
	apiName = strings.ReplaceAll(apiName, ".prototype.", ".")
	apiName = strings.ReplaceAll(apiName, ".", "/")
	// Capitalize apiName
	apiName = strings.ToUpper(apiName[0:1]) + apiName[1:]
	return apiName, displayName, anchor, ignoreWrap, nil
}

func domxref(env *environment.Environment, _ *registry.Registry, args string) (string, error) {
	apiName, displayName, anchor, ignoreWrap, err := parseDomxrefArgs(env, args)
	if err != nil {
		return "", err
	}
	basePath := "/" + env.Locale + "/docs/Web/API/"
	href := basePath + apiName
	if anchor != "" {
		href += "#" + anchor
	}
	aParams := renderhtml.AParams{
		Href: href,
	}
	if ignoreWrap {
		aParams.Text = displayName
	} else {
		aParams.InnerHtml = template.HTML(preprocessor_helpers.WrapAsCode(displayName))
	}
	return renderhtml.RenderA(&aParams)
}

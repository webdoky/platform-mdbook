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

const BASE_SLUG = "Web/JavaScript/Reference/"
const GLOBAL_OBJECTS = "Global_Objects"

func parseJsxrefArgs(args string) (string, string, string, bool, error) {
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
	return apiName, displayName, anchor, ignoreWrap, nil
}

func jsxref(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	apiName, displayName, anchor, ignoreWrap, err := parseJsxrefArgs(args)
	if err != nil {
		return "", err
	}
	slug := strings.Replace(apiName, "()", "", -1)
	localUrl := "/" + env.Locale + "/docs/" + BASE_SLUG
	slug = strings.Replace(slug, ".prototype.", ".", -1)
	if strings.Contains(apiName, ".") && !strings.Contains(apiName, "..") {
		// E.g. "Array.filter", but not "try...catch".
		slug = strings.Replace(slug, ".", "/", -1)
	}
	var basePath string
	if reg.HasPath(env.Locale + "/docs/" + BASE_SLUG + slug) {
		basePath = localUrl
	} else if reg.HasPath(env.Locale + "/docs/" + BASE_SLUG + GLOBAL_OBJECTS + "/" + slug) {
		basePath = localUrl + GLOBAL_OBJECTS + "/"
	} else {
		basePath = localUrl
	}
	href := basePath + slug
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
	aHtml, err := renderhtml.RenderA(&aParams)
	if err != nil {
		return "", err
	}
	return aHtml, err
}

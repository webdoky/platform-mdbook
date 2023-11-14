package macros

import (
	"errors"
	"log"
	"strings"
	"webdoky3/preprocessors/src/helpers"
)

const BASE_SLUG = "Web/JavaScript/Reference/"
const GLOBAL_OBJECTS = "Global_Objects"

func parseJsxrefArgs(args string) (string, string, string, error) {
	// Split the args string into a slice of strings
	// using the comma as the separator
	// (e.g., "termName,displayName" -> ["termName", "displayName"])
	argSlice := strings.Split(args, ",")
	var apiName, displayName, anchor string
	if len(argSlice) == 0 {
		return "", "", "", errors.New("no arguments")
	}
	switch len(argSlice) {
	case 0:
		return "", "", "", errors.New("no arguments")
	case 4:
		ignoreWrap := helpers.UnwrapBoolean(argSlice[3])
		if ignoreWrap {
			displayName = helpers.UnwrapString(argSlice[1])
			if displayName == "" {
				displayName = apiName
			}
		}
		fallthrough
	case 3:
		anchor = strings.TrimPrefix(helpers.UnwrapString(argSlice[2]), "#")
		fallthrough
	case 2:
		if displayName == "" {
			displayName = "<code>" + helpers.UnwrapString(argSlice[1]) + "</code>"
		}
		fallthrough
	case 1:
		apiName = helpers.UnwrapString(argSlice[0])
		if displayName == "" {
			displayName = "<code>" + apiName + "</code>"
		}
	default:
		return "", "", "", errors.New("too many arguments")
	}
	return apiName, displayName, anchor, nil
}

func jsxref(env *Environment, registry Registry, args string) (string, error) {
	apiName, displayName, anchor, err := parseJsxrefArgs(args)
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
	if registry.HasPath(env.Locale + "/docs/" + BASE_SLUG + slug) {
		basePath = localUrl
	} else if registry.HasPath(env.Locale + "/docs/" + BASE_SLUG + GLOBAL_OBJECTS + "/" + slug) {
		basePath = localUrl + GLOBAL_OBJECTS + "/"
	} else {
		basePath = localUrl
	}
	log.Printf("basePath: %s", basePath)
	href := basePath + slug
	log.Printf("href: %s", href)
	if anchor != "" {
		href += "#" + anchor
	}
	return "<a href=\"" + href + "\">" + displayName + "</a>", nil
}

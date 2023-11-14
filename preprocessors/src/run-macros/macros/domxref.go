package macros

import (
	"errors"
	"strings"
	"webdoky3/preprocessors/src/helpers"

	"golang.org/x/exp/slices"
)

var RTL_LOCALES = []string{"ar", "fa", "he"}

func parseDomxrefArgs(env *Environment, args string) (string, string, string, error) {
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
		if displayName == "" {
			displayName = helpers.UnwrapString(argSlice[1])
			if displayName == "" {
				if apiName == "" {
					displayName = "<code>" + helpers.UnwrapString(argSlice[0]) + "</code>"
				} else {
					displayName = "<code>" + apiName + "</code>"
				}
			}
		}
		displayName += "." + anchor
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
	apiName = strings.ReplaceAll(apiName, " ", "_")
	apiName = strings.ReplaceAll(apiName, "()", "")
	apiName = strings.ReplaceAll(apiName, ".prototype.", ".")
	apiName = strings.ReplaceAll(apiName, ".", "/")
	if slices.Contains(RTL_LOCALES, strings.ToLower(env.Locale)) {
		displayName = "<bdi>" + displayName + "</bdi>"
	}
	// Capitalize apiName
	apiName = strings.ToUpper(apiName[0:1]) + apiName[1:]
	return apiName, displayName, anchor, nil
}

func domxref(env *Environment, registry Registry, args string) (string, error) {
	apiName, displayName, anchor, err := parseDomxrefArgs(env, args)
	if err != nil {
		return "", err
	}
	basePath := "/" + env.Locale + "/docs/Web/API/"
	href := basePath + apiName
	if anchor != "" {
		href += "#" + anchor
	}
	return "<a href=\"" + href + "\">" + displayName + "</a>", nil
}

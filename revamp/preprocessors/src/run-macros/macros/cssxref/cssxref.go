package cssxref

import (
	"errors"
	"html"
	"html/template"
	"regexp"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

var ANGLES_TYPE_SLUG_REGEX = regexp.MustCompile(`<(.*)>`)

func parseCssxrefArgs(args string) (string, string, string, error) {
	// Split the args string into a slice of strings
	// using the comma as the separator
	// (e.g., "termName,displayName" -> ["termName", "displayName"])
	argSlice := strings.Split(args, ",")
	var slug, displayName, anchor string
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
		slug = helpers.UnwrapString(argSlice[0])
	default:
		return "", "", "", errors.New("too many arguments")
	}
	return slug, displayName, anchor, nil
}

func Cssxref(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	slug, displayName, anchor, err := parseCssxrefArgs(args)
	if err != nil {
		return "", err
	}
	slug = html.UnescapeString(slug)
	switch slug {
	case "<color>":
		slug = "color_value"
	case "<flex>":
		slug = "flex_value"
	case "<position>":
		slug = "position_value"
	default:
		slug = ANGLES_TYPE_SLUG_REGEX.ReplaceAllString(slug, "$1")
	}
	basePath := "/" + env.Locale + "/docs/Web/CSS/"
	urlWithoutAnchor := basePath + slug
	url := urlWithoutAnchor
	if anchor != "" {
		url += "#" + anchor
	}
	if displayName == "" {
		frontmatterData, err := helpers.GetFrontmatterDataBySlug("web/css/"+slug, env.Locale)
		if err != nil {
			return "", err
		}
		if frontmatterData != nil {
			switch frontmatterData.PageType {
			case "css-function":
				displayName = slug + "()"
			case "css-type":
				if !strings.HasPrefix(slug, "<") && !strings.HasSuffix(slug, ">") {
					displayName = "<" + slug + ">"
				}
			default:
				displayName = slug
			}
		} else {
			displayName = slug
		}
	}
	aParams := renderhtml.AParams{
		Href:      url,
		InnerHtml: template.HTML(helpers.WrapAsCode(displayName)),
	}
	aHtml, err := renderhtml.RenderA(&aParams)
	if err != nil {
		return "", err
	}
	return aHtml, err
}

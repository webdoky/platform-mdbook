package macros

import (
	"errors"
	"html/template"
	"strings"
	preprocessor_helpers "webdoky3/revamp/preprocessors/src/helpers"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"

	"golang.org/x/exp/slices"
)

var supportedHeights = []string{"shorter", "taller", "tabbed-shorter", "tabbed-standard", "tabbed-taller"}

func parseEmbedinteractiveexampleArgs(args string) (string, string, error) {
	argSlice := strings.Split(args, ",")
	var url, height string
	switch len(argSlice) {
	case 0:
		return "", "", errors.New("no arguments")
	case 1:
		url = preprocessor_helpers.UnwrapString(argSlice[0])
	case 2:
		url = preprocessor_helpers.UnwrapString(argSlice[0])
		height = preprocessor_helpers.UnwrapString(argSlice[1])
		if !slices.Contains(supportedHeights, height) {
			return "", "", errors.New("invalid height")
		}
	default:
		return "", "", errors.New("too many arguments")
	}
	return url, height, nil
}

func embedinteractiveexample(env *environment.Environment, reg *registry.Registry, args string) (string, error) {
	url, height, err := parseEmbedinteractiveexampleArgs(args)
	if err != nil {
		return "", err
	}
	heightClass := "is-default-height"
	if strings.Contains(url, "/js/") {
		heightClass = "is-js-height"
	}
	if height != "" {
		heightClass = "is-" + height + "-height"
	}
	h2, err := renderhtml.RenderH2(&renderhtml.WrapperParams{
		Text: "Інтерактивний приклад",
	})
	if err != nil {
		return "", err
	}
	iframe, err := renderhtml.RenderIframe(&renderhtml.IframeParams{
		Class:  template.HTMLAttr("interactive " + heightClass),
		Height: "200",
		Src:    template.HTMLAttr("/interactive-examples/" + url),
		Title:  template.HTMLAttr("Інтерактивний приклад ВебДоків"),
	})
	if err != nil {
		return "", err
	}
	return h2 + iframe, nil
}

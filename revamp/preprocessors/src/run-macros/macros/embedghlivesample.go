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

func parseEmbedghlivesampleArgs(args string) (string, string, string, error) {
	// args: "path,width,height"
	// Split the args string into a slice of strings
	// using the comma as the separator
	// (e.g., "path,width,height" -> ["path", "width", "height"])
	argSlice := strings.Split(args, ",")
	switch len(argSlice) {
	case 0:
		return "", "", "", errors.New("no arguments")
	case 1:
		return preprocessor_helpers.UnwrapString(argSlice[0]), "", "", nil
	case 2:
		return preprocessor_helpers.UnwrapString(argSlice[0]), preprocessor_helpers.UnwrapString(argSlice[1]), "", nil
	case 3:
		return preprocessor_helpers.UnwrapString(argSlice[0]), preprocessor_helpers.UnwrapString(argSlice[1]), preprocessor_helpers.UnwrapString(argSlice[2]), nil
	default:
		return "", "", "", errors.New("too many arguments")
	}
}

func embedghlivesample(env *environment.Environment, _ *registry.Registry, args string) (string, error) {
	path, width, height, err := parseEmbedghlivesampleArgs(args)
	if err != nil {
		return "", err
	}
	return renderhtml.RenderIframe(&renderhtml.IframeParams{
		Height: template.HTMLAttr(height),
		Src:    template.HTMLAttr("https://webdoky.github.io/" + path),
		Width:  template.HTMLAttr(width),
	})

}

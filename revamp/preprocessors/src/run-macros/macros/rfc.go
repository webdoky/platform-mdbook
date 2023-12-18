package macros

import (
	"errors"
	"strings"
	preprocessor_helpers "webdoky3/revamp/preprocessors/src/helpers"
	"webdoky3/revamp/preprocessors/src/helpers/l10n"
	renderhtml "webdoky3/revamp/preprocessors/src/helpers/render_html"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func parseRfcArgs(args string) (string, string, string, error) {
	argSlice := strings.Split(args, ",")
	var id, text, section string
	if len(argSlice) == 0 {
		return "", "", "", errors.New("no arguments")
	}
	if len(argSlice) >= 1 {
		id = preprocessor_helpers.UnwrapString(argSlice[0])
	}
	if len(argSlice) >= 2 {
		text = ": " + preprocessor_helpers.UnwrapString(argSlice[1])
	}
	if len(argSlice) >= 3 {
		section = preprocessor_helpers.UnwrapString(argSlice[2])
	}
	return id, text, section, nil
}

func rfc(env *environment.Environment, registry *registry.Registry, args string) (string, error) {
	var id, text, section string
	var err error
	id, text, section, err = parseRfcArgs(args)
	if err != nil {
		return "", err
	}
	link := "https://datatracker.ietf.org/doc/html/rfc" + id
	if section != "" {
		link += "#section-" + section
		text = ", " + l10n.L10nCommon(env.Locale, "section") + " " + section + text
	}
	aParams := renderhtml.AParams{
		Href: link,
		Text: "RFC " + id + text,
	}
	aHtml, err := renderhtml.RenderA(&aParams)
	if err != nil {
		return "", err
	}
	return aHtml, nil
}

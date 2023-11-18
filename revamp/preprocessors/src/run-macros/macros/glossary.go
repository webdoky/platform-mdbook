package macros

import (
	"errors"
	"log"
	"strings"
	"webdoky3/revamp/preprocessors/src/helpers"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func parseGlossaryArgs(args string) (string, string, error) {
	// Split the args string into a slice of strings
	// using the comma as the separator
	// (e.g., "termName,displayName" -> ["termName", "displayName"])
	argSlice := strings.Split(args, ",")
	// If the args string is empty, return an empty slice
	for i, arg := range argSlice {
		argSlice[i] = helpers.UnwrapString(arg)
	}
	if len(argSlice) == 1 {
		return argSlice[0], "", nil
	}
	// If the args string has more than two elements,
	// return an error
	if len(argSlice) > 2 {
		return "", "", errors.New("too many arguments")
	}
	// If the args string has two elements,
	// return the slice of strings as a slice of interfaces
	if len(argSlice) == 2 {
		return argSlice[0], argSlice[1], nil
	}
	// If the args string has one element,
	// return the slice of strings as a
	// slice of interfaces with the second element empty
	return "", "", nil
}

func glossary(env *environment.Environment, _ *registry.Registry, args string) (string, error) {
	termName, displayName, err := parseGlossaryArgs(args)
	log.Printf("glossary(%s, %s)", termName, displayName)
	if err != nil {
		return "", err
	}
	basePath := "/" + env.Locale + "/docs/Glossary/"
	// Replace space characters with underscores
	subPath := strings.Replace(termName, " ", "_", -1)
	if displayName == "" {
		displayName = termName
	}
	return "[" + displayName + "](" + basePath + subPath + ")", nil
}

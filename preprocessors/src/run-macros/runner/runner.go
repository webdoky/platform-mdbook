package runner

import (
	"log"
	"regexp"
	"strings"
	"webdoky3/preprocessors/src/helpers"
	"webdoky3/preprocessors/src/run-macros/macros"
)

var MACRO_REGEXP = regexp.MustCompile(`{{(\w+)(?:\(([^{}]+)\))?}}`)

type MacrosRunner struct {
	environment *macros.Environment
	registry    macros.Registry
}

func (mr *MacrosRunner) lookupMacro(macroName string) func(*macros.Environment, macros.Registry, string) (string, error) {
	return macros.MacrosIndex[strings.ToLower(macroName)]
}

func (mr *MacrosRunner) Run(markdownCode string) string {

	return helpers.ReplaceAllStringSubmatchFunc(MACRO_REGEXP, markdownCode, func(match []string) string {
		macroName := match[1]
		macroArgs := match[2]
		macro := mr.lookupMacro(macroName)
		if macro == nil {
			log.Printf("Unknown macro %s", macroName)
			return "<span class=\"wd-expunged\" title=\"Ця частина функціональності ще не готова\">" + match[0] + "</span>"
		}
		macroResult, err := macro(mr.environment, mr.registry, macroArgs)
		if err != nil {
			log.Printf("Error running macro %s: %s", macroName, err)
			return "<span class=\"wd-failed-macro\" title=\"Щось тут пішло не так\">" + match[0] + "</span>"
		}
		return macroResult
	})
}

func NewMacrosRunner(environment *macros.Environment, registry macros.Registry) *MacrosRunner {
	return &MacrosRunner{
		environment: environment,
		registry:    registry,
	}
}

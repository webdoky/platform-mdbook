package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func referToLineHeight(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	lineHeightRef, err := cssxref.Cssxref(env, reg, "line-height")
	if err != nil {
		return "", err
	}
	return "стосується " + lineHeightRef + " самого елемента", nil
}

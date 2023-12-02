package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func elementsWithDisplayBoxOrInlineBox(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	displayRef, err := cssxref.Cssxref(env, reg, "display")
	if err != nil {
		return "", err
	}
	return "елементи, чиє значення " + displayRef + " – <code>box</code> або <code>inline-box</code>", nil
}

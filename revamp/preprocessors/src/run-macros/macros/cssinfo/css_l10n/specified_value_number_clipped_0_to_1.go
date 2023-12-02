package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func specifiedValueNumberClipped0To1(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	numberRef, err := cssxref.Cssxref(env, reg, "number")
	if err != nil {
		return "", err
	}
	return "Таке ж, як задане значення після застискання " + numberRef + " в діапазон [0.0, 1.0]", nil
}

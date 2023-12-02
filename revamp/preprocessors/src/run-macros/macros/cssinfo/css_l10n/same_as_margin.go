package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func sameAsMargin(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	marginRef, err := cssxref.Cssxref(env, reg, "margin")
	if err != nil {
		return "", err
	}
	return "те саме, що й " + marginRef, nil
}

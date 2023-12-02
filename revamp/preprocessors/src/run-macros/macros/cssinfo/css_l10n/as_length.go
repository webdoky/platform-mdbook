package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func asLength(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	lengthRef, err := cssxref.Cssxref(env, reg, "length")
	if err != nil {
		return "", err
	}
	return "як " + lengthRef, nil
}

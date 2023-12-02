package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func absoluteLength(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	ref, err := cssxref.Cssxref(env, reg, "length, довжина")
	if err != nil {
		return "", err
	}
	return "абсолютна " + ref, nil
}

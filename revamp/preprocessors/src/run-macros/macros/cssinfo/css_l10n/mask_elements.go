package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/svgelement"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func maskElements(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	maskRef, err := svgelement.Svgelement(env, reg, "mask")
	if err != nil {
		return "", err
	}
	return "елементи " + maskRef, nil
}

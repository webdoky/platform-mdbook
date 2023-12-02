package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func sameAsWidthAndHeight(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	widthRef, err := cssxref.Cssxref(env, reg, "width")
	if err != nil {
		return "", err
	}
	heightRef, err := cssxref.Cssxref(env, reg, "height")
	if err != nil {
		return "", err
	}
	return "те саме, що й " + widthRef + " і " + heightRef, nil
}

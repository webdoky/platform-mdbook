package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func sameAsMaxWidthAndMaxHeight(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	maxWidthRef, err := cssxref.Cssxref(env, reg, "max-width")
	if err != nil {
		return "", err
	}
	maxHeightRef, err := cssxref.Cssxref(env, reg, "max-height")
	if err != nil {
		return "", err
	}
	return "те саме, що й " + maxWidthRef + " і " + maxHeightRef, nil
}

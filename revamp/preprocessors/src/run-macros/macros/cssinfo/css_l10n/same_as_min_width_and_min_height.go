package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func sameAsMinWidthAndMinHeight(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	minWidthRef, err := cssxref.Cssxref(env, reg, "min-width")
	if err != nil {
		return "", err
	}
	minHeightRef, err := cssxref.Cssxref(env, reg, "min-height")
	if err != nil {
		return "", err
	}
	return "те саме, що й " + minWidthRef + " та " + minHeightRef, nil
}

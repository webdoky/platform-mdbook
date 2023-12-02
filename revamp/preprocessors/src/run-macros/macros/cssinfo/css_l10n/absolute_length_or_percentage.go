package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func absoluteLengthOrPercentage(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	ref, err := cssxref.Cssxref(env, reg, "length")
	if err != nil {
		return "", err
	}
	return "для " + ref + " – абсолютне значення, інакше – відсотки", nil
}

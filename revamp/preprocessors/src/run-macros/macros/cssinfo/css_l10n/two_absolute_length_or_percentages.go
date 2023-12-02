package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func twoAbsoluteLengthOrPercentages(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	lengthRef, err := cssxref.Cssxref(env, reg, "length")
	if err != nil {
		return "", err
	}
	percentageRef, err := cssxref.Cssxref(env, reg, "percentage")
	if err != nil {
		return "", err
	}
	return "два абсолютні значення " + lengthRef + " або " + percentageRef, nil
}

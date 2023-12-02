package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func eachOfShorthandPropertiesExceptUnicodeBiDiAndDirection(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	unicodeBidiRef, err := cssxref.Cssxref(env, reg, "unicode-bidi")
	if err != nil {
		return "", err
	}
	directionRef, err := cssxref.Cssxref(env, reg, "direction")
	if err != nil {
		return "", err
	}
	return "як кожна зі складових скорочення (усі властивості, крім " + unicodeBidiRef + " і " + directionRef + ")", nil
}

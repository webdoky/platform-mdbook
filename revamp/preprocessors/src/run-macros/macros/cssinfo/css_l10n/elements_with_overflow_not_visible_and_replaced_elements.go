package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func elementsWithOverflowNotVisibleAndReplacedElements(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	overflowRef, err := cssxref.Cssxref(env, reg, "overflow")
	if err != nil {
		return "", err
	}
	return "елементи, чиє значення " + overflowRef + " відмінне від <code>visible</code>, і необов'язково – заміщені елементи, що представляють зображення чи відео, а також вбудовані фрейми", nil
}

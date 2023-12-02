package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func firstLetterPseudoElementsAndInlineLevelFirstChildren(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	firstLetterRef, err := cssxref.Cssxref(env, reg, "::first-letter")
	if err != nil {
		return "", err
	}
	return "псевдоелементи " + firstLetterRef + " та перші дочірні елементи рядкового рівня блокових контейнерів", nil
}

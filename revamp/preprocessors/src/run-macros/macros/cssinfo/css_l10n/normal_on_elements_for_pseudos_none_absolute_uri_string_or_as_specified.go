package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func normalOnElementsForPseudosNoneAbsoluteURIStringOrAsSpecified(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	beforeRef, err := cssxref.Cssxref(env, reg, "::before")
	if err != nil {
		return "", err
	}
	afterRef, err := cssxref.Cssxref(env, reg, "::after")
	if err != nil {
		return "", err
	}
	return "На елементах завжди обчислюється в <code>normal</code>. На " + beforeRef + " та " + afterRef + ", якщо задано <code>none</code>, обчислюється в <code>none</code>. Інакше, для для значень URI, абсолютний URI; для значень <code>attr()</code> – результівний рядок; для інших ключових слів – як задано.", nil
}

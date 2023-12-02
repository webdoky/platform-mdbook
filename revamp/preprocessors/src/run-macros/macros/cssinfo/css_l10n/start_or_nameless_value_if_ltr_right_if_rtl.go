package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func startOrNamelessValueIfLTRRightIfRTL(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	directionRef, err := cssxref.Cssxref(env, reg, "direction")
	if err != nil {
		return "", err
	}
	return "<code>start</code>, або безіменне значення, що діє як <code>left</code>, якщо " + directionRef + " – <code>ltr</code>, чи <code>right</code>, якщо " + directionRef + " – <code>rtl</code>, якщо значення <code>start</code> не підтримується браузером.", nil
}

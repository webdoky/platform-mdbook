package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func lpc(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	return arg1 + ", <a href=\"/uk/docs/Web/CSS/percentage#interpoliatsiia\" title=\"Значення типу даних CSS &gt;percentage&lt; інтерполюються як дійсні числа з рухомою комою.\">відсотки</a> або calc();" + arg2, nil
}

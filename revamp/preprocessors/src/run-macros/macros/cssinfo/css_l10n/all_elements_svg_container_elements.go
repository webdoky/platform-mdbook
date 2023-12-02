package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/svgelement"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func allElementsSVGContainerElements(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	defsRef, err := svgelement.Svgelement(env, reg, "defs")
	if err != nil {
		return "", err
	}
	return "всі елементи; в SVG це застосовується до контейнерів, окрім елемента " + defsRef + "і всіх графічних елементів", nil
}

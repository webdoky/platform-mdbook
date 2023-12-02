package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func allElementsThatCanReferenceImages(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	backroundImageRef, err := cssxref.Cssxref(env, reg, "background-image")
	if err != nil {
		return "", err
	}
	borderImageRef, err := cssxref.Cssxref(env, reg, "border-image")
	if err != nil {
		return "", err
	}
	listStyleImageRef, err := cssxref.Cssxref(env, reg, "list-style-image")
	if err != nil {
		return "", err
	}
	return "Всі елементи, до яких може бути застосовано зображення, наприклад, " + backroundImageRef + ", " + borderImageRef + " або " + listStyleImageRef, nil
}

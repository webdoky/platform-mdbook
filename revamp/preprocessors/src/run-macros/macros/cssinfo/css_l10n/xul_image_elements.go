package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/xulelem"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func xulImageElements(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	xulImageRef, err := xulelem.Xulelem(env, reg, "image")
	if err != nil {
		return "", err
	}
	mozTreeImageRef, err := cssxref.Cssxref(env, reg, ":-moz-tree-image")
	if err != nil {
		return "", err
	}
	mozTreeTwistyRef, err := cssxref.Cssxref(env, reg, ":-moz-tree-twisty")
	if err != nil {
		return "", err
	}
	mozTreeCheckboxRef, err := cssxref.Cssxref(env, reg, ":-moz-tree-checkbox")
	if err != nil {
		return "", err
	}
	listStyleImageRef, err := cssxref.Cssxref(env, reg, "list-style-image")
	if err != nil {
		return "", err
	}
	return "Елементи XUL " + xulImageRef + ", а також псевдоелементи " + mozTreeImageRef + ", " + mozTreeTwistyRef + " і " + mozTreeCheckboxRef + ". <strong>Примітка:</strong> <code>-moz-image-region</code> працює лише на елементах " + xulImageRef + ", чия піктограма задана за допомогою " + listStyleImageRef + ". Це не працюватиме з XUL <code>&lt;image src=\"url\" /&gt;</code>.", nil
}

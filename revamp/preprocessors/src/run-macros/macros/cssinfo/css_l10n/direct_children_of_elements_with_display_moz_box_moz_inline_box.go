package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func directChildrenOfElementsWithDisplayMozBoxMozInlineBox(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	displayRef, err := cssxref.Cssxref(env, reg, "display")
	if err != nil {
		return "", err
	}
	mozBoxRef, err := cssxref.Cssxref(env, reg, "-moz-box")
	if err != nil {
		return "", err
	}
	mozInlineBoxRef, err := cssxref.Cssxref(env, reg, "-moz-inline-box")
	if err != nil {
		return "", err
	}
	webkitBoxRef, err := cssxref.Cssxref(env, reg, "-webkit-box")
	if err != nil {
		return "", err
	}
	webkitInlineBoxRef, err := cssxref.Cssxref(env, reg, "-webkit-inline-box")
	if err != nil {
		return "", err
	}
	return "елементи, що є безпосередніми нащадками елементів, чиє значення " + displayRef + " – " + mozBoxRef + ", " + mozInlineBoxRef + ", " + webkitBoxRef + " або " + webkitInlineBoxRef, nil
}

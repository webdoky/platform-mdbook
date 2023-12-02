package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func sameAsBoxOffsets(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	topRef, err := cssxref.Cssxref(env, reg, "top")
	if err != nil {
		return "", err
	}
	rightRef, err := cssxref.Cssxref(env, reg, "right")
	if err != nil {
		return "", err
	}
	bottomRef, err := cssxref.Cssxref(env, reg, "bottom")
	if err != nil {
		return "", err
	}
	leftRef, err := cssxref.Cssxref(env, reg, "left")
	if err != nil {
		return "", err
	}
	return "дорівнює зміщенням рамки – властивостям " + topRef + ", " + rightRef + ", " + bottomRef + ", " + leftRef + ", причому ці напрямки – логічні", nil
}

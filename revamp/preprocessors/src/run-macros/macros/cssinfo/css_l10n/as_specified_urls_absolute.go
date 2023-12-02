package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func asSpecifiedURLsAbsolute(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	urlRef, err := cssxref.Cssxref(env, reg, "url")
	if err != nil {
		return "", err
	}
	return "як задано, але з " + urlRef + ", зробленими абсолютними", nil
}

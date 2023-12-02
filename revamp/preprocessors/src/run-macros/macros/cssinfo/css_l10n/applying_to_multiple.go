package css_l10n

import (
	"errors"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func applyingToMultiple(_ *environment.Environment, _ *registry.Registry, arg1 string, arg2 string) (string, error) {
	if arg1 == "" {
		return "", errors.New("arg1 is empty")
	}
	if arg2 == "" {
		return "", errors.New("arg2 is empty")
	}
	return arg1 + ". Також застосовується до " + arg2, nil
}

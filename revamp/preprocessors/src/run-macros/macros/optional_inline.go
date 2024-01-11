package macros

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func optional_inline(env *environment.Environment, reg *registry.Registry, _ string) (string, error) {
	return "<span class=\"badge inline optional\">Необов'язкове</span>", nil
}

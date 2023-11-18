package macros

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func blank(env *environment.Environment, reg *registry.Registry, _ string) (string, error) {
	return "", nil
}

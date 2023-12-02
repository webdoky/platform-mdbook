package cssinfo

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func appliesto(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	if data.AppliesTo == "" {
		return "", nil
	}
	return as_single(env, reg, data, data.AppliesTo)
}

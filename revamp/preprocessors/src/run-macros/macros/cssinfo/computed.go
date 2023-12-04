package cssinfo

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func computed(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	if data.Computed == nil || len(data.Computed) == 0 {
		return "", nil
	}
	if len(data.Computed) == 1 {
		return as_single(env, reg, data, data.Computed[0], false)
	}
	return as_longhands(env, reg, data.Computed, computed)
}

package cssinfo

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func percentages(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	if data.Percentages == nil || len(data.Percentages) == 0 {
		return "", nil
	}
	if len(data.Percentages) == 1 {
		return as_single(env, reg, data, data.Percentages[0])
	}
	return as_longhands(env, reg, data.Percentages)
}

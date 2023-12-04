package cssinfo

import (
	"log"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func initial(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	initialValue := data.Initial
	log.Printf("initial: %v", initialValue)
	if len(initialValue) == 1 {
		return as_single(env, reg, data, initialValue[0], false)
	} else {
		return as_longhands(env, reg, initialValue, initial)
	}
}

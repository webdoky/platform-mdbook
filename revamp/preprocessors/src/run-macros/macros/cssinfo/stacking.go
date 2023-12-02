package cssinfo

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssinfo/css_l10n"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func stacking(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	var value string
	if data.Stacking {
		value = "yes"
	} else {
		value = "no"
	}
	localizedValue, err := css_l10n.Localize(env, reg, value, "", "")
	if err != nil {
		return "", err
	}
	return localizedValue, nil
}

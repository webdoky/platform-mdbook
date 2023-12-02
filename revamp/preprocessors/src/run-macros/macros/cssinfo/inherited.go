package cssinfo

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssinfo/css_l10n"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func inherited(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	key := "no"
	if data.Inherited {
		key = "yes"
	}
	localizedString, err := css_l10n.Localize(env, reg, key, "", "")
	if err != nil {
		return "", err
	}
	return localizedString, nil
}

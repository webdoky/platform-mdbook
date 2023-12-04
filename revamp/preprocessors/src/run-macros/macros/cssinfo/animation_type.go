package cssinfo

import (
	"strings"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssinfo/css_l10n"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func animationType(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	if len(data.AnimationType) == 0 {
		return "", nil
	}
	if len(data.AnimationType) == 1 {
		values := strings.Split(data.AnimationType[0], " ")
		parsedValues := []string{}
		for _, value := range values {
			firstParam := ""
			if value == "lpc" {
				var err error
				firstParam, err = css_l10n.Localize(env, reg, "length", "", "")
				if err != nil {
					return "", err
				}
			}
			localizedValue, err := css_l10n.Localize(env, reg, value, firstParam, "")
			if err != nil {
				return "", err
			}
			parsedValues = append(parsedValues, localizedValue)
		}
		listSeparator, err := css_l10n.Localize(env, reg, "listSeparator", "", "")
		if err != nil {
			return "", err
		}
		return strings.Join(parsedValues, listSeparator), nil
	}
	return as_longhands(env, reg, data.AnimationType, animationType)
}

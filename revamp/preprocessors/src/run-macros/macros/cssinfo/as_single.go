package cssinfo

import (
	"strings"
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssinfo/css_l10n"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func as_single(env *environment.Environment, reg *registry.Registry, data *CssData, value string, includeAlso bool) (string, error) {
	keywords := strings.Split(value, ", ")
	localizedKeywords := []string{}
	for _, keyword := range keywords {
		localizedKeyword, err := css_l10n.Localize(env, reg, keyword, "", "")
		if err != nil {
			return "", err
		}
		// if localizedKeyword != keyword {
		// 	log.Printf("Localized keyword %s: %s", keyword, localizedKeyword)
		// }
		localizedKeywords = append(localizedKeywords, localizedKeyword)
	}
	result := strings.Join(localizedKeywords, ", ")
	if !includeAlso {
		return result, nil
	}
	alsoAppliesToOutput, err := also_applies_to(env, reg, data)
	if err != nil {
		return "", err
	}
	return css_l10n.Localize(env, reg, "applyingToMultiple", result, alsoAppliesToOutput)
}

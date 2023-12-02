package cssinfo

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func also_applies_to(env *environment.Environment, reg *registry.Registry, data *CssData) (string, error) {
	if data.AlsoAppliesTo == nil || len(data.AlsoAppliesTo) == 0 {
		return "", nil
	}
	// Remove '::placeholder' from array to avoid displaying it,
	// because it's not standardized yet
	alsoAppliesTo := []string{}
	for _, item := range data.AlsoAppliesTo {
		if item != "::placeholder" {
			alsoAppliesTo = append(alsoAppliesTo, item)
		}
	}
	if len(alsoAppliesTo) == 0 {
		return "", nil
	}
	result := ""
	for index, item := range alsoAppliesTo {
		ref, err := cssxref.Cssxref(env, reg, item)
		if err != nil {
			return "", err
		}
		result += ref
		if index < len(alsoAppliesTo)-2 {
			result += ", "
		} else if index < len(alsoAppliesTo)-1 {
			result += " і "
		}
	}
	return result, nil
}

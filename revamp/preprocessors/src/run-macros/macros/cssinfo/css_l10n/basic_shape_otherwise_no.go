package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func basicShapeOtherwiseNo(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	basicShapeRef, err := cssxref.Cssxref(env, reg, "basic-shape")
	if err != nil {
		return "", err
	}
	return "так, як задано для " + basicShapeRef + ", інакше – ні", nil
}

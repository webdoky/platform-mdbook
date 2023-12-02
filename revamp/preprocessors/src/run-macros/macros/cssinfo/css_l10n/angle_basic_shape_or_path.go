package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func angleBasicShapeOrPath(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	angleRef, err := cssxref.Cssxref(env, reg, "<angle>")
	if err != nil {
		return "", err
	}
	basicShapeRef, err := cssxref.Cssxref(env, reg, "<basic-shape>")
	if err != nil {
		return "", err
	}
	pathRef, err := cssxref.Cssxref(env, reg, "<path()>")
	if err != nil {
		return "", err
	}
	return angleRef + ", " + basicShapeRef + " або " + pathRef, nil
}

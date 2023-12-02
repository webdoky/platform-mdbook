package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func asDefinedForBasicShapeWithAbsoluteURIOtherwiseAsSpecified(env *environment.Environment, reg *registry.Registry, _ string, _ string) (string, error) {
	basicShapeRef, err := cssxref.Cssxref(env, reg, "basic-shape")
	if err != nil {
		return "", err
	}
	shapeBoxRef, err := cssxref.Cssxref(env, reg, "shape-box")
	if err != nil {
		return "", err
	}
	imageRef, err := cssxref.Cssxref(env, reg, "image")
	if err != nil {
		return "", err
	}
	return "як визначено для " + basicShapeRef + " (з наступним " + shapeBoxRef + ", якщо задано), " + imageRef + " з його абсолютним URI, а інакше – як задано.", nil
}

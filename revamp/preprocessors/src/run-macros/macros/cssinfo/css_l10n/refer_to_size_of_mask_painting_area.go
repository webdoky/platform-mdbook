package css_l10n

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

func referToSizeOfMaskPaintingArea(env *environment.Environment, reg *registry.Registry, arg1 string, arg2 string) (string, error) {
	backgroundPositionRef, err := cssxref.Cssxref(env, reg, "background-position")
	if err != nil {
		return "", err
	}
	return "стосується розміру маски області малювання мінус розмір маски шару зображення (дивіться опис " + backgroundPositionRef + ")", nil
}

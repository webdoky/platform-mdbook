package macros

import (
	"webdoky3/preprocessors/src/run-macros/environment"
	"webdoky3/preprocessors/src/run-macros/macros/embedlivesample"
	"webdoky3/preprocessors/src/run-macros/registry"
)

var MacrosIndex = map[string]func(*environment.Environment, *registry.Registry, string) (string, error){
	"domxref":                domxref,
	"embedlivesample":        embedlivesample.Embedlivesample,
	"glossary":               glossary,
	"glossarydisambiguation": glossarydisambiguation,
	"glossarysidebar":        blank,
	"htmlelement":            htmlelement,
	"jsref":                  blank,
	"jsxref":                 jsxref,
}

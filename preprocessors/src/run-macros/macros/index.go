package macros

import (
	"webdoky3/preprocessors/src/run-macros/environment"
	"webdoky3/preprocessors/src/run-macros/macros/embedlivesample"
	"webdoky3/preprocessors/src/run-macros/registry"
)

var MacrosIndex = map[string]func(*environment.Environment, *registry.Registry, string) (string, error){
	"cssref":                 blank,
	"cssxref":                cssxref,
	"deprecated_header":      deprecated_header,
	"domxref":                domxref,
	"embedlivesample":        embedlivesample.Embedlivesample,
	"glossary":               glossary,
	"glossarydisambiguation": glossarydisambiguation,
	"glossarysidebar":        blank,
	"htmlelement":            htmlelement,
	"httpheader":             httpheader,
	"jsref":                  blank,
	"jsxref":                 jsxref,
	"listsubpages":           listsubpages,
	"non-standard_header":    non_standard_header,
	"rfc":                    rfc,
}

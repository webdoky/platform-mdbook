package macros

import (
	"webdoky3/revamp/preprocessors/src/run-macros/environment"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/compat"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssinfo"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/cssxref"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/embedlivesample"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/htmlelement"
	"webdoky3/revamp/preprocessors/src/run-macros/macros/svgelement"
	"webdoky3/revamp/preprocessors/src/run-macros/registry"
)

var MacrosIndex = map[string]func(*environment.Environment, *registry.Registry, string) (string, error){
	"compat":                  compat.Compat,
	"cssinfo":                 cssinfo.Cssinfo,
	"cssref":                  blank,
	"csssyntax":               csssyntax,
	"cssxref":                 cssxref.Cssxref,
	"deprecated_header":       deprecated_header,
	"domxref":                 domxref,
	"embedghlivesample":       embedghlivesample,
	"embedinteractiveexample": embedinteractiveexample,
	"embedlivesample":         embedlivesample.Embedlivesample,
	"experimental_inline":     experimental_inline,
	"experimentalbadge":       experimentalbadge,
	"glossary":                glossary,
	"glossarydisambiguation":  glossarydisambiguation,
	"glossarysidebar":         blank,
	"htmlelement":             htmlelement.Htmlelement,
	"httpheader":              httpheader,
	"jsref":                   blank,
	"jsxref":                  jsxref,
	"listsubpages":            listsubpages,
	"non-standard_header":     non_standard_header,
	"optional_inline":         optional_inline,
	"rfc":                     rfc,
	"seecompattable":          seecompattable,
	"specifications":          specifications,
	"svgelement":              svgelement.Svgelement,
}

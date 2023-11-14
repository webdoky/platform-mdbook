package macros

var MacrosIndex = map[string]func(*Environment, Registry, string) (string, error){
	"domxref":         domxref,
	"glossary":        glossary,
	"glossarysidebar": blank,
	"jsref":           blank,
	"jsxref":          jsxref,
}

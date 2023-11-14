package macros

type Registry interface {
	HasPath(path string) bool
}

type Environment struct {
	Locale string
}

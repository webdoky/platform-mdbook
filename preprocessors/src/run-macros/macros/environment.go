package macros

import "webdoky3/preprocessors/src/preprocessor"

type Registry interface {
	GetSubItems(path string) []preprocessor.Section
	HasPath(path string) bool
}

type Environment struct {
	Locale  string
	Path    string
	Summary string
}

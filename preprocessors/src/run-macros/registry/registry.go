package registry

import (
	"webdoky3/preprocessors/src/preprocessor"
	"webdoky3/preprocessors/src/run-macros/macros"
)

type Registry struct {
	book *preprocessor.Book
}

func NewRegistry(book *preprocessor.Book) macros.Registry {
	return &Registry{
		book: book,
	}
}

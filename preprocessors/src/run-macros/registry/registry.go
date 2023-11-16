package registry

import (
	"webdoky3/preprocessors/src/preprocessor"
)

type Registry struct {
	book *preprocessor.Book
}

func NewRegistry(book *preprocessor.Book) *Registry {
	return &Registry{
		book: book,
	}
}

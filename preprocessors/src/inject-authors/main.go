package main

import (
	"log"
	"webdoky3/preprocessors/src/preprocessor"
)

func main() {
	p := preprocessor.NewPreprocessor(injectAuthors)
	err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

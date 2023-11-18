package main

import (
	"encoding/json"
	"log"
	"os"
	"webdoky3/revamp/preprocessors/src/preprocessor"
)

func saveJson(book *preprocessor.Book, context *preprocessor.Context) (*preprocessor.Book, error) {
	log.Println("Saving JSON")
	bookData, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}

	contextData, err := json.Marshal(context)
	if err != nil {
		return nil, err
	}

	if os.Getenv("PRODUCTION") != "true" {
		err = os.WriteFile("./contextData.json", contextData, 0644)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile("./bookData.json", bookData, 0644)
		if err != nil {
			return nil, err
		}
	}

	return book, nil
}

func main() {
	log.Println("Starting writer")
	p := preprocessor.NewPreprocessor(saveJson)
	err := p.Run()
	if err != nil {

		log.Fatal(err)
	}
}

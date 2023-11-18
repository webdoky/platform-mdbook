package preprocessor

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

// func (s *Section) UnmarshalJSON(data []byte) error {
// 	helpers.WriteToStderr("Unmarshal section\n")
// 	var raw interface{}
// 	err := json.Unmarshal(data, &raw)
// 	if err != nil {
// 		return err
// 	}
// 	switch raw.(type) {
// 	case string:
// 		s.Chapter = nil
// 		s.IsSeparator = true
// 		return nil
// 	case map[string]interface{}:
// 		s.IsSeparator = false
// 		err := json.Unmarshal(data, &s)
// 		helpers.WriteToStderr((s.Chapter.Name))
// 		// s.Chapter, err = raw.(*map[string]interface{})["chapter"].(*Chapter)
// 		return err
// 	}
// 	return errors.New("invalid section")
// }

type Preprocessor struct {
	process func(book *Book, context *Context) (*Book, error)
}

func NewPreprocessor(process func(book *Book, context *Context) (*Book, error)) *Preprocessor {
	return &Preprocessor{process: process}
}

func (p *Preprocessor) Run() error {
	if len(os.Args) > 2 && os.Args[1] == "supports" && os.Args[2] == "html" {
		return nil
	}
	logFile := LogToFile()
	defer logFile.Close()

	log.Println("ReadAll from Stdin starts")
	jsonInput, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	// helpers.WriteToStderr(string(jsonInput))
	// jsonInput, err := ReadStdinToBuffer()
	if err != nil {
		return err
	}
	// helpers.WriteToStderr(fmt.Sprintf("%+v\n", jsonInput.String()))
	log.Println("ReadAll from Stdin ended")

	var input []interface{}
	if err := json.Unmarshal(jsonInput, &input); err != nil {
		return err
	}
	// if err := decoder.Decode(&input); err != nil {
	// 	return err
	// }
	// helpers.WriteToStderr(fmt.Sprintf("%+v\n", input))

	if len(input) != 2 {
		return errors.New("invalid input")
	}

	contextMap, ok := input[0].(map[string]interface{})
	if !ok {
		return errors.New("invalid context")
	}
	context := ConvertMapToContext(contextMap)

	bookMap, ok := input[1].(map[string]interface{})
	if !ok {
		return errors.New("invalid book")
	}
	// if bookMap["sections"] == nil || len(bookMap["sections"].([]interface{})) == 0 {
	// 	return errors.New("empty book")
	// }
	// contextBytes, err := json.Marshal(contextMap)
	// if err != nil {
	// 	return err
	// }
	// var context Context
	// if err := json.Unmarshal(contextBytes, &context); err != nil {
	// 	return err
	// }
	// coerce input as a tuple of (context, book)
	//
	book := ConvertMapToBook(bookMap)
	// context := input[1].(*Context)

	// bookBytes, err := json.Marshal(bookMap)
	// if err != nil {
	// 	return err
	// }
	// var book Book
	// if err := json.Unmarshal(bookBytes, &book); err != nil {
	// 	return err
	// }
	// helpers.WriteToStderr((book.Sections[0].Chapter.Name))

	processedBook, err := p.process(book, context)
	if err != nil {
		return err
	}

	processedBookBytes, err := json.Marshal(processedBook)
	if err != nil {
		return err
	}

	// Write processedBookBytes to a JSON file
	// called as the process called
	// (e.g., "run-macros.json")
	if os.Getenv(("PRODUCTION")) != "true" {
		if err := os.WriteFile("./"+os.Args[0]+".json", processedBookBytes, 0644); err != nil {
			return err
		}
	}

	if _, err := os.Stdout.Write(processedBookBytes); err != nil {
		return err
	}

	return nil
}

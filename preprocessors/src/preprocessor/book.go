package preprocessor

import (
	"encoding/json"
	"log"
)

type Chapter struct {
	Content     string    `json:"content"`
	Name        string    `json:"name"`
	Number      []int     `json:"number"`
	ParentNames []string  `json:"parent_names"`
	Path        string    `json:"path"`
	SourcePath  string    `json:"source_path"`
	SubItems    []Section `json:"sub_items"`
}

type Section struct {
	Chapter     *Chapter `json:"Chapter"`
	IsSeparator bool
}

func (s *Section) MarshalJSON() ([]byte, error) {
	if s.IsSeparator {
		return json.Marshal("Separator")
	} else {
		return json.Marshal(&struct {
			Chapter Chapter `json:"Chapter"`
		}{
			Chapter: *s.Chapter,
		})
	}
}

type Book struct {
	Sections      []Section    `json:"sections"`
	NonExhaustive *interface{} `json:"__non_exhaustive"`
}

func ConvertInterfaceToIntArray(data interface{}) []int {
	var intArray []int
	for _, number := range data.([]interface{}) {
		intArray = append(intArray, int(number.(float64)))
	}
	//return empty array instead of nil
	if intArray == nil {
		return []int{}
	}
	return intArray
}

func ConvertInterfaceToStringArray(data interface{}) []string {
	var stringArray []string
	for _, name := range data.([]interface{}) {
		stringArray = append(stringArray, name.(string))
	}
	//return empty array instead of nil
	if stringArray == nil {
		return []string{}
	}
	return stringArray
}

func ConvertMapToChapter(data map[string]interface{}) Chapter {
	var chapter Chapter
	chapter.Content = data["content"].(string)
	chapter.Name = data["name"].(string)
	chapter.Number = ConvertInterfaceToIntArray(data["number"])
	chapter.ParentNames = ConvertInterfaceToStringArray(data["parent_names"])
	chapter.Path = data["path"].(string)
	chapter.SourcePath = data["source_path"].(string)
	chapter.SubItems = ConvertInterfaceToSections(data["sub_items"])
	return chapter
}

func ConvertInterfaceToSection(data interface{}) Section {
	var section Section
	switch data.(type) {
	case string:
		section.IsSeparator = true
	case map[string]interface{}:
		section.IsSeparator = false
		chapter := ConvertMapToChapter(data.(map[string]interface{})["Chapter"].(map[string]interface{}))
		section.Chapter = &chapter
	}
	return section
}

func ConvertInterfaceToSections(data interface{}) []Section {
	var sections []Section
	for _, section := range data.([]interface{}) {
		sections = append(sections, ConvertInterfaceToSection(section))
	}
	if sections == nil {
		return []Section{}
	}
	return sections
}

func ConvertMapToBook(data map[string]interface{}) *Book {
	var book Book
	book.Sections = ConvertInterfaceToSections(data["sections"].([]interface{}))
	log.Printf("Book.Sections: %d\n", len(book.Sections))
	return &book
}

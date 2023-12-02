package l10n

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const cssJsonFilePath = "./content/files/jsondata/L10n-CSS.json"

var cssTranslationData map[string]map[string]string

func L10nCss(locale string, key string) string {
	keyData, ok := cssTranslationData[key]
	if !ok {
		return key
	}
	translation, ok := keyData[locale]
	if !ok || translation == "" {
		return key
	}
	return translation
}

func init() {
	cssTranslationData = make(map[string]map[string]string)
	jsonFile, err := os.Open(cssJsonFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &cssTranslationData)
	if err != nil {
		log.Fatal(err)
	}
}

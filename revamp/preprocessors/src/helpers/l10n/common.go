package l10n

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const jsonFilePath = "./content/files/jsondata/L10n-Common.json"

var translationData map[string]map[string]string

func L10nCommon(locale string, key string) string {
	keyData, ok := translationData[key]
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
	translationData = make(map[string]map[string]string)
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &translationData)
	if err != nil {
		log.Fatal(err)
	}
}

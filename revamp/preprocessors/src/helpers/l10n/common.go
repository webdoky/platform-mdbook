package l10n

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const commonJsonFilePath = "./content/files/jsondata/L10n-Common.json"

var commonL10nData map[string]map[string]string

func L10nCommon(locale string, key string) string {
	keyData, ok := commonL10nData[key]
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
	commonL10nData = make(map[string]map[string]string)
	jsonFile, err := os.Open(commonJsonFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &commonL10nData)
	if err != nil {
		log.Fatal(err)
	}
}

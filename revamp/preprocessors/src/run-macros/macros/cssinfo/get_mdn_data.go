package cssinfo

import (
	"encoding/json"
	"log"
	"os"
)

var mdnData = make(map[string]map[string]CssData)

func get_mdn_data(blockName string, key string) (*CssData, error) {
	var block, ok = mdnData[blockName]
	if !ok {
		block = make(map[string]CssData)
		rawJson, err := os.ReadFile("./data/css/" + blockName + ".json")
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(rawJson, &block)
		if err != nil {
			return nil, err
		}
		mdnData[blockName] = block
	}
	data, ok := block[key]
	if !ok {
		log.Printf("Key %s not found in block %s", key, blockName)
		return nil, nil
	}
	return &data, nil
}

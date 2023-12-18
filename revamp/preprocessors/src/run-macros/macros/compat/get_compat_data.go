package compat

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

// var compatDataMap map[string]*CompatData

type CompatDataBlock map[string]map[string]map[string]map[string]CompatData

func get_compat_data(key string) (*CompatData, error) {
	// compatData, ok := compatDataMap[key]
	// if !ok {
	key = strings.ReplaceAll(key, ".", "/")
	rawJson, err := os.ReadFile("./browser-compat-data/" + key + ".json")
	if err != nil {
		return nil, err
	}
	var data CompatDataBlock
	err = json.Unmarshal(rawJson, &data)
	if err != nil {
		return nil, err
	}
	keyParts := strings.Split(key, "/")
	fourthKey := "__compat"
	if len(keyParts) == 4 {
		fourthKey = keyParts[3]
	}
	compatDataRecord, ok := data[keyParts[0]][keyParts[1]][keyParts[2]][fourthKey]
	if !ok {
		return nil, errors.New("compat data not found")
	}
	compatData := &compatDataRecord
	// compatDataMap[key] = compatData
	// }
	return compatData, nil
}

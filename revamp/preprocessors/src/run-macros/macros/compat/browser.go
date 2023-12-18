package compat

import (
	"encoding/json"
	"errors"
	"os"
)

type BrowserData struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Upstream string `json:"upstream"`
}

var browsersData map[string]*BrowserData

func get_browser_data(browserId string) (*BrowserData, error) {
	browserData, ok := browsersData[browserId]
	if !ok {
		rawJson, err := os.ReadFile("./browser-compat-data/browsers/" + browserId + ".json")
		if err != nil {
			return nil, err
		}
		var data map[string]map[string]BrowserData
		err = json.Unmarshal(rawJson, &data)
		if err != nil {
			return nil, err
		}
		browserDataRecord, ok := data["browsers"][browserId]
		if !ok {
			return nil, errors.New("browser data not found")
		}
		browserData = &browserDataRecord
		browsersData[browserId] = browserData
	}
	return browserData, nil
}

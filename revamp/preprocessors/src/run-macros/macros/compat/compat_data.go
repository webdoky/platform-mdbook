package compat

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

type BrowserSupportData struct {
	Notes          []string `json:"notes"`
	VersionAdded   string   `json:"version_added"`
	VersionRemoved string   `json:"version_removed"`
}

func (bsd *BrowserSupportData) UnmarshalJSON(data []byte) error {
	//	bsd = &BrowserSupportData{}
	var rawFields map[string]*json.RawMessage
	log.Println("Trying to unmarshal bsd")
	err := json.Unmarshal(data, &rawFields)
	if err != nil {
		return err
	}
	if data, ok := rawFields["version_added"]; ok {
		stringData := ""
		err := json.Unmarshal(*data, &stringData)
		if err != nil {
			if strings.Contains(err.Error(), "cannot unmarshal") {
				var value bool
				err := json.Unmarshal(*data, &value)
				if err != nil {
					return err
				}
				if value {
					stringData = "true"
				} else {
					stringData = "false"
				}
			} else {
				return err
			}
		}
		bsd.VersionAdded = stringData
	}
	if data, ok := rawFields["version_removed"]; ok {
		stringData := ""
		err := json.Unmarshal(*data, &stringData)
		if err != nil {
			if strings.Contains(err.Error(), "cannot unmarshal") {
				var value bool
				err := json.Unmarshal(*data, &value)
				if err != nil {
					return err
				}
				if value {
					stringData = "true"
				} else {
					stringData = "false"
				}
			} else {
				return err
			}
		}
		bsd.VersionRemoved = stringData
	}
	if data, ok := rawFields["notes"]; ok {
		stringData := ""
		err := json.Unmarshal(*data, &stringData)
		if err != nil {
			if strings.Contains(err.Error(), "cannot unmarshal") {
				var stringsValue []string
				err := json.Unmarshal(*data, &stringsValue)
				if err != nil {
					return err
				}
				bsd.Notes = stringsValue
			} else {
				return err
			}
		} else {
			bsd.Notes = []string{stringData}
		}
	}
	log.Printf("bsd: %v\n", *bsd)
	return nil
}

type StatusData struct {
	IsDeprecated    bool `json:"deprecated"`
	IsExperimental  bool `json:"experimental"`
	IsStandardTrack bool `json:"standard_track"`
}

type SupportMap map[string]*BrowserSupportData

func resolve_mirror(browserId string, rawSupportData *map[string]*json.RawMessage) (*BrowserSupportData, error) {
	browserData, err := get_browser_data(browserId)
	if err != nil {
		return nil, err
	}
	supportData := &BrowserSupportData{}
	err = json.Unmarshal(*(*rawSupportData)[browserData.Upstream], supportData)
	if err != nil {
		if strings.Contains(err.Error(), "cannot unmarshal") {
			supportData, err = resolve_mirror(browserData.Upstream, rawSupportData)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	log.Printf("Resolved mirror for %s: %v\n", browserId, supportData)
	return supportData, nil
}

func (s *SupportMap) UnmarshalJSON(data []byte) error {
	s = &SupportMap{}
	var rawFields map[string]*json.RawMessage
	log.Println("Trying to unmarshal")
	err := json.Unmarshal(data, &rawFields)
	if err != nil {
		return err
	}
	fmt.Printf("raw fields: %v\n", rawFields)
	for browserId, rawField := range rawFields {
		stringData := ""
		err := json.Unmarshal(*rawField, &stringData)
		if err != nil {
			if strings.Contains(err.Error(), "cannot unmarshal") {
				var supportData BrowserSupportData
				err := json.Unmarshal(*rawField, &supportData)
				if err != nil {
					return err
				}
				(*s)[browserId] = &supportData
			} else {
				return err
			}
		} else {
			if stringData != "mirror" {
				return errors.New("stringData is not mirror")
			}
			supportData, err := resolve_mirror(browserId, &rawFields)
			if err != nil {
				return err
			}
			(*s)[browserId] = supportData
		}
	}
	return nil
}

type CompatData struct {
	MdnUrl  string      `json:"mdn_url"`
	Status  *StatusData `json:"status"`
	Support SupportMap  `json:"support"`
}

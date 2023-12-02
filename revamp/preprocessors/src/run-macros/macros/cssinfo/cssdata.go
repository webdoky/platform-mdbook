package cssinfo

import (
	"encoding/json"
	"errors"
)

type StringOrArray []string

func (sa *StringOrArray) UnmarshalJSON(data []byte) error {
	var jsonObj interface{}
	err := json.Unmarshal(data, &jsonObj)
	if err != nil {
		return err
	}
	switch obj := jsonObj.(type) {
	case string:
		*sa = StringOrArray([]string{obj})
		return nil
	case []interface{}:
		s := make([]string, 0, len(obj))
		for _, v := range obj {
			value, ok := v.(string)
			if !ok {
				return errors.New("unsupported type")
			}
			s = append(s, value)
		}
		*sa = StringOrArray(s)
		return nil
	}
	return errors.New("unsupported type")
}

type CssData struct {
	AlsoAppliesTo []string            `json:"alsoAppliesTo"`
	AnimationType StringOrArray       `json:"animationType"`
	AppliesTo     string              `json:"appliesto"`
	Computed      StringOrArray       `json:"computed"`
	Descriptors   map[string]*CssData `json:"descriptors"`
	Groups        []string            `json:"groups"`
	Inherited     bool                `json:"inherited"`
	Initial       StringOrArray       `json:"initial"`
	Interfaces    []string            `json:"interfaces"`
	MdnUrl        string              `json:"mdn_url"`
	Media         StringOrArray       `json:"media"`
	Order         string              `json:"order"`
	Percentages   StringOrArray       `json:"percentages"`
	Stacking      bool                `json:"stacking"`
	Status        string              `json:"status"`
	Syntax        string              `json:"syntax"`
}

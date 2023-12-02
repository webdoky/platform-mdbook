package cssinfo

import (
	"encoding/json"
	"fmt"
	"os"
)

type CssDescriptor struct {
	Computed    string
	Initial     string
	Media       []string
	Order       string
	Percentages string
	Status      string
	Syntax      string
}

type AtRule struct {
	Descriptors map[string]CssDescriptor
	Groups      []string `json:"groups"`
	Interfaces  []string `json:"interfaces"`
	MdnUrl      string   `json:"mdn_url"`
	Status      string   `json:"status"`
	Syntax      string   `json:"syntax"`
}

var atRules = make(map[string]AtRule)

func loadAtRules() error {
	data, err := os.ReadFile("./data/css/at-rules.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &atRules); err != nil {
		return err
	}
	return nil
}

func getAtRuleInfo(atRuleName string) (*AtRule, error) {
	if atRules == nil {
		var err error
		err = loadAtRules()
		if err != nil {
			return nil, err
		}
	}
	atRule, ok := atRules[atRuleName]
	if !ok {
		return nil, fmt.Errorf("At-rule %s not found", atRuleName)
	}
	return &atRule, nil
}

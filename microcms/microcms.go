package microcms

import (
	"encoding/json"
	"io/ioutil"
)

type (
	Schema struct {
		APIFields    []APIField    `json:"apiFields"`
		CustomFields []CustomField `json:"customFields"`
	}

	APIField struct {
		IDValue                 string                  `json:"idValue"`
		FieldID                 string                  `json:"fieldId"`
		Name                    string                  `json:"name"`
		Description             string                  `json:"description"`
		Kind                    string                  `json:"kind"`
		IsUnique                bool                    `json:"isUnique"`
		Required                bool                    `json:"required"`
		SelectItems             []SelectItem            `json:"selectItems"`
		MultipleSelect          bool                    `json:"multipleSelect"`
		DateFormat              bool                    `json:"dateFormat"`
		TextSizeLimitValidation TextSizeLimitValidation `json:"textSizeLimitValidation"`
		PatternMatchValidation  PatternMatchValidation  `json:"patternMatchValidation"`
	}

	SelectItem struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	}

	TextSizeLimitValidation struct {
		TextSize TextSize `json:"textSize"`
	}

	TextSize struct {
		Min uint64 `json:"min"`
		Max uint64 `json:"max"`
	}

	PatternMatchValidation struct {
		Regexp Regexp `json:"regexp"`
	}

	Regexp struct {
		Pattern string `json:"pattern"`
	}

	CustomField struct {
		FieldID     string     `json:"fieldId"`
		Name        string     `json:"name"`
		Fields      []APIField `json:"fields"`
		CreatedAt   string     `json:"createdAt"`
		UpdatedAt   string     `json:"updatedAt"`
		ViewerGroup string     `json:"viewerGroup"`
	}
)

func LoadSchemaFromFile(path string) (Schema, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Schema{}, err
	}

	s := Schema{}
	if err := json.Unmarshal(bytes, &s); err != nil {
		return Schema{}, err
	}

	return s, nil
}

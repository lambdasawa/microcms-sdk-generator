package metadata

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type (
	Metadata struct {
		Info *Info  `yaml:"info"`
		APIs []*API `yaml:"apis"`
	}

	Info struct {
		ServiceID string `yaml:"service-id"`
		BasePath  string `yaml:"base-path"`
	}

	API struct {
		Name string `yaml:"name"`

		Path string `yaml:"path"`

		// Relations is map of field name and api name
		Relations map[string]string `yaml:"relations"`

		// Form is `list` or `object`.
		Form string `yaml:"form"`

		// CustomFields is map of field name and custom field name
		CustomFields map[string]string `yaml:"custom-fields"`
	}
)

func LoadFromFile(path string) (Metadata, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Metadata{}, err
	}

	m := Metadata{}
	if err := yaml.Unmarshal(bytes, &m); err != nil {
		return Metadata{}, err
	}

	m.SetDefaultValues()

	return m, nil
}

func (meta *Metadata) SetDefaultValues() {
	meta.Info.SetDefaultValues()
	for i := range meta.APIs {
		meta.APIs[i].SetDefaultValues()
	}
}

func (info *Info) SetDefaultValues() {
	if info.BasePath == "" {
		info.BasePath = "."
	}
}

func (api *API) SetDefaultValues() {
	if api.Path == "" {
		api.Path = fmt.Sprintf("%s.json", api.Name)
	}
}

func (api *API) BuildPath(meta Metadata) string {
	return filepath.Join(meta.Info.BasePath, api.Path)
}

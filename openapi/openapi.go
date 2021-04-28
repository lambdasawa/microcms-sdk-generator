package openapi

import (
	"encoding/json"
	"io/ioutil"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
	"github.com/lambdasawa/microcms-sdk-generator/microcms"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/components"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/info"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/paths"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/servers"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/tags"
	"gopkg.in/yaml.v2"
)

func BuildFrom(meta metadata.Metadata, schemaMap map[string]microcms.Schema) (t openapi3.T, err error) {
	t = openapi3.T{}

	t.OpenAPI = "3.0.0"

	t.Info = info.BuildFrom(meta)

	t.Servers = servers.BuildFrom(meta)

	t.Tags = tags.Build(meta)

	t.Paths, err = paths.BuildFrom(meta)
	if err != nil {
		return openapi3.T{}, err
	}

	t.Components, err = components.BuildFrom(meta, schemaMap)
	if err != nil {
		return openapi3.T{}, err
	}

	return t, nil
}

func SaveAsFile(t openapi3.T, path string) error {
	jsonBytes, err := t.MarshalJSON()
	if err != nil {
		return err
	}

	i := map[string]interface{}{}
	if err := json.Unmarshal(jsonBytes, &i); err != nil {
		return err
	}

	yamlBytes, err := yaml.Marshal(i)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, yamlBytes, 0644)
}

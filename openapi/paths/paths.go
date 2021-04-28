package paths

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/paths/list"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/paths/object"
)

func BuildFrom(meta metadata.Metadata) (openapi3.Paths, error) {
	paths := openapi3.Paths{}

	for _, api := range meta.APIs {
		switch api.Form {
		case "", "list":
			paths[fmt.Sprintf("/%s", api.Name)] = list.BuildListFrom(api.Name)
			paths[fmt.Sprintf("/%s/{content-id}", api.Name)] = list.BuildSingleFrom(api.Name)
		case "object":
			paths[fmt.Sprintf("/%s", api.Name)] = object.BuildFrom(api.Name)
		default:
			return openapi3.Paths{}, fmt.Errorf("unknowo form: %+v", api)
		}
	}

	return paths, nil
}

package schemas

import (
	"fmt"

	"github.com/lambdasawa/microcms-sdk-generator/metadata"
	"github.com/lambdasawa/microcms-sdk-generator/microcms"
	mcms "github.com/lambdasawa/microcms-sdk-generator/openapi/components/schemas/microcms"

	"github.com/getkin/kin-openapi/openapi3"
)

func BuildFrom(meta metadata.Metadata, schemaMap map[string]microcms.Schema) (schemas openapi3.Schemas, err error) {
	schemas = openapi3.Schemas{}

	for _, api := range meta.APIs {
		schema := schemaMap[api.Name]

		schemas[api.Name], err = buildSchema(api, schema)
		if err != nil {
			return openapi3.Schemas{}, err
		}

		schemas[fmt.Sprintf("%s-list", api.Name)] = buildListSchema(api)

		schemas[fmt.Sprintf("%s-create-request", api.Name)], err = buildSchemaForCreateReq(api, schema)
		if err != nil {
			return openapi3.Schemas{}, err
		}

		schemas[fmt.Sprintf("%s-patch-request", api.Name)], err = buildSchemaForPatchReq(api, schema)
		if err != nil {
			return openapi3.Schemas{}, err
		}

		customSchemas, err := mcms.BuildCustomProperties(api, schema, mcms.Option{
			IncludeCustomFieldID: true,
			IncludeMedia:         true,
			IncludeFile:          true,
		})
		if err != nil {
			return openapi3.Schemas{}, err
		}
		for name, ref := range customSchemas {
			schemas[name] = ref
		}
	}

	schemas["common-update-result"] = buildCommonUpdateResult()

	return schemas, nil
}

func buildSchema(api *metadata.API, schema microcms.Schema) (*openapi3.SchemaRef, error) {
	return mcms.BuildFrom(api, schema, mcms.Option{
		IncludeID:         true,
		IncludeTimestamps: true,
		IncludeMedia:      true,
		IncludeFile:       true,
		ForceOptional:     false,
	})
}

func buildListSchema(api *metadata.API) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Required: []string{"totalCount", "offset", "limit", "contents"},
			Type:     "object",
			Properties: openapi3.Schemas{
				"totalCount": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number"}},
				"offset":     &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number"}},
				"limit":      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number"}},
				"contents": &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:  "array",
						Items: &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s", api.Name)},
					},
				},
			},
		},
	}
}

func buildSchemaForCreateReq(api *metadata.API, schema microcms.Schema) (*openapi3.SchemaRef, error) {
	return mcms.BuildFrom(api, schema, mcms.Option{
		IncludeID:            false,
		IncludeTimestamps:    false,
		IncludeCustomFieldID: true,
		ForceOptional:        false,
		UseIDInsteadOfRef:    true,
	})
}

func buildSchemaForPatchReq(api *metadata.API, schema microcms.Schema) (*openapi3.SchemaRef, error) {
	return mcms.BuildFrom(api, schema, mcms.Option{
		IncludeID:            false,
		IncludeTimestamps:    false,
		IncludeCustomFieldID: true,
		ForceOptional:        true,
		UseIDInsteadOfRef:    true,
	})
}

func buildCommonUpdateResult() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Required: []string{"id"},
			Type:     "object",
			Properties: openapi3.Schemas{
				"id": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "string"}},
			},
		},
	}
}

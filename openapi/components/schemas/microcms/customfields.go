package microcms

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
	"github.com/lambdasawa/microcms-sdk-generator/microcms"
)

func BuildCustomProperties(api *metadata.API, m microcms.Schema, option Option) (map[string]*openapi3.SchemaRef, error) {
	schemas := map[string]*openapi3.SchemaRef{}

	for _, customField := range m.CustomFields {
		customFieldSchema := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:       "object",
				Properties: openapi3.Schemas{},
			},
		}

		for _, field := range customField.Fields {
			ref, err := buildProperty(api, field, option)
			if err != nil {
				return nil, err
			}

			if ref == nil {
				continue
			}
			customFieldSchema.Value.Properties[field.FieldID] = ref
		}

		if option.IncludeCustomFieldID {
			customFieldSchema.Value.Properties["fieldId"] = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "string",
					Enum: []interface{}{customField.FieldID},
				},
			}
		}

		schemas[fmt.Sprintf("%s-%s", api.Name, customField.FieldID)] = customFieldSchema
	}

	return schemas, nil
}

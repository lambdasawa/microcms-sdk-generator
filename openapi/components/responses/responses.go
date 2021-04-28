package responses

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
)

func BuildFrom(meta metadata.Metadata) (openapi3.Responses, error) {
	responses := openapi3.Responses{}

	description := ""

	for _, api := range meta.APIs {
		responses[fmt.Sprintf("%s-list", api.Name)] = &openapi3.ResponseRef{
			Value: &openapi3.Response{
				Description: &description,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Required: []string{"totalCount", "offset", "limit", "contents"},
								Type:     "object",
								Properties: openapi3.Schemas{
									"totalCount": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number", Nullable: false}},
									"offset":     &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number", Nullable: false}},
									"limit":      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number", Nullable: false}},
									"contents": &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Type:  "array",
											Items: &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s", api.Name)},
										},
									},
								},
							},
						},
					},
				},
			},
		}
	}

	return responses, nil
}

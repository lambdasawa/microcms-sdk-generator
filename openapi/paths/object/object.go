package object

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func BuildFrom(name string) *openapi3.PathItem {
	description := ""

	return &openapi3.PathItem{
		Get: &openapi3.Operation{
			OperationID: fmt.Sprintf("fetch-%s", name),
			Tags:        []string{name},
			Security:    &openapi3.SecurityRequirements{{"ApiKeyAuth": {}}},
			Responses: openapi3.Responses{
				"200": &openapi3.ResponseRef{
					Value: &openapi3.Response{
						Description: &description,
						Content: openapi3.Content{
							"application/json": {
								Schema: &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s", name)},
							},
						},
					},
				},
			},
		},

		Patch: &openapi3.Operation{
			OperationID: fmt.Sprintf("update-%s", name),
			Tags:        []string{name},
			Security:    &openapi3.SecurityRequirements{{"WriteApiKeyAuth": {}}},
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Description: "",
					Required:    true,
					Content: openapi3.Content{
						"application/json": {
							Schema: &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s-patch-request", name)},
						},
					},
				},
			},
			Responses: openapi3.Responses{
				"201": &openapi3.ResponseRef{
					Value: &openapi3.Response{
						Description: &description,
						Content: openapi3.Content{
							"application/json": {
								Schema: &openapi3.SchemaRef{Ref: "#/components/schemas/common-update-result"},
							},
						},
					},
				},
			},
		},
	}
}

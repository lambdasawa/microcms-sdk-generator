package list

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func BuildListFrom(name string) *openapi3.PathItem {
	description := ""

	return &openapi3.PathItem{
		Post: &openapi3.Operation{
			OperationID: fmt.Sprintf("create-%s", name),
			Tags:        []string{name},
			Security:    &openapi3.SecurityRequirements{{"WriteApiKeyAuth": {}}},
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Description: "",
					Required:    true,
					Content: openapi3.Content{
						"application/json": {
							Schema: &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s-create-request", name)},
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

		Get: &openapi3.Operation{
			OperationID: fmt.Sprintf("search-%s", name),
			Tags:        []string{name},
			Security: &openapi3.SecurityRequirements{
				{"ApiKeyAuth": {}},
				{"GlobalDraftKeyAuth": {}},
			},
			Parameters: []*openapi3.ParameterRef{
				{Ref: "#/components/parameters/draft-key-querystring"},
				{Ref: "#/components/parameters/limit-querystring"},
				{Ref: "#/components/parameters/offset-querystring"},
				{Ref: fmt.Sprintf("#/components/parameters/%s-orders-querystring", name)},
				{Ref: "#/components/parameters/q-querystring"},
				{Ref: "#/components/parameters/ids-querystring"},
				{Ref: fmt.Sprintf("#/components/parameters/%s-fields-querystring", name)},
				{Ref: "#/components/parameters/filters-querystring"},
				{Ref: "#/components/parameters/depth-querystring"},
			},
			Responses: openapi3.Responses{
				"200": &openapi3.ResponseRef{
					Value: &openapi3.Response{
						Description: &description,
						Content: openapi3.Content{
							"application/json": {
								Schema: &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s-list", name)},
							},
						},
					},
				},
			},
		},
	}
}

func BuildSingleFrom(name string) *openapi3.PathItem {
	description := ""

	return &openapi3.PathItem{
		Get: &openapi3.Operation{
			OperationID: fmt.Sprintf("fetch-%s", name),
			Tags:        []string{name},
			Security:    &openapi3.SecurityRequirements{{"ApiKeyAuth": {}}},
			Parameters: []*openapi3.ParameterRef{
				{Ref: "#/components/parameters/content-id-path"},
				{Ref: "#/components/parameters/draft-key-querystring"},
				{Ref: fmt.Sprintf("#/components/parameters/%s-fields-querystring", name)},
				{Ref: "#/components/parameters/depth-querystring"},
			},
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

		Put: &openapi3.Operation{
			OperationID: fmt.Sprintf("put-%s", name),
			Tags:        []string{name},
			Security:    &openapi3.SecurityRequirements{{"WriteApiKeyAuth": {}}},
			Parameters: []*openapi3.ParameterRef{
				{Ref: "#/components/parameters/content-id-path"},
			},
			RequestBody: &openapi3.RequestBodyRef{
				Value: &openapi3.RequestBody{
					Description: "",
					Required:    true,
					Content: openapi3.Content{
						"application/json": {
							Schema: &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s-create-request", name)},
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

		Patch: &openapi3.Operation{
			OperationID: fmt.Sprintf("update-%s", name),
			Tags:        []string{name},
			Security:    &openapi3.SecurityRequirements{{"WriteApiKeyAuth": {}}},
			Parameters: []*openapi3.ParameterRef{
				{Ref: "#/components/parameters/content-id-path"},
			},
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

		Delete: &openapi3.Operation{
			OperationID: fmt.Sprintf("delete-%s", name),
			Tags:        []string{name},
			Security:    &openapi3.SecurityRequirements{{"WriteApiKeyAuth": {}}},
			Parameters: []*openapi3.ParameterRef{
				{Ref: "#/components/parameters/content-id-path"},
			},
			Responses: openapi3.Responses{
				"202": &openapi3.ResponseRef{
					Value: &openapi3.Response{
						Description: &description,
					},
				},
			},
		},
	}
}

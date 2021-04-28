package parameters

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
	"github.com/lambdasawa/microcms-sdk-generator/microcms"
)

func BuildFrom(meta metadata.Metadata, schemaMap map[string]microcms.Schema) (openapi3.ParametersMap, error) {
	parameters := openapi3.ParametersMap{
		"content-id-path":       buildContent(),
		"draft-key-querystring": buildDraftKey(),
		"limit-querystring":     buildLimit(),
		"offset-querystring":    buildOffset(),
		"q-querystring":         buildQ(),
		"ids-querystring":       buildIDs(),
		"filters-querystring":   buildFilter(),
		"depth-querystring":     buildDepth(),
	}

	for name, ref := range buildOrders(meta, schemaMap) {
		parameters[name] = ref
	}

	for name, ref := range buildFields(meta, schemaMap) {
		parameters[name] = ref
	}

	return parameters, nil
}

func buildContent() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:   "path",
			Name: "content-id",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "string",
				},
			},
			Required: true,
		},
	}
}

func buildDraftKey() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:          "query",
			Name:        "draftKey",
			Description: "[doc](https://document.microcms.io/content-api/get-list-contents#hab2c474417)",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "string",
				},
			},
		},
	}
}

func buildLimit() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:          "query",
			Name:        "limit",
			Description: "[doc](https://document.microcms.io/content-api/get-list-contents#h4cd61f9fa1)",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "number",
				},
			},
		},
	}
}

func buildOffset() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:          "query",
			Name:        "offset",
			Description: "[doc](https://document.microcms.io/content-api/get-list-contents#h41838110ca)",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "number",
				},
			},
		},
	}
}

func buildOrders(meta metadata.Metadata, schemaMap map[string]microcms.Schema) openapi3.ParametersMap {
	parametersMap := openapi3.ParametersMap{}
	for _, api := range meta.APIs {
		name := fmt.Sprintf("%s-orders-querystring", api.Name)

		schema := schemaMap[api.Name]
		var enum, xEnumVarNames []interface{}
		for _, field := range schema.APIFields {
			enum = append(enum, field.FieldID, "-"+field.FieldID)
			xEnumVarNames = append(xEnumVarNames, field.FieldID, field.FieldID+"Descending")
		}
		for _, field := range []string{"id", "publishedAt", "createdAt", "updatedAt", "revisedAt"} {
			enum = append(enum, field, "-"+field)
			xEnumVarNames = append(xEnumVarNames, field, field+"Descending")
		}

		parametersMap[name] = &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:          "query",
				Name:        "orders",
				Description: "[doc](https://document.microcms.io/content-api/get-list-contents#hf1af2f632c)",
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: "array",
						Items: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
								Enum: enum,
								ExtensionProps: openapi3.ExtensionProps{
									Extensions: map[string]interface{}{
										"x-enum-varnames": xEnumVarNames,
									},
								},
							},
						},
					},
				},
				Explode: (func(b bool) *bool { return &b })(false),
			},
		}
	}

	return parametersMap
}

func buildQ() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:          "query",
			Name:        "q",
			Description: "[doc](https://document.microcms.io/content-api/get-list-contents#ha8abec0b2f)",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "string",
				},
			},
		},
	}
}

func buildFields(meta metadata.Metadata, schemaMap map[string]microcms.Schema) openapi3.ParametersMap {
	parametersMap := openapi3.ParametersMap{}
	for _, api := range meta.APIs {
		name := fmt.Sprintf("%s-fields-querystring", api.Name)

		var enum []interface{}
		schema := schemaMap[api.Name]
		for _, field := range schema.APIFields {
			enum = append(enum, field.FieldID)
		}

		parametersMap[name] = &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				In:          "query",
				Name:        "fields",
				Description: "[doc](https://document.microcms.io/content-api/get-list-contents#h7462d83de4)",
				Schema: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: "array",
						Items: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
								Enum: enum,
							},
						},
					},
				},
				Explode: (func(b bool) *bool { return &b })(false),
			},
		}
	}

	return parametersMap

}

func buildIDs() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:          "query",
			Name:        "ids",
			Description: "[doc](https://document.microcms.io/content-api/get-list-contents#h6b992e0fe4)",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "array",
					Items: &openapi3.SchemaRef{
						Value: &openapi3.Schema{Type: "string"},
					},
				},
			},
			Explode: (func(b bool) *bool { return &b })(false),
		},
	}
}

func buildFilter() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:          "query",
			Name:        "filters",
			Description: "[doc](https://document.microcms.io/content-api/get-list-contents#hdebbdc8e86)",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "array",
					Items: &openapi3.SchemaRef{
						Value: &openapi3.Schema{Type: "string"},
					},
				},
			},
			Explode: (func(b bool) *bool { return &b })(false),
		},
	}
}

func buildDepth() *openapi3.ParameterRef {
	return &openapi3.ParameterRef{
		Value: &openapi3.Parameter{
			In:          "query",
			Name:        "depth",
			Description: "[doc](https://document.microcms.io/content-api/get-list-contents#h30fce9c966)",
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "number",
				},
			},
		},
	}
}

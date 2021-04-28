package security

import "github.com/getkin/kin-openapi/openapi3"

func Build() (openapi3.SecuritySchemes, error) {
	return openapi3.SecuritySchemes{
		"ApiKeyAuth": &openapi3.SecuritySchemeRef{
			Value: &openapi3.SecurityScheme{
				Type: "apiKey",
				In:   "header",
				Name: "X-API-Key",
			},
		},
		"WriteApiKeyAuth": &openapi3.SecuritySchemeRef{
			Value: &openapi3.SecurityScheme{
				Type: "apiKey",
				In:   "header",
				Name: "X-Write-API-Key",
			},
		},
		"GlobalDraftKeyAuth": &openapi3.SecuritySchemeRef{
			Value: &openapi3.SecurityScheme{
				Type: "apiKey",
				In:   "header",
				Name: "X-Global-Draft-Key",
			},
		},
	}, nil
}

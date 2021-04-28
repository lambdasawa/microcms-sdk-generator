package components

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
	"github.com/lambdasawa/microcms-sdk-generator/microcms"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/components/parameters"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/components/responses"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/components/schemas"
	"github.com/lambdasawa/microcms-sdk-generator/openapi/components/security"
)

func BuildFrom(meta metadata.Metadata, schemaMap map[string]microcms.Schema) (c openapi3.Components, err error) {
	c = openapi3.Components{}

	c.Schemas, err = schemas.BuildFrom(meta, schemaMap)
	if err != nil {
		return openapi3.Components{}, err
	}

	c.SecuritySchemes, err = security.Build()
	if err != nil {
		return openapi3.Components{}, err
	}

	c.Parameters, err = parameters.BuildFrom(meta, schemaMap)
	if err != nil {
		return openapi3.Components{}, err
	}

	c.Responses, err = responses.BuildFrom(meta)
	if err != nil {
		return openapi3.Components{}, err
	}

	return c, nil
}

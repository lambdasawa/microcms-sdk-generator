package info

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
)

func BuildFrom(meta metadata.Metadata) *openapi3.Info {
	return &openapi3.Info{
		Title:   meta.Info.ServiceID,
		Version: "1.0.0",
	}
}

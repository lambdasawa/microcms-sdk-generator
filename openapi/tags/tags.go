package tags

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
)

func Build(meta metadata.Metadata) openapi3.Tags {
	tags := openapi3.Tags{}

	for _, api := range meta.APIs {
		tags = append(tags, &openapi3.Tag{
			Name: api.Name,
			ExternalDocs: &openapi3.ExternalDocs{
				URL: fmt.Sprintf("https://%s.microcms.io/apis/%s/settings", meta.Info.ServiceID, api.Name),
			},
		})
	}

	return tags
}

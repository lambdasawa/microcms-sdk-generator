package servers

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
)

func BuildFrom(meta metadata.Metadata) openapi3.Servers {
	return openapi3.Servers{
		&openapi3.Server{
			URL: fmt.Sprintf("https://%s.microcms.io/api/v1/", meta.Info.ServiceID),
		},
	}
}

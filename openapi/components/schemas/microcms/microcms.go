package microcms

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/lambdasawa/microcms-sdk-generator/metadata"
	"github.com/lambdasawa/microcms-sdk-generator/microcms"
)

type (
	Option struct {
		IncludeID            bool
		IncludeTimestamps    bool
		IncludeMedia         bool
		IncludeCustomFieldID bool
		ForceOptional        bool
		UseIDInsteadOfRef    bool
	}
)

func BuildFrom(api *metadata.API, m microcms.Schema, option Option) (*openapi3.SchemaRef, error) {
	properties, err := buildProperties(api, m, option)
	if err != nil {
		return nil, err
	}

	requiredField := buildRequiredField(m, option)

	schemaRef := &openapi3.SchemaRef{
		Value: &openapi3.Schema{
			Type:       "object",
			Properties: properties,
			Required:   requiredField,
		},
	}

	return schemaRef, nil
}

func buildProperties(api *metadata.API, m microcms.Schema, option Option) (openapi3.Schemas, error) {
	properties := openapi3.Schemas{}

	apiProperties, err := buildAPIProperties(api, m, option)
	if err != nil {
		return nil, err
	}

	for name, prop := range apiProperties {
		properties[name] = prop
	}

	return properties, nil
}

func buildAPIProperties(api *metadata.API, m microcms.Schema, option Option) (openapi3.Schemas, error) {
	properties := openapi3.Schemas{}

	for name, ref := range buildCommonProperties(option) {
		properties[name] = ref
	}

	for _, field := range m.APIFields {
		ref, err := buildProperty(api, field, option)
		if err != nil {
			return nil, err
		}

		if ref == nil {
			continue
		}
		properties[field.FieldID] = ref
	}

	return properties, nil
}

func buildCommonProperties(option Option) map[string]*openapi3.SchemaRef {
	m := map[string]*openapi3.SchemaRef{}

	if option.IncludeID {
		m["id"] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "string"}}
	}
	if option.IncludeTimestamps {
		dateTimeRef := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:   "string",
				Format: "date-time",
			},
		}
		m["createdAt"] = dateTimeRef
		m["updatedAt"] = dateTimeRef
		m["publishedAt"] = dateTimeRef
		m["revisedAt"] = dateTimeRef
	}

	return m
}

func buildProperty(api *metadata.API, field microcms.APIField, option Option) (*openapi3.SchemaRef, error) {
	var ref *openapi3.SchemaRef

	switch field.Kind {

	case "text", "textArea", "richEditor":
		ref = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "string"}}
		if field.TextSizeLimitValidation.TextSize.Min != 0 {
			ref.Value.MinLength = field.TextSizeLimitValidation.TextSize.Min
		}
		if field.TextSizeLimitValidation.TextSize.Max != 0 {
			ref.Value.MaxLength = &field.TextSizeLimitValidation.TextSize.Max
		}
		if field.PatternMatchValidation.Regexp.Pattern != "" {
			ref.Value.Pattern = field.PatternMatchValidation.Regexp.Pattern
		}

	case "date":
		format := "date-time"
		if field.DateFormat {
			format = "date"
		}

		ref = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:   "string",
				Format: format,
			},
		}

	case "number":
		ref = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number"}}

	case "boolean":
		ref = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "boolean"}}

	case "select":
		var enum []interface{}
		for _, item := range field.SelectItems {
			enum = append(enum, item.Value)
		}

		var maxItems *uint64
		if field.MultipleSelect {
			maxItems = func(n uint64) *uint64 { return &n }(1)
		}

		ref = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:     "array",
				MaxItems: maxItems,
				Items: &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type: "string",
						Enum: enum,
					},
				},
			},
		}

	case "media":
		if !option.IncludeMedia {
			return nil, nil
		}

		ref = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type: "object",
				Properties: openapi3.Schemas{
					"url":    &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "string"}},
					"height": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number"}},
					"width":  &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "number"}},
				},
			},
		}

	case "custom":
		ref = &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s-%s", api.Name, api.CustomFields[field.FieldID])}

	case "relation":
		if option.UseIDInsteadOfRef {
			ref = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "string"}}
		} else {
			ref = &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s", api.Relations[field.FieldID])}
		}

	case "relationList":
		var items *openapi3.SchemaRef
		if option.UseIDInsteadOfRef {
			items = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: "string"}}
		} else {
			items = &openapi3.SchemaRef{Ref: fmt.Sprintf("#/components/schemas/%s", api.Relations[field.FieldID])}
		}
		ref = &openapi3.SchemaRef{
			Value: &openapi3.Schema{Type: "array", Items: items},
		}

	default:
		return nil, fmt.Errorf("unknown kind: %+v", field)
	}

	if ref.Value != nil {
		switch {
		case len(field.Name) != 0 && len(field.Description) != 0:
			ref.Value.Description = fmt.Sprintf("%s: %s", field.Name, field.Description)
		case len(field.Description) != 0:
			ref.Value.Description = field.Description
		case len(field.Name) != 0:
			ref.Value.Description = field.Name
		default:
			ref.Value.Description = ""
		}
	}

	return ref, nil
}

func buildRequiredField(m microcms.Schema, option Option) (required []string) {
	if option.ForceOptional {
		return nil
	}

	required = append(required, []string{"id", "createdAt", "updatedAt", "publishedAt", "revisedAt"}...)
	for _, field := range m.APIFields {
		if field.Required {
			required = append(required, field.FieldID)
		}
	}
	return required
}

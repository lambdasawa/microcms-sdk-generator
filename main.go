package main

import (
	"flag"

	"github.com/lambdasawa/microcms-sdk-generator/metadata"
	"github.com/lambdasawa/microcms-sdk-generator/microcms"
	"github.com/lambdasawa/microcms-sdk-generator/openapi"
)

func main() {
	var inputFilePath, outputFilePath string
	flag.StringVar(&inputFilePath, "input", "metadata.yml", "input file path")
	flag.StringVar(&outputFilePath, "output", "openapi.yml", "output file path")
	flag.Parse()

	meta, err := metadata.LoadFromFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	microCMSSchemaMap := map[string]microcms.Schema{}
	for _, api := range meta.APIs {
		schema, err := microcms.LoadSchemaFromFile(api.BuildPath(meta))
		if err != nil {
			panic(err)
		}
		microCMSSchemaMap[api.Name] = schema
	}

	openapiSchema, err := openapi.BuildFrom(meta, microCMSSchemaMap)
	if err != nil {
		panic(err)
	}

	if err := openapi.SaveAsFile(openapiSchema, outputFilePath); err != nil {
		panic(err)
	}
}

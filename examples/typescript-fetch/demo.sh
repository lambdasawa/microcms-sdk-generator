#!/bin/bash

set -xeu

docker run \
  --rm \
  --pull always \
  --volume ${PWD}/microcms/:/app/microcms/ \
  --env METADATA_PATH=/app/microcms/metadata.yml \
  --env OPENAPI_PATH=/app/microcms/openapi.yml \
  --env OUTPUT_PATH=/app/microcms/ \
  lambdasawa/microcms-sdk-generator:latest \
  --generator-name typescript-fetch \
  --additional-properties=typescriptThreePlus=true,allowUnicodeIdentifiers=true

npm install

npm run demo

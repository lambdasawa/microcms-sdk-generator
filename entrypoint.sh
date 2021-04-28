#!/bin/bash

set -xeu

/app/microcms-sdk-generator-cli \
  -input $METADATA_PATH \
  -output $OPENAPI_PATH

java -jar openapi-generator-cli.jar generate \
  -i $OPENAPI_PATH \
  -o $OUTPUT_PATH \
  $@

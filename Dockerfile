FROM golang:1.16 AS builder

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY metadata metadata
COPY microcms microcms
COPY openapi openapi

RUN GOOS=linux go build -o microcms-sdk-generator-cli

FROM openjdk:17

ENV METADATA_PATH /app/metadata.yml
ENV OPENAPI_PATH /app/openapi.yml
ENV OUTPUT_PATH /app/sdk
ENV OPENAPI_GENERATOR_CLI_VERSION 5.1.0

WORKDIR /app/

RUN curl -sSL -o openapi-generator-cli.jar https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$OPENAPI_GENERATOR_CLI_VERSION/openapi-generator-cli-$OPENAPI_GENERATOR_CLI_VERSION.jar

COPY --from=builder /app/microcms-sdk-generator-cli /app/microcms-sdk-generator-cli

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

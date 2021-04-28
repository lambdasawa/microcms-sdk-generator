.PHONY: \
	run-demo-all \
	run-dmeo \
	generate-example-all \
	generate-example \
	test

run-demo-all:
	make run-demo EXAMPLE=typescript-fetch

generate-example-all:
	make generate-example EXAMPLE=typescript-fetch PROPS=typescriptThreePlus=true,allowUnicodeIdentifiers=true

run-demo:
	cd examples/${EXAMPLE}/ && ./demo.sh

generate-example:
	docker rmi lambdasawa/microcms-sdk-generator:local
	docker build -t lambdasawa/microcms-sdk-generator:local .
	cd examples/${EXAMPLE}/ &&\
 		docker run \
        	--rm \
        	--volume ${PWD}/examples/${EXAMPLE}/microcms/:/app/microcms/ \
        	--env METADATA_PATH=/app/microcms/metadata.yml \
        	--env OUTPUT_PATH=/app/microcms/ \
        	lambdasawa/microcms-sdk-generator:local \
        	--generator-name ${EXAMPLE} \
        	--additional-properties=${PROPS}

test:
	go fmt ./...
	go build -v ./...
	golangci-lint run ./...
	go test -v ./...

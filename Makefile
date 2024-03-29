.PHONY: generate-client
generate-client:
	wget http://cbs:8080/api-docs -O openapi.json
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
	oapi-codegen -package cbs openapi.json > cbs/cbsapi.gen.go

build: generate-client
	go build ./... -o cbs

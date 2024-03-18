.PHONY: generate-client
generate-client:
	wget http://cbs:8080/api-docs -O openapi.json
	openapi-generator-cli generate \
		--skip-validate-spec \
		-i openapi.json \
		-g go \
		-o cbsapi \
		--additional-properties=packageName=cbsapi \
		--git-repo-id build-environment-builder/cbsapi --git-user-id eed-web-application

build: generate-client
	go build ./... -o cbs

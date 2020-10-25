-include .common.mk

export

## Build the application
build: .DEFAULT

codegen: deps generate

deps:
	go mod vendor

generate:
	go generate ./...

## Test runs all project unit tests.
test:
	go test -coverprofile=cover.out $(TEST_OPTS) $(TESTS)

tool-chain:
	go get -u \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/envoyproxy/protoc-gen-validate \
		github.com/awalterschulze/goderive \
		github.com/rleszilm/grpc-graphql-gateway/protoc-gen-graphql

.DEFAULT: codegen test

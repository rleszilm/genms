-include .common.mk

export

## Build the application
build: .DEFAULT

codegen: deps generate

deps:
	go mod vendor

generate:
	go generate ./...

proto: proto-annotations proto-example

proto-annotations:
	protoc \
		-I . \
		--go_out=. \
		--go_opt=paths=source_relative \
		`ls protoc-gen-microservice/annotations/*.proto`

proto-example:
	protoc \
		-I . \
		-I ~/git/rleszilm/grpc-graphql-gateway/include \
		-I ~/git/itsgameday/gameday/.proto/github.com/googleapis/googleapis \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. \
		--grpc-gateway_opt=logtostderr=true,paths=source_relative \
		--graphql_out=. \
		--graphql_opt=paths=source_relative \
		--microservice_out=. \
		--microservice_opt=paths=source_relative \
		`ls protoc-gen-microservice/example/*.proto`

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

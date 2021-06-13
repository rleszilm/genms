-include .common.mk

export

## Build the application
build: .DEFAULT

codegen: deps generate

deps:
	go mod vendor

generate:
	go generate ./...

proto: proto-annotations

proto-annotations:
	protoc \
		-I . \
		--go_out=. \
		--go_opt=paths=source_relative \
		`ls cmd/protoc-gen-go-genms/annotations/*.proto`
	protoc \
		-I . \
		--go_out=. \
		--go_opt=paths=source_relative \
		`ls cmd/protoc-gen-go-genms-dal/annotations/*.proto`
	protoc \
		-I . \
		--go_out=. \
		--go_opt=paths=source_relative \
		`ls cmd/protoc-gen-go-genms-dal/annotations/types/*.proto`

## Test runs all project unit tests.
test:
	go test $(TESTS_OPTS) $(TESTS)

test-all: test test-integration

test-integration: test-env
	docker compose up make

test-env:
	docker compose up -d make

test-clean:
	docker compose down --remove-orphans

tool-chain:
	go get \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/envoyproxy/protoc-gen-validate \
		github.com/awalterschulze/goderive \
		github.com/rleszilm/grpc-graphql-gateway/protoc-gen-graphql \
		github.com/rleszilm/genms/cmd/protoc-gen-go-genms \
		github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal

.DEFAULT: codegen test

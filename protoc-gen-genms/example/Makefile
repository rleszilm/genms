## Build the application
build: .DEFAULT

proto:
	protoc \
		--experimental_allow_proto3_optional \
		-I ../.. \
		-I ../../../.proto/googleapis \
		-I ../../../.proto/grpc-graphql-gateway/include \
		-I greeter \
		--go_out=greeter \
		--go_opt=paths=source_relative \
		--go-grpc_out=greeter \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=greeter \
		--grpc-gateway_opt=logtostderr=true,paths=source_relative \
		--go-genms_out=greeter \
		--go-genms_opt=paths=source_relative \
		`ls greeter/*.proto`

.DEFAULT: proto
	go build main.go

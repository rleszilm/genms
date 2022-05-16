## Build the application
build: .DEFAULT

proto:
	protoc \
		--experimental_allow_proto3_optional \
		-I ../.. \
		-I ../../.proto/github.com/googleapis/googleapis \
		-I ../../.proto/github.com/rleszilm \
		-I greeter \
		--go_out=greeter \
		--go_opt=paths=source_relative \
		--go-grpc_out=greeter \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=greeter \
		--grpc-gateway_opt=logtostderr=true,paths=source_relative \
		--genms_out=greeter \
		--genms_opt=paths=source_relative \
		`ls greeter/*.proto`

.DEFAULT: proto
	go build main.go

## Build the application
build: .DEFAULT

proto:
	protoc \
		-I ../../../.proto/grpc-graphql-gateway/include \
		-I ../../../.proto/googleapis \
		-I ../.. \
		-I users \
		--go_out=users \
		--go_opt=paths=source_relative \
		--go-genms-dal_out=users \
		--go-genms-dal_opt=paths=source_relative \
		`ls users/*.proto`

.DEFAULT: proto

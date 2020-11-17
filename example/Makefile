## Build the application
build: .DEFAULT

proto:
	protoc \
		-I ../.proto/grpc-graphql-gateway/include \
		-I ../.proto/googleapis \
		-I .. \
		-I greeter \
		--go_out=greeter \
		--go_opt=paths=source_relative \
		--go-grpc_out=greeter \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=greeter \
		--grpc-gateway_opt=logtostderr=true,paths=source_relative \
		--graphql_out=greeter \
		--graphql_opt=paths=source_relative \
		--microservice_out=greeter \
		--microservice_opt=paths=source_relative \
		`ls greeter/*.proto`

.DEFAULT:
	go build main.go

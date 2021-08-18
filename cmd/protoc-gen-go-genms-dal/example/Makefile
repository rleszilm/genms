## Build the application
build: .DEFAULT

proto: proto-single proto-multi

proto-single:
	protoc \
		--experimental_allow_proto3_optional \
		-I ../../../.proto/grpc-graphql-gateway/include \
		-I ../../../.proto/googleapis \
		-I ../.. \
		-I single \
		--go_out=single \
		--go_opt=paths=source_relative \
		--go-genms-dal_out=single \
		--go-genms-dal_opt=paths=source_relative \
		`ls single/*.proto`

proto-multi:
	protoc \
		--experimental_allow_proto3_optional \
		-I ../../../.proto/grpc-graphql-gateway/include \
		-I ../../../.proto/googleapis \
		-I ../.. \
		-I multi \
		--go_out=multi \
		--go_opt=paths=source_relative \
		--go-genms-dal_out=multi \
		--go-genms-dal_opt=paths=source_relative \
		`ls multi/*.proto`

.DEFAULT: proto

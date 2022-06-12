package tools

import (
	_ "github.com/awalterschulze/goderive"
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
	_ "github.com/rleszilm/genms/protoc-gen-genms"
	_ "honnef.co/go/tools/cmd/staticcheck"
)

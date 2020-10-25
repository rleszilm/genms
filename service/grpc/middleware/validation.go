package middleware_grpc

import (
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

// WithValidationUnary is a wrapper for validator.UnaryServerInterceptor
func WithValidationUnary() grpc.UnaryServerInterceptor {
	return grpc_validator.UnaryServerInterceptor()
}

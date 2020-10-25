package middleware_grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

// WithTimeoutUnary returns a UnaryServerInterceptor that aplies a default timeout.
func WithTimeoutUnary(dur time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, cancel := context.WithTimeout(ctx, dur)
		defer cancel()

		return handler(ctx, req)
	}
}

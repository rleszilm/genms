package middleware_grpc

import (
	"context"
	"log"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// WithStatsUnary returns a UnaryServerInterceptor that reports stats.
func WithStatsUnary() grpc.UnaryServerInterceptor {
	if err := registerMetrics(); err != nil {
		log.Panicln("Unable to register metrics", err)
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		_, service, method := parseMethod(info.FullMethod)

		ctx, err := tag.New(ctx,
			tag.Upsert(requestService, service),
			tag.Upsert(requestMethod, method),
		)
		if err != nil {
			return nil, err
		}
		err = nil

		var resp interface{}
		stats.Record(ctx, requestsInflight.M(1))
		start := time.Now()
		defer func(ctx context.Context) {
			stop := time.Now()
			dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

			code := status.Code(err)
			ctx, err := tag.New(ctx,
				tag.Upsert(requestCode, code.String()),
			)
			if err != nil {
				log.Println("Error when setting response code in metrics", err, requestCode, requestCode.Name(), code.String(), len(code.String()))
			}

			stats.Record(ctx, requests.M(dur), requestsInflight.M(-1))
		}(ctx)

		resp, err = handler(ctx, req)
		return resp, err
	}
}

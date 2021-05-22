package middleware_grpc

import (
	"context"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// WithStatsUnary returns a UnaryServerInterceptor that reports stats.
func WithStatsUnary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		_, service, method := parseMethod(info.FullMethod)

		ctx, err := tag.New(ctx,
			tag.Upsert(tagReqService, service),
			tag.Upsert(tagReqMethod, method),
		)
		if err != nil {
			logs.Error("could not apply base metrics:", err)
			return nil, err
		}

		var resp interface{}
		stats.Record(ctx, measureReqInflight.M(1))
		start := time.Now()
		defer func(ctx context.Context) {
			stop := time.Now()
			dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

			code := status.Code(err)
			ctx, err := tag.New(ctx,
				tag.Upsert(tagReqCode, code.String()),
			)
			if err != nil {
				logs.Error("could not apply response code to metrics:", err, tagReqCode, tagReqCode.Name(), code.String(), len(code.String()))
			}

			stats.Record(ctx, measureReqLatency.M(dur), measureReqInflight.M(-1))
		}(ctx)

		resp, err = handler(ctx, req)
		if err != nil {
			stats.Record(ctx, measureError.M(1))
		}
		return resp, err
	}
}

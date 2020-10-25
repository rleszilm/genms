package middleware_grpc

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	requestService   tag.Key
	requestMethod    tag.Key
	requestCode      tag.Key
	requests         = stats.Float64("request_latency", "Latency of requests", stats.UnitMilliseconds)
	requestsInflight = stats.Int64("request_inflight", "Number of requests in flight", stats.UnitDimensionless)
)

func registerMetrics() error {
	var err error
	requestService, err = tag.NewKey("service")
	if err != nil {
		return err
	}

	requestMethod, err = tag.NewKey("method")
	if err != nil {
		return err
	}

	requestCode, err = tag.NewKey("code")
	if err != nil {
		return err
	}

	views := []*view.View{
		{
			Name:        "request_latency",
			Measure:     requests,
			Description: "The distribution of the request call latencies",
			TagKeys:     []tag.Key{requestMethod, requestService, requestCode},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "request_inflight",
			Measure:     requestsInflight,
			Description: "The number of requests being processed",
			TagKeys:     []tag.Key{requestMethod, requestService},
			Aggregation: view.Sum(),
		},
	}

	if err := view.Register(views...); err != nil {
		return err
	}
	return nil
}

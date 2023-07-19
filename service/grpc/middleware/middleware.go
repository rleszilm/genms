package middleware_grpc

import (
	"strings"

	"github.com/rleszilm/genms/logging"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	logs = logging.NewChannel("grpc-mid")
)

var (
	tagReqService = tag.MustNewKey("genms_request_service")
	tagReqMethod  = tag.MustNewKey("genms_request_method")
	tagReqCode    = tag.MustNewKey("genms_request_code")

	measureError       = stats.Int64("genms_request_error", "Count of request errors", stats.UnitDimensionless)
	measureReqLatency  = stats.Float64("genms_request_latency", "Latency of requests", stats.UnitMilliseconds)
	measureReqInflight = stats.Int64("genms_request_inflight", "Count of requests in flight", stats.UnitDimensionless)
)

func init() {
	views := []*view.View{
		{
			Name:        "genms_request_error",
			Measure:     measureError,
			Description: "Count of request errors",
			TagKeys:     []tag.Key{tagReqService, tagReqMethod, tagReqCode},
			Aggregation: view.Count(),
		},
		{
			Name:        "genms_request_latency",
			Measure:     measureReqLatency,
			Description: "The distribution of the request latencies",
			TagKeys:     []tag.Key{tagReqService, tagReqMethod, tagReqCode},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "genms_request_inflight",
			Measure:     measureReqInflight,
			Description: "The number of requests being processed",
			TagKeys:     []tag.Key{tagReqService, tagReqMethod, tagReqCode},
			Aggregation: view.LastValue(),
		},
	}

	if err := view.Register(views...); err != nil {
		logs.Fatal("Cannot register metrics:", err)
	}
}

func parseMethod(method string) (string, string, string) {
	// /package.service/method
	methodTokens := strings.Split(method, "/")
	packageTokens := strings.Split(methodTokens[1], ".")
	return packageTokens[0], packageTokens[1], methodTokens[2]
}

package http_service

import (
	"github.com/rleszilm/genms/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// define metrics
var (
	TagInstance     = tag.MustNewKey("genms_http_instance")
	TagMethod       = tag.MustNewKey("genms_http_method")
	TagEndpoint     = tag.MustNewKey("genms_http_endpoint")
	TagResponseCode = tag.MustNewKey("genms_http_response_code")

	MeasureError    = stats.Int64("genms_http_error", "Count of cache errors", stats.UnitDimensionless)
	MeasureLatency  = stats.Float64("genms_http_latency", "Latency of queries", stats.UnitMilliseconds)
	MeasureInflight = stats.Int64("genms_http_inflight", "Count of queries in flight", stats.UnitDimensionless)
)

func init() {
	views := []*view.View{
		{
			Name:        "genms_http_error",
			Measure:     MeasureError,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{TagInstance, TagMethod, TagEndpoint, TagResponseCode},
			Aggregation: view.Count(),
		},
		{
			Name:        "genms_http_latency",
			Measure:     MeasureLatency,
			Description: "The distribution of the query latencies",
			TagKeys:     []tag.Key{TagInstance, TagMethod, TagEndpoint, TagResponseCode},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "genms_http_inflight",
			Measure:     MeasureInflight,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{TagInstance, TagMethod, TagEndpoint, TagResponseCode},
			Aggregation: view.LastValue(),
		},
	}

	if err := view.Register(views...); err != nil {
		log.Fatal("Cannot register metrics:", err)
	}
}

package rest_service

import (
	"github.com/rleszilm/genms/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// define metrics
var (
	TagCollection   = tag.MustNewKey("genms_rest_collection")
	TagInstance     = tag.MustNewKey("genms_rest_instance")
	TagMethod       = tag.MustNewKey("genms_rest_method")
	TagRestMethod   = tag.MustNewKey("genms_rest_rest_method")
	TagResponseCode = tag.MustNewKey("genms_rest_response_code")

	MeasureError    = stats.Int64("genms_rest_error", "Count of cache errors", stats.UnitDimensionless)
	MeasureLatency  = stats.Float64("genms_rest_latency", "Latency of queries", stats.UnitMilliseconds)
	MeasureInflight = stats.Int64("genms_rest_inflight", "Count of queries in flight", stats.UnitDimensionless)

	logs = log.NewChannel("genms-rest")
)

func init() {
	views := []*view.View{
		{
			Name:        "genms_rest_error",
			Measure:     MeasureError,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagMethod, TagRestMethod, TagResponseCode},
			Aggregation: view.Count(),
		},
		{
			Name:        "genms_rest_latency",
			Measure:     MeasureLatency,
			Description: "The distribution of the query latencies",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagMethod, TagRestMethod, TagResponseCode},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "genms_rest_inflight",
			Measure:     MeasureInflight,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagMethod, TagRestMethod, TagResponseCode},
			Aggregation: view.LastValue(),
		},
	}

	if err := view.Register(views...); err != nil {
		logs.Fatal("Cannot register metrics:", err)
	}
}

// Logs returns the cache logs channel.
func Logs() *log.Channel {
	return logs
}

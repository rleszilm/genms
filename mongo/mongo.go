package mongo

import (
	"github.com/rleszilm/genms/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// define metrics
var (
	TagCollection   = tag.MustNewKey("genms_mongo_collection")
	TagInstance     = tag.MustNewKey("genms_mongo_instance")
	TagMethod       = tag.MustNewKey("genms_mongo_method")
	TagResponseCode = tag.MustNewKey("genms_mongo_response_code")

	MeasureError    = stats.Int64("genms_mongo_error", "Count of cache errors", stats.UnitDimensionless)
	MeasureLatency  = stats.Float64("genms_mongo_latency", "Latency of TypeOne queries", stats.UnitMilliseconds)
	MeasureInflight = stats.Int64("genms_mongo_inflight", "Count of TypeOne queries in flight", stats.UnitDimensionless)

	logs = log.NewChannel("genms-mongo")
)

func init() {
	views := []*view.View{
		{
			Name:        "genms_mongo_error",
			Measure:     MeasureError,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagMethod, TagResponseCode},
			Aggregation: view.Count(),
		},
		{
			Name:        "genms_mongo_latency",
			Measure:     MeasureLatency,
			Description: "The distribution of the query latencies",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagMethod, TagResponseCode},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "genms_mongo_inflight",
			Measure:     MeasureInflight,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagMethod, TagResponseCode},
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

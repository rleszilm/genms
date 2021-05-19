package sql

import (
	"github.com/rleszilm/genms/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// define metrics
var (
	TagCollection = tag.MustNewKey("genms_sql_collection")
	TagInstance   = tag.MustNewKey("genms_sql_instance")
	TagQuery      = tag.MustNewKey("genms_sql_query")
	TagDriver     = tag.MustNewKey("genms_sql_driver")

	MeasureError    = stats.Int64("genms_sql_error", "Count of cache errors", stats.UnitDimensionless)
	MeasureLatency  = stats.Float64("genms_sql_latency", "Latency of TypeOne queries", stats.UnitMilliseconds)
	MeasureInflight = stats.Int64("genms_sql_inflight", "Count of TypeOne queries in flight", stats.UnitDimensionless)

	logs = log.NewChannel("genms-sql")
)

func init() {
	views := []*view.View{
		{
			Name:        "genms_sql_inflight",
			Measure:     MeasureInflight,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagQuery, TagDriver},
			Aggregation: view.Count(),
		},
		{
			Name:        "genms_sql_latency",
			Measure:     measureLatency,
			Description: "The distribution of the query latencies",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagQuery, TagDriver},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "genms_sql_inflight",
			Measure:     measureInflight,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagQuery, TagDriver},
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

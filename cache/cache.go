package cache

import (
	"github.com/rleszilm/genms/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	TagCacheCollection = tag.MustNewKey("cache_collection")
	TagCacheInstance   = tag.MustNewKey("cache_instance")
	TagCacheMethod     = tag.MustNewKey("cache_method")
	TagCacheType       = tag.MustNewKey("cache_type")

	MeasureError           = stats.Int64("measure_error", "Count of cache errors", stats.UnitDimensionless)
	MeasureHit             = stats.Int64("measure_hit", "Count of cache hits", stats.UnitDimensionless)
	MeasureInflight        = stats.Int64("measure_inflight", "Count of cache lookups", stats.UnitDimensionless)
	MeasureLatency         = stats.Float64("measure_latency", "Latency of cache lookups", stats.UnitMilliseconds)
	MeasureMiss            = stats.Int64("measure_miss", "Count of cache misses", stats.UnitDimensionless)
	MeasureUpdatesInflight = stats.Int64("measure_inflight", "Count of cache updates in flight", stats.UnitDimensionless)
	MeasureUpdatesLatency  = stats.Float64("measure_latency", "Latency of cache updates.", stats.UnitMilliseconds)

	logs = log.NewChannel("cache")
)

func init() {
	views := []*view.View{
		{
			Name:        "cache_error",
			Measure:     MeasureError,
			Description: "Count of cache operations that resulted in an error.",
			TagKeys:     []tag.Key{TagCacheCollection, TagCacheInstance, TagCacheType, TagCacheMethod},
			Aggregation: view.Sum(),
		},
		{
			Name:        "cache_hit",
			Measure:     MeasureHit,
			Description: "Count of cache lookups where a value was present.",
			TagKeys:     []tag.Key{TagCacheCollection, TagCacheInstance, TagCacheType, TagCacheMethod},
			Aggregation: view.Sum(),
		},
		{
			Name:        "cache_miss",
			Measure:     MeasureMiss,
			Description: "Count of cache lookups where a value was not present.",
			TagKeys:     []tag.Key{TagCacheCollection, TagCacheInstance, TagCacheType, TagCacheMethod},
			Aggregation: view.Sum(),
		},
		{
			Name:        "cache_latency",
			Measure:     MeasureLatency,
			Description: "The distribution of cache lookup latencies.",
			TagKeys:     []tag.Key{TagCacheCollection, TagCacheInstance, TagCacheType, TagCacheMethod},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "cache_inflight",
			Measure:     MeasureInflight,
			Description: "The number of cache lookups being processed",
			TagKeys:     []tag.Key{TagCacheCollection, TagCacheInstance, TagCacheType, TagCacheMethod},
			Aggregation: view.Sum(),
		},
		{
			Name:        "cache_update_latency",
			Measure:     MeasureUpdatesLatency,
			Description: "The distribution of cache update latencies.",
			TagKeys:     []tag.Key{TagCacheCollection, TagCacheInstance, TagCacheType, TagCacheMethod},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "cache_update_inflight",
			Measure:     MeasureUpdatesInflight,
			Description: "The number of cache updates being processed",
			TagKeys:     []tag.Key{TagCacheCollection, TagCacheInstance, TagCacheType, TagCacheMethod},
			Aggregation: view.Sum(),
		},
	}

	if err := view.Register(views...); err != nil {
		logs.Fatal("Cannot register metrics:", err)
	}
}

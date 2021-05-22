package cache

import (
	"errors"

	"github.com/rleszilm/genms/log"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	logs = log.NewChannel("genms-cache")

	ErrAll      = errors.New("cache: cannot get all values")
	ErrGetNone  = errors.New("cache: no value for key")
	ErrGetValue = errors.New("cache: cannot get value from cache")
	ErrSetValue = errors.New("cache: cannot set value in cache")

	TagCollection = tag.MustNewKey("genms_cache_collection")
	TagInstance   = tag.MustNewKey("genms_cache_instance")
	TagMethod     = tag.MustNewKey("genms_cache_method")
	TagType       = tag.MustNewKey("genms_cache_type")

	MeasureError           = stats.Int64("genms_cache_error", "Count of cache errors", stats.UnitDimensionless)
	MeasureHit             = stats.Int64("genms_cache_hit", "Count of cache hits", stats.UnitDimensionless)
	MeasureInflight        = stats.Int64("genms_cache_inflight", "Count of cache lookups", stats.UnitDimensionless)
	MeasureLatency         = stats.Float64("genms_cache_latency", "Latency of cache lookups", stats.UnitMilliseconds)
	MeasureMiss            = stats.Int64("genms_cache_miss", "Count of cache misses", stats.UnitDimensionless)
	MeasureUpdatesInflight = stats.Int64("genms_cache_inflight", "Count of cache updates in flight", stats.UnitDimensionless)
	MeasureUpdatesLatency  = stats.Float64("genms_cache_latency", "Latency of cache updates.", stats.UnitMilliseconds)
)

func init() {
	views := []*view.View{
		{
			Name:        "genms_cache_error",
			Measure:     MeasureError,
			Description: "Count of cache operations that resulted in an error.",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagType, TagMethod},
			Aggregation: view.Count(),
		},
		{
			Name:        "genms_cache_hit",
			Measure:     MeasureHit,
			Description: "Count of cache lookups where a value was present.",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagType, TagMethod},
			Aggregation: view.Count(),
		},
		{
			Name:        "genms_cache_miss",
			Measure:     MeasureMiss,
			Description: "Count of cache lookups where a value was not present.",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagType, TagMethod},
			Aggregation: view.Count(),
		},
		{
			Name:        "genms_cache_latency",
			Measure:     MeasureLatency,
			Description: "The distribution of cache lookup latencies.",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagType, TagMethod},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "genms_cache_inflight",
			Measure:     MeasureInflight,
			Description: "The number of cache lookups being processed",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagType, TagMethod},
			Aggregation: view.LastValue(),
		},
		{
			Name:        "genms_cache_update_latency",
			Measure:     MeasureUpdatesLatency,
			Description: "The distribution of cache update latencies.",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagType, TagMethod},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "genms_cache_update_inflight",
			Measure:     MeasureUpdatesInflight,
			Description: "The number of cache updates being processed",
			TagKeys:     []tag.Key{TagCollection, TagInstance, TagType, TagMethod},
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

// Package cache_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_multi

import (
	context "context"
	time "time"

	cache "github.com/rleszilm/genms/cache"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal/keyvalue"
	service "github.com/rleszilm/genms/service"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// TypeTwoUpdater is an autogenerated implementation of dal.TypeTwoUpdater.
type TypeTwoUpdater struct {
	service.Dependencies

	name     string
	reader   keyvalue.TypeTwoReadAller
	writer   keyvalue.TypeTwoWriter
	key      keyvalue.TypeTwoKeyFunc
	interval time.Duration
	done     chan struct{}
}

// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *TypeTwoUpdater) Initialize(ctx context.Context) error {
	go x.update(ctx)
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *TypeTwoUpdater) Shutdown(_ context.Context) error {
	return nil
}

// String returns the name of the updater.
func (x *TypeTwoUpdater) String() string {
	if x.name != "" {
		return "cache-dal-multi-type-two-updater-" + x.name
	}
	return "cache-dal-multi-type-two-updater"
}

// NameOf returns the name of the updater.
func (x *TypeTwoUpdater) NameOf() string {
	return x.String()
}

func (x *TypeTwoUpdater) update(ctx context.Context) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_two"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "update"),
		tag.Upsert(cache.TagType, "updater"),
	)

	ticker := time.NewTicker(1)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			cache.Logs().Infof("starting update for %s", x.name)
			start := time.Now()
			stats.Record(ctx, cache.MeasureInflight.M(1))

			vals, err := x.reader.All(ctx)
			if err != nil {
			}

			for _, val := range vals {
				cache.Logs().Trace("updater TypeTwo storing value:", x.key(val), val)
				if _, err = x.writer.SetByKey(ctx, x.key(val), val); err != nil {
					cache.Logs().Error("updater TypeTwo could not store value:", x.key(val), val, err)
					break
				}
			}

			stats.Record(ctx, cache.MeasureInflight.M(-1))

			if err != nil {
				stats.Record(ctx, cache.MeasureError.M(1))
			}

			stop := time.Now()
			dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
			stats.Record(ctx, cache.MeasureLatency.M(dur))

			if x.interval == 0 {
				cache.Logs().Infof("updater %s is terminating", x.name)
				return
			}
			cache.Logs().Infof("scheduling next update for %v", x.interval)
			ticker.Reset(x.interval)
		}
	}
}

// WithReadAller tells the TypeTwoMap where to source values from if they don't exist in cache.
func (x *TypeTwoUpdater) WithReadAller(r keyvalue.TypeTwoReadAller) {
	x.reader = r
}

// WithWriter tells the TypeTwoMap where to source values from if they don't exist in cache.
func (x *TypeTwoUpdater) WithWriter(w keyvalue.TypeTwoWriter) {
	x.writer = w
}

// NewTypeTwoUpdater returns a new TypeTwoUpdater.
func NewTypeTwoUpdater(name string, k keyvalue.TypeTwoKeyFunc, i time.Duration) *TypeTwoUpdater {
	return &TypeTwoUpdater{
		name:     name,
		key:      k,
		interval: i,
		done:     make(chan struct{}),
	}
}

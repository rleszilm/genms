// Package cache_dal_single is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_single

import (
	context "context"
	time "time"

	golang_lru "github.com/hashicorp/golang-lru"
	cache "github.com/rleszilm/genms/cache"
	single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single/dal/keyvalue"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// SingleLRU defines a LRU cache implementing keyvalue.SingleReadWriter.
// If a key is queries that does not exist an attempt to read and store it is made.
type SingleLRU struct {
	name   string
	reader keyvalue.SingleReader
	writer keyvalue.SingleWriter
	lru    *golang_lru.ARCCache
	all    []*single.Single
}

// All implements implements keyvalue.SingleReadAller.
func (x *SingleLRU) All(ctx context.Context) ([]*single.Single, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "single"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "all"),
		tag.Upsert(cache.TagType, "lru"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	return x.all, nil
}

// GetByKey implements keyvalue.SingleReader.
func (x *SingleLRU) GetByKey(ctx context.Context, key keyvalue.SingleKey) (*single.Single, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "single"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "get"),
		tag.Upsert(cache.TagType, "lru"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	if val, ok := x.lru.Get(key); ok {
		stats.Record(ctx, cache.MeasureHit.M(1))
		return val.(*single.Single), nil
	}
	stats.Record(ctx, cache.MeasureMiss.M(1))

	if x.reader != nil {
		val, err := x.reader.GetByKey(ctx, key)
		if err != nil {
			return nil, err
		}
		x.lru.Add(key, val)
		return val, nil
	}

	stats.Record(ctx, cache.MeasureError.M(1))
	return nil, nil
}

// SetByKey implements keyvalue.SingleWriter.
func (x *SingleLRU) SetByKey(ctx context.Context, key keyvalue.SingleKey, val *single.Single) error {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "single"),
		tag.Upsert(cache.TagInstance, x.name),
		tag.Upsert(cache.TagMethod, "set"),
		tag.Upsert(cache.TagType, "lru"),
	)
	stats.Record(ctx, cache.MeasureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, cache.MeasureLatency.M(dur), cache.MeasureInflight.M(-1))
	}()

	if x.writer != nil {
		if err := x.writer.SetByKey(ctx, key, val); err != nil {
			stats.Record(ctx, cache.MeasureError.M(1))
			return err
		}
	}

	x.lru.Add(key, val)

	all := []*single.Single{}
	for _, k := range x.lru.Keys() {
		y, _ := x.lru.Get(k)
		all = append(all, y.(*single.Single))
	}
	x.all = all

	return nil
}

// WithReader tells the SingleLRU where to source values from if they don't exist in cache.
func (x *SingleLRU) WithReader(r keyvalue.SingleReader) {
	x.reader = r
}

// WithWriter tells the SingleLRU where to source values from if they don't exist in cache.
func (x *SingleLRU) WithWriter(w keyvalue.SingleWriter) {
	x.writer = w
}

// NewSingleLRU returns a new SingleLRU cache.
func NewSingleLRU(name string, i int) (*SingleLRU, error) {
	arc, err := golang_lru.NewARC(i)
	if err != nil {
		return nil, err
	}

	return &SingleLRU{
		name: name,
		lru:  arc,
	}, nil
}

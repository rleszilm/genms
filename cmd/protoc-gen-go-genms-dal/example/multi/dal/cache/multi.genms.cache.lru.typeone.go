// Package cache_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_multi

import (
	context "context"
	time "time"

	golang_lru "github.com/hashicorp/golang-lru"
	cache "github.com/rleszilm/genms/cache"
	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal/keyvalue"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// TypeOneLRU defines a LRU cache implementing keyvalue.TypeOneReadWriter.
// If a key is queries that does not exist an attempt to read and store it is made.
type TypeOneLRU struct {
	name   string
	reader keyvalue.TypeOneReader
	writer keyvalue.TypeOneWriter
	lru    *golang_lru.ARCCache
	all    []*multi.TypeOne
}

// All implements implements keyvalue.TypeOneReadAller.
func (x *TypeOneLRU) All(ctx context.Context) ([]*multi.TypeOne, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_one"),
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

// GetByKey implements keyvalue.TypeOneReadWriter.
func (x *TypeOneLRU) GetByKey(ctx context.Context, key keyvalue.TypeOneKey) (*multi.TypeOne, error) {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_one"),
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
		return val.(*multi.TypeOne), nil
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

// SetByKey implements keyvalue.TypeOneReadWriter.
func (x *TypeOneLRU) SetByKey(ctx context.Context, key keyvalue.TypeOneKey, val *multi.TypeOne) error {
	start := time.Now()
	ctx, _ = tag.New(ctx,
		tag.Upsert(cache.TagCollection, "type_one"),
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

	all := []*multi.TypeOne{}
	for _, k := range x.lru.Keys() {
		y, _ := x.lru.Get(k)
		all = append(all, y.(*multi.TypeOne))
	}
	x.all = all

	return nil
}

// WithReader tells the TypeOneLRU where to source values from if they don't exist in cache.
func (x *TypeOneLRU) WithReader(r keyvalue.TypeOneReader) {
	x.reader = r
}

// WithWriter tells the TypeOneLRU where to source values from if they don't exist in cache.
func (x *TypeOneLRU) WithWriter(w keyvalue.TypeOneWriter) {
	x.writer = w
}

// NewTypeOneLRU returns a new TypeOneLRU cache.
func NewTypeOneLRU(name string, i int) (*TypeOneLRU, error) {
	arc, err := golang_lru.NewARC(i)
	if err != nil {
		return nil, err
	}

	return &TypeOneLRU{
		name: name,
		lru:  arc,
	}, nil
}
package mongo_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-test/deep"
	pkgMongo "github.com/rleszilm/genms/mongo"
	pkgBson "github.com/rleszilm/genms/mongo/bson"
	"github.com/rleszilm/genms/mongo/pool"
)

type value struct {
	ID    pkgBson.ObjectID `bson:"_id"`
	Value string           `bson:"val"`
	Seed  int64            `bson:"seed"`
}

func TestReadWrite(t *testing.T) {
	ids := []string{
		"deadbeefdeadbeef00000000",
		"deadbeefdeadbeef00000001",
		"deadbeefdeadbeef00000002",
		"deadbeefdeadbeef00000003",
	}

	input := map[string]*value{}
	for _, id := range ids {
		oid, err := pkgBson.ObjectIDFromHex(id)
		if err != nil {
			t.Error(err)
			return
		}
		input[id] = &value{ID: oid, Value: id, Seed: time.Now().UnixNano()}
	}

	config := &pool.Config{
		URI:             "mongodb://mongo:27017",
		AppName:         "genms",
		MaxPoolSize:     25,
		MaxConnIdleTime: 30 * time.Second,
		Database:        "genms-test-data",
		Timeout:         5 * time.Second,
		ReadPref:        "primarypreferred",
	}

	dialer, err := pool.NewDialer(config)
	if err != nil {
		t.Error("no dials:", err)
		return
	}

	ctx := context.Background()
	if err := dialer.Initialize(ctx); err != nil {
		t.Error("no dialer:", err)
		return
	}

	client, err := dialer.Dial(ctx)
	if err != nil {
		t.Error("no client:", err)
		return
	}

	for _, v := range input {
		filter := pkgBson.M{
			"_id": v.ID,
		}
		update := pkgBson.M{
			"$set": v,
		}
		opts := &pkgMongo.UpdateOptions{}
		opts.SetUpsert(true)
		_, err := client.Database("genms-test-data").Collection("test-collection").UpdateOne(ctx, filter, update, opts)
		if err != nil {
			t.Error("no write:", err)
			return
		}
	}

	selector := pkgBson.M{
		"_id":  1,
		"val":  1,
		"seed": 1,
	}
	find := pkgBson.M{}
	cur, err := client.Database("genms-test-data").Collection("test-collection").Find(ctx, find, selector)
	if err != nil {
		t.Error("no results:", err)
		return
	}

	output := map[string]*value{}
	for cur.Next(ctx) {
		record := &value{}
		if err := cur.Decode(record); err != nil {
			t.Error("no decode:", err)
			return
		}

		output[record.ID.Hex()] = record
	}

	if diff := deep.Equal(output, input); diff != nil {
		t.Error("read values are not as they were written:", diff)
	}
}

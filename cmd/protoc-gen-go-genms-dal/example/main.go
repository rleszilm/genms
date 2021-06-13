package main

import (
	"context"
	"log"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
	cache_dal_single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single/dal/cache"
)

func main() {
	hash, err := cache_dal_single.NewSingleMap("map")
	if err != nil {
		log.Fatal(err)
	}

	lru, err := cache_dal_single.NewSingleLRU("lru", 5)
	if err != nil {
		log.Fatal(err)
	}
	lru.WithWriter(hash)
	lru.WithReader(hash)

	ctx := context.Background()
	for i := 0; i < 11; i++ {
		if err := lru.SetByKey(ctx, i, &single.Single{ScalarInt32: int32(i)}); err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 11; i++ {
		hval, err := hash.GetByKey(ctx, i)
		if err != nil {
			log.Fatal(err)
		}

		val, err := lru.GetByKey(ctx, i)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d - %+v - %+v", i, val, hval)
	}
}

// Package cache_dal_multi is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package cache_dal_multi

import (
	context "context"

	multi "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi"
	keyvalue "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/multi/dal/keyvalue"
)

// NilTypeTwoCache is a KV ReadWriter that takes no action on read or write.
type NilTypeTwoCache struct {
}

// GetAll implements keyvalue.TypeTwoReadAller.
func (x *NilTypeTwoCache) All(_ context.Context) (*multi.TypeTwo, error) {
	return nil, nil
}

// GetByKey implements keyvalue.TypeTwoReader.
func (x *NilTypeTwoCache) GetByKey(_ context.Context, _ keyvalue.TypeTwoKey) (*multi.TypeTwo, error) {
	return nil, nil
}

// SetByKey implements keyvalue.TypeTwoWriter.
func (x *NilTypeTwoCache) SetByKey(_ context.Context, _ keyvalue.TypeTwoKey, _ *multi.TypeTwo) error {
	return nil
}
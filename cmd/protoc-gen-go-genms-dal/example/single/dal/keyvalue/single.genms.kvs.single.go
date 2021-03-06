// Package keyvalue_dal_single is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package keyvalue_dal_single

import (
	context "context"

	single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
)

// SingleKey defines a Key in the kv store.
type SingleKey interface{}

// SingleReader is defines the interface for getting values from a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SingleReader
type SingleReader interface {
	GetByKey(context.Context, SingleKey) (*single.Single, error)
}

// SingleReadeAllr is defines the interface for getting values from a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SingleReadAller
type SingleReadAller interface {
	All(context.Context) ([]*single.Single, error)
}

// SingleWriter is defines the interface for setting values in a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SingleWriter
type SingleWriter interface {
	SetByKey(context.Context, SingleKey, *single.Single) (*single.Single, error)
}

// SingleReadWriter is defines the interface for setting values in a KV store.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SingleReadWriter
type SingleReadWriter interface {
	SingleReader
	SingleWriter
}

// SingleKeyFunc is a function that generates a unique deterministic key for the single.Single.
type SingleKeyFunc func(*single.Single) interface{}

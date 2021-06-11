package mongo

import (
	"context"

	"github.com/rleszilm/genms/service"
)

// Dialer is the interface used to obtain a new mongo client.
type Dialer interface {
	service.Service

	// Dial creates a client connected to mongo.
	Dial(context.Context) (Client, error)
}

// Cursor is an interface that mirrors the mongo driver Cursor struct.
type Cursor interface {
	Decode(interface{}) error
	Next(context.Context) bool
}

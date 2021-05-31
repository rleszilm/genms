package mongo

import (
	"context"
)

// Dialer is the interface used to obtain a new mongo client.
type Dialer interface {
	// Initialize starts up the underlying mongo driver and ensures connectivity.
	Initialize(context.Context) error
	// Shutdown shuts down the underlying mongo driver.
	Shutdown(context.Context) error
	// Dial creates a client connected to mongo.
	Dial(context.Context) (Client, error)
}

// Cursor is an interface that mirrors the mongo driver Cursor struct.
type Cursor interface {
	Decode(interface{}) error
	Next(context.Context) bool
}

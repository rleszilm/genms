package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client is the interface for a mongo client.
type Client interface {
	// Database returns a reference to the mongo database.
	Database(db string, opts ...*DatabaseOptions) Database
	// Close closes the client connection to mongo.
	Close(ctx context.Context) error
	// Ping pings the remote mongo server to ensure connectivity is valid.
	Ping(ctx context.Context, rp *readpref.ReadPref) error
}

// Client is a wrapper for mongo.Client
type SimpleClient struct {
	session mongo.Session
}

// Database implements mongo.Client.Database
func (c *SimpleClient) Database(db string, opts ...*DatabaseOptions) Database {
	dOpts := []*options.DatabaseOptions{}
	for _, opt := range opts {
		dOpts = append(dOpts, &opt.DatabaseOptions)
	}
	return &SimpleDatabase{Database: c.session.Client().Database(db, dOpts...)}
}

// Close implements mongo.Client.Close
func (c *SimpleClient) Close(ctx context.Context) error {
	c.session.EndSession(ctx)
	return nil
}

// Ping implements mongo.Client.Ping
func (c *SimpleClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.session.Client().Ping(ctx, rp)
}

func NewSimpleClient(s mongo.Session) *SimpleClient {
	return &SimpleClient{session: s}
}

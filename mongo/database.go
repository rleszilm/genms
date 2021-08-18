package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database is an interface that mirrors the mongo driver Database struct.
type Database interface {
	Client() Client
	Name() string
	Collection(name string, opts ...*CollectionOptions) Collection
	RunCommand(ctx context.Context, cmd interface{}, opts ...*RunCmdOptions) SingleResult
	RunCommandCursor(ctx context.Context, cmd interface{}, opts ...*RunCmdOptions) (Cursor, error)
	Drop(ctx context.Context) error
}

// SimpleDatabase is a wrapper for mongo.Database
type SimpleDatabase struct {
	*mongo.Database
}

func (d *SimpleDatabase) Client() Client {
	return &SimpleClient{}
}

// Collection implements mongo.Database.Collection
func (d *SimpleDatabase) Collection(col string, opts ...*CollectionOptions) Collection {
	dOpts := []*options.CollectionOptions{}
	for _, opt := range opts {
		dOpts = append(dOpts, &opt.CollectionOptions)
	}

	return &SimpleCollection{Collection: d.Database.Collection(col, dOpts...)}
}

// RunCommand implements mongo.Database.RunCommand
func (d *SimpleDatabase) RunCommand(ctx context.Context, runCommand interface{}, opts ...*RunCmdOptions) SingleResult {
	dOpts := []*options.RunCmdOptions{}
	for _, opt := range opts {
		dOpts = append(dOpts, &opt.RunCmdOptions)
	}

	res := d.Database.RunCommand(ctx, runCommand, dOpts...)
	return &SimpleSingleResult{SingleResult: res}
}

// RunCommandCursor implements mongo.Database.RunCommandCursor
func (d *SimpleDatabase) RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*RunCmdOptions) (Cursor, error) {
	dOpts := []*options.RunCmdOptions{}
	for _, opt := range opts {
		dOpts = append(dOpts, &opt.RunCmdOptions)
	}

	return d.Database.RunCommandCursor(ctx, runCommand, dOpts...)
}

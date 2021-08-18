package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WriteModel interface {
	mongo.WriteModel
}

// Collection is an interface that mirrors the mongo driver Collection struct.
type Collection interface {
	Clone(...*CollectionOptions) (Collection, error)
	Name() string
	Database() Database
	BulkWrite(ctx context.Context, models []WriteModel, opts ...*BulkWriteOptions) (*BulkWriteResult, error)
	InsertOne(ctx context.Context, obj interface{}, opts ...*InsertOneOptions) (*InsertOneResult, error)
	InsertMany(ctx context.Context, objs []interface{}, opts ...*InsertManyOptions) (*InsertManyResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*DeleteOptions) (*DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*DeleteOptions) (*DeleteResult, error)
	UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*UpdateOptions) (*UpdateResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*UpdateOptions) (*UpdateResult, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*UpdateOptions) (*UpdateResult, error)
	ReplaceOne(ctx context.Context, filter interface{}, obj interface{}, opts ...*ReplaceOptions) (*UpdateResult, error)
	Find(ctx context.Context, filter interface{}, projection interface{}, opts ...*FindOptions) (Cursor, error)
	FindOne(ctx context.Context, filter interface{}, projection interface{}, opts ...*FindOneOptions) SingleResult
	FindOneAndDelete(ctx context.Context, filter interface{}, projection interface{}, opts ...*FindOneAndDeleteOptions) SingleResult
	FindOneAndReplace(ctx context.Context, filter interface{}, projection interface{}, obj interface{}, opts ...*FindOneAndReplaceOptions) SingleResult
	FindOneAndUpdate(ctx context.Context, filter interface{}, projection interface{}, obj interface{}, opts ...*FindOneAndUpdateOptions) SingleResult
}

// SimpleCollection is a wrapper for mongo.Collection
type SimpleCollection struct {
	*mongo.Collection
}

// Clone implements mongo.Collection.Clone
func (c *SimpleCollection) Clone(opts ...*CollectionOptions) (Collection, error) {
	cOpts := []*options.CollectionOptions{}
	for _, opt := range opts {
		cOpts = append(cOpts, &opt.CollectionOptions)
	}

	col, err := c.Collection.Clone(cOpts...)
	if err != nil {
		return nil, err
	}

	return &SimpleCollection{Collection: col}, nil
}

// Database implements mongo.Collection.Database
func (c *SimpleCollection) Database() Database {
	return &SimpleDatabase{Database: c.Collection.Database()}
}

// BulkWrite implements mongo.Collection.BulkWrite
func (c *SimpleCollection) BulkWrite(ctx context.Context, wms []WriteModel, opts ...*BulkWriteOptions) (*BulkWriteResult, error) {
	bwOpts := []*options.BulkWriteOptions{}
	for _, opt := range opts {
		bwOpts = append(bwOpts, &opt.BulkWriteOptions)
	}

	mongoWMs := []mongo.WriteModel{}
	for _, wm := range wms {
		mongoWMs = append(mongoWMs, wm)
	}

	res, err := c.Collection.BulkWrite(ctx, mongoWMs, bwOpts...)
	if err != nil {
		return nil, err
	}

	return &BulkWriteResult{BulkWriteResult: res}, err
}

// InsertOne implements mongo.Collection.insertOne
func (c *SimpleCollection) InsertOne(ctx context.Context, obj interface{}, opts ...*InsertOneOptions) (*InsertOneResult, error) {
	iOpts := []*options.InsertOneOptions{}
	for _, opt := range opts {
		iOpts = append(iOpts, &opt.InsertOneOptions)
	}

	res, err := c.Collection.InsertOne(ctx, obj, iOpts...)
	if err != nil {
		return nil, err
	}

	return &InsertOneResult{InsertOneResult: res}, nil
}

// InsertMany implements mongo.Collection.insertMany
func (c *SimpleCollection) InsertMany(ctx context.Context, objs []interface{}, opts ...*InsertManyOptions) (*InsertManyResult, error) {
	iOpts := []*options.InsertManyOptions{}
	for _, opt := range opts {
		iOpts = append(iOpts, &opt.InsertManyOptions)
	}

	res, err := c.Collection.InsertMany(ctx, objs, iOpts...)
	if err != nil {
		return nil, err
	}

	return &InsertManyResult{InsertManyResult: res}, nil
}

// DeleteOne implements mongo.Collection.DeleteOne
func (c *SimpleCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*DeleteOptions) (*DeleteResult, error) {
	iOpts := []*options.DeleteOptions{}
	for _, opt := range opts {
		iOpts = append(iOpts, &opt.DeleteOptions)
	}

	res, err := c.Collection.DeleteOne(ctx, filter, iOpts...)
	if err != nil {
		return nil, err
	}

	return &DeleteResult{DeleteResult: res}, nil
}

// DeleteMany implements mongo.Collection.DeleteMany
func (c *SimpleCollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*DeleteOptions) (*DeleteResult, error) {
	iOpts := []*options.DeleteOptions{}
	for _, opt := range opts {
		iOpts = append(iOpts, &opt.DeleteOptions)
	}

	res, err := c.Collection.DeleteMany(ctx, filter, iOpts...)
	if err != nil {
		return nil, err
	}

	return &DeleteResult{DeleteResult: res}, nil
}

// UpdateByID implements mongo.Collection.UpdateByID
func (c *SimpleCollection) UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*UpdateOptions) (*UpdateResult, error) {
	uOpts := []*options.UpdateOptions{}
	for _, opt := range opts {
		uOpts = append(uOpts, &opt.UpdateOptions)
	}

	res, err := c.Collection.UpdateByID(ctx, id, update, uOpts...)
	if err != nil {
		return nil, err
	}

	return &UpdateResult{UpdateResult: res}, nil
}

// UpdateOne implements mongo.Collection.UpdateOne
func (c *SimpleCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*UpdateOptions) (*UpdateResult, error) {
	uOpts := []*options.UpdateOptions{}
	for _, opt := range opts {
		uOpts = append(uOpts, &opt.UpdateOptions)
	}

	res, err := c.Collection.UpdateOne(ctx, filter, update, uOpts...)
	if err != nil {
		return nil, err
	}

	return &UpdateResult{UpdateResult: res}, nil
}

// UpdateMany implements mongo.Collection.UpdateMany
func (c *SimpleCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*UpdateOptions) (*UpdateResult, error) {
	uOpts := []*options.UpdateOptions{}
	for _, opt := range opts {
		uOpts = append(uOpts, &opt.UpdateOptions)
	}

	res, err := c.Collection.UpdateMany(ctx, filter, update, uOpts...)
	if err != nil {
		return nil, err
	}

	return &UpdateResult{UpdateResult: res}, nil
}

// ReplaceOne implements mongo.Collection.ReplaceOne
func (c *SimpleCollection) ReplaceOne(ctx context.Context, filter interface{}, update interface{}, opts ...*ReplaceOptions) (*UpdateResult, error) {
	rOpts := []*options.ReplaceOptions{}
	for _, opt := range opts {
		rOpts = append(rOpts, &opt.ReplaceOptions)
	}

	res, err := c.Collection.ReplaceOne(ctx, filter, update, rOpts...)
	if err != nil {
		return nil, err
	}

	return &UpdateResult{UpdateResult: res}, nil
}

// Find implements mongo.Collection.Find
func (c *SimpleCollection) Find(ctx context.Context, filter interface{}, projection interface{}, opts ...*FindOptions) (Cursor, error) {
	projOpts := options.Find()
	projOpts.SetProjection(projection)

	fOpts := []*options.FindOptions{projOpts}
	for _, opt := range opts {
		fOpts = append(fOpts, &opt.FindOptions)
	}

	cur, err := c.Collection.Find(ctx, filter, fOpts...)
	if err != nil {
		return nil, err
	}

	return cur, nil
}

// FindOne implements mongo.Collection.FindOne
func (c *SimpleCollection) FindOne(ctx context.Context, filter interface{}, projection interface{}, opts ...*FindOneOptions) SingleResult {
	projOpts := options.FindOne()
	projOpts.SetProjection(projection)

	fOpts := []*options.FindOneOptions{projOpts}
	for _, opt := range opts {
		fOpts = append(fOpts, &opt.FindOneOptions)
	}

	res := c.Collection.FindOne(ctx, filter, fOpts...)
	return &SimpleSingleResult{SingleResult: res}
}

// FindOneAndDelete implements mongo.Collection.FindOneAndDelete
func (c *SimpleCollection) FindOneAndDelete(ctx context.Context, filter interface{}, projection interface{}, opts ...*FindOneAndDeleteOptions) SingleResult {
	projOpts := options.FindOneAndDelete()
	projOpts.SetProjection(projection)

	fOpts := []*options.FindOneAndDeleteOptions{projOpts}
	for _, opt := range opts {
		fOpts = append(fOpts, &opt.FindOneAndDeleteOptions)
	}

	res := c.Collection.FindOneAndDelete(ctx, filter, fOpts...)
	return &SimpleSingleResult{SingleResult: res}
}

// FindOneAndReplace implements mongo.Collection.FindOneAndReplace
func (c *SimpleCollection) FindOneAndReplace(ctx context.Context, filter interface{}, projection interface{}, replacement interface{}, opts ...*FindOneAndReplaceOptions) SingleResult {
	projOpts := options.FindOneAndReplace()
	projOpts.SetProjection(projection)

	fOpts := []*options.FindOneAndReplaceOptions{projOpts}
	for _, opt := range opts {
		fOpts = append(fOpts, &opt.FindOneAndReplaceOptions)
	}

	res := c.Collection.FindOneAndReplace(ctx, filter, replacement, fOpts...)
	return &SimpleSingleResult{SingleResult: res}
}

// FindOneAndUpdate implements mongo.Collection.FindOneAndUpdate
func (c *SimpleCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, projection interface{}, update interface{}, opts ...*FindOneAndUpdateOptions) SingleResult {
	projOpts := options.FindOneAndUpdate()
	projOpts.SetProjection(projection)

	fOpts := []*options.FindOneAndUpdateOptions{projOpts}
	for _, opt := range opts {
		fOpts = append(fOpts, &opt.FindOneAndUpdateOptions)
	}

	res := c.Collection.FindOneAndUpdate(ctx, filter, update, fOpts...)
	return &SimpleSingleResult{SingleResult: res}
}

package mongo

import "go.mongodb.org/mongo-driver/mongo/options"

// AggregateOptions is a wrapper for the mongo AggregateOptions
type AggregateOptions struct {
	options.AggregateOptions
}

// AutoEncryptionOptions is a wrapper for the mongo AutoEncryptionOptions
type AutoEncryptionOptions struct {
	options.AutoEncryptionOptions
}

// BucketOptions is a wrapper for the mongo BucketOptions
type BucketOptions struct {
	options.BucketOptions
}

// BulkWriteOptions is a wrapper for the mongo BulkWriteOptions
type BulkWriteOptions struct {
	options.BulkWriteOptions
}

// ChangeStreamOptions is a wrapper for the mongo ChangeStreamOptions
type ChangeStreamOptions struct {
	options.ChangeStreamOptions
}

// ClientEncryptionOptions is a wrapper for the mongo ClientEncryptionOptions
type ClientEncryptionOptions struct {
	options.ClientEncryptionOptions
}

// ClientOptions is a wrapper for the mongo ClientOptions
type ClientOptions struct {
	options.ClientOptions
}

// Collation is a wrapper for the mongo Collation
type Collation struct {
	options.Collation
}

// CollectionOptions is a wrapper for the mongo CollectionOptions
type CollectionOptions struct {
	options.CollectionOptions
}

// CountOptions is a wrapper for the mongo CountOptions
type CountOptions struct {
	options.CountOptions
}

// CreateCollectionOptions is a wrapper for the mongo CreateCollectionOptions
type CreateCollectionOptions struct {
	options.CreateCollectionOptions
}

// CreateIndexesOptions is a wrapper for the mongo CreateIndexesOptions
type CreateIndexesOptions struct {
	options.CreateIndexesOptions
}

// DataKeyOptions is a wrapper for the mongo DataKeyOptions
type DataKeyOptions struct {
	options.DataKeyOptions
}

// DatabaseOptions is a wrapper for the mongo DatabaseOptions
type DatabaseOptions struct {
	options.DatabaseOptions
}

// DeleteOptions is a wrapper for the mongo DeleteOptions
type DeleteOptions struct {
	options.DeleteOptions
}

// DistinctOptions is a wrapper for the mongo DistinctOptions
type DistinctOptions struct {
	options.DistinctOptions
}

// DropIndexesOptions is a wrapper for the mongo DropIndexesOptions
type DropIndexesOptions struct {
	options.DropIndexesOptions
}

// EncryptOptions is a wrapper for the mongo EncryptOptions
type EncryptOptions struct {
	options.EncryptOptions
}

// EstimatedDocumentCountOptions is a wrapper for the mongo EstimatedDocumentCountOptions
type EstimatedDocumentCountOptions struct {
	options.EstimatedDocumentCountOptions
}

// FindOptions is a wrapper for the mongo FindOptions
type FindOptions struct {
	options.FindOptions
}

// FindOneOptions is a wrapper for the mongo FindOneOptions
type FindOneOptions struct {
	options.FindOneOptions
}

// FindOneAndDeleteOptions is a wrapper for the mongo FindOneAndDeleteOptions
type FindOneAndDeleteOptions struct {
	options.FindOneAndDeleteOptions
}

// FindOneAndReplaceOptions is a wrapper for the mongo FindOneAndReplaceOptions
type FindOneAndReplaceOptions struct {
	options.FindOneAndReplaceOptions
}

// FindOneAndUpdateOptions is a wrapper for the mongo FindOneAndUpdateOptions
type FindOneAndUpdateOptions struct {
	options.FindOneAndUpdateOptions
}

// GridFSFindOptions is a wrapper for the mongo GridFSFindOptions
type GridFSFindOptions struct {
	options.GridFSFindOptions
}

// IndexOptions is a wrapper for the mongo IndexOptions
type IndexOptions struct {
	options.IndexOptions
}

// InsertManyOptions is a wrapper for the mongo InsertManyOptions
type InsertManyOptions struct {
	options.InsertManyOptions
}

// InsertOneOptions is a wrapper for the mongo InsertOneOptions
type InsertOneOptions struct {
	options.InsertOneOptions
}

// ListCollectionsOptions is a wrapper for the mongo ListCollectionsOptions
type ListCollectionsOptions struct {
	options.ListCollectionsOptions
}

// ListDatabasesOptions is a wrapper for the mongo ListDatabasesOptions
type ListDatabasesOptions struct {
	options.ListDatabasesOptions
}

// ListIndexesOptions is a wrapper for the mongo ListIndexesOptions
type ListIndexesOptions struct {
	options.ListIndexesOptions
}

// ReplaceOptions is a wrapper for the mongo ReplaceOptions
type ReplaceOptions struct {
	options.ReplaceOptions
}

// RunCmdOptions is a wrapper for the mongo RunCmdOptions
type RunCmdOptions struct {
	options.RunCmdOptions
}

// SessionOptions is a wrapper for the mongo SessionOptions
type SessionOptions struct {
	options.SessionOptions
}

// TransactionOptions is a wrapper for the mongo TransactionOptions
type TransactionOptions struct {
	options.TransactionOptions
}

// UpdateOptions is a wrapper for the mongo UpdateOptions
type UpdateOptions struct {
	options.UpdateOptions
}

// UploadOptions is a wrapper for the mongo UploadOptions
type UploadOptions struct {
	options.UploadOptions
}

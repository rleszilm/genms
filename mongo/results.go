package mongo

import (
	"github.com/rleszilm/genms/mongo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BulkWriteResult wraps mongo.BulkWriteResult
type BulkWriteResult struct {
	*mongo.BulkWriteResult
}

// InsertOneResult wraps mongo.InsertOneResult
type InsertOneResult struct {
	*mongo.InsertOneResult
}

// InsertManyResult wraps mongo.InsertManyResult
type InsertManyResult struct {
	*mongo.InsertManyResult
}

// DeleteResult wraps mongo.DeleteResult
type DeleteResult struct {
	*mongo.DeleteResult
}

// UpdateResult wraps mongo.UpdateResult
type UpdateResult struct {
	*mongo.UpdateResult
}

// SingleResult is an interface that mirrors the mongo driver SingleResult struct.
type SingleResult interface {
	Decode(obj interface{}) error
	DecodeBytes() (bson.Raw, error)
	Err() error
}

// SimpleSingleResult is a wrapper for mongo.SingleResult
type SimpleSingleResult struct {
	*mongo.SingleResult
}

// ResumeToken wraps cmongo.SingleResult.ResumeToken
func (s *SimpleSingleResult) DecodeBytes() (bson.Raw, error) {
	raw, err := s.SingleResult.DecodeBytes()
	if err != nil {
		return nil, err
	}
	return bson.Raw(raw), err
}

package bson

import "errors"

var (
	// ErrNonObjectID is returned if a value is not or cannot be converted into an ObjectID
	ErrNonObjectID = errors.New("not an objectid")
)

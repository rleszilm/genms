package mongo

import "errors"

var (
	// ErrNoID is returned when an auto-generated method is called for a type that does not
	// have a _id field defined.
	ErrNoID = errors.New("mongo: no _id")

	// ErrBadObjID is returned when an interface cannot be converted into a bson.ObjectID.
	ErrBadObjID = errors.New("mongo: bad obj id")
)

package bson

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectIDHex returns an ObjectID from the given hex string. It panics if the
// hexstring is not valid.
func ObjectIDHex(in string) ObjectID {
	out, err := primitive.ObjectIDFromHex(in)
	if err != nil {
		panic(err)
	}
	return ObjectID(out)
}

// NewObjectID generates a new ObjectID.
func NewObjectID() ObjectID {
	return NewObjectIDFromTimestamp(time.Now())
}

// NewObjectIDFromTimestamp generates a new ObjectID based on the given time.
func NewObjectIDFromTimestamp(timestamp time.Time) ObjectID {
	return ObjectID(primitive.NewObjectIDFromTimestamp(timestamp))
}

// Timestamp extracts the time part of the ObjectId.
func (id ObjectID) Timestamp() time.Time {
	return primitive.ObjectID(id).Timestamp()
}

// Hex returns the hex encoding of the ObjectID as a string.
func (id ObjectID) Hex() string {
	return primitive.ObjectID(id).Hex()
}

func (id ObjectID) String() string {
	return primitive.ObjectID(id).String()
}

// IsZero returns true if id is the empty ObjectID.
func (id ObjectID) IsZero() bool {
	return primitive.ObjectID(id).IsZero()
}

// ObjectIDFromHex creates a new ObjectID from a hex string. It returns an error if the hex string is not a
// valid ObjectID.
func ObjectIDFromHex(s string) (ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(s)
	return ObjectID(id), err
}

// IsValidObjectID returns true if the provided hex string represents a valid ObjectID and false if not.
func IsValidObjectID(s string) bool {
	return primitive.IsValidObjectID(s)
}

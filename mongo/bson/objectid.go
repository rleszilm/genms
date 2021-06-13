package bson

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectIDHex returns an ObjectID from the given hex string.
func ObjectIDHex(in string) (ObjectID, error) {
	out, err := primitive.ObjectIDFromHex(in)
	if err != nil {
		return ObjectID{}, err
	}
	return ObjectID(out), nil
}

// MustObjectIDHex panics if the provided hexstring is not valid.
func MustObjectIDHex(in string) ObjectID {
	out, err := ObjectIDHex(in)
	if err != nil {
		panic(err)
	}
	return out
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

// IsValidObjectID returns true if the provided hex string represents a valid ObjectID and false if not.
func IsValidObjectID(s string) bool {
	return primitive.IsValidObjectID(s)
}

// ToObjectID converts the given value into an ObjectID.
func ToObjectID(in interface{}) (ObjectID, error) {
	switch in.(type) {
	case ObjectID:
		return in.(ObjectID), nil
	case primitive.ObjectID:
		return ObjectID(in.(primitive.ObjectID)), nil
	case [12]byte:
		return ObjectID(in.([12]byte)), nil
	case []byte:
		oid := in.([]byte)
		if len(oid) != 12 {
			return ObjectID{}, ErrNonObjectID
		}

		out := [12]byte{}
		copy(out[:], oid)
		return ObjectID(out), nil
	case string:
		oid, err := primitive.ObjectIDFromHex(in.(string))
		if err != nil {
			return ObjectID{}, err
		}
		return ObjectID(oid), nil
	default:
		return ObjectID{}, ErrNonObjectID
	}
}

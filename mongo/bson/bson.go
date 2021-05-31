package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

var (
	dvd bsoncodec.DefaultValueDecoders
	dve bsoncodec.DefaultValueEncoders
	rb  *bsoncodec.RegistryBuilder
)

func init() {
	rb = bsoncodec.NewRegistryBuilder()

	// defaults
	dvd.RegisterDefaultDecoders(rb)
	dve.RegisterDefaultEncoders(rb)

	registerBinary(rb)
	registerCodeWithScope(rb)
	registerDBPointer(rb)
	registerDateTime(rb)
	registerDecimal128(rb)
	registerJavaScript(rb)
	registerMaxKey(rb)
	registerMinKey(rb)
	registerNull(rb)
	registerObjectID(rb)
	registerRegex(rb)
	registerSymbol(rb)
	registerTimestamp(rb)

	bson.DefaultRegistry = rb.Build()
}

// Raw exports bson.Raw
type Raw bson.Raw

// A exports bson.A
type A bson.A

// D exports bson.D
type D bson.D

// E exports bson.E
type E bson.E

// M exports bson.M
type M bson.M

// ValueCodec exports bsoncodec.ValueCodec
type ValueCodec interface {
	bsoncodec.ValueCodec
}

// WithValueCodec adds a type specific ValueCodec.
func WithValueCodec(k reflect.Type, c ValueCodec) {
	rb.RegisterCodec(k, c)
	bson.DefaultRegistry = rb.Build()
}

// Marshal exports bson.Marshal
func Marshal(in interface{}) ([]byte, error) {
	return bson.Marshal(in)
}

// Unmarshal exports bson.Unmarshal
func Unmarshal(in []byte, out interface{}) error {
	return bson.Unmarshal(in, out)
}

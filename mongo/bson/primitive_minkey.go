package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MinKey exports the driver specific MinKey so it can be used externally.
type MinKey primitive.MinKey

var (
	typeMinKey          reflect.Type
	typePrimitiveMinKey reflect.Type
)

func registerMinKey(rb *bsoncodec.RegistryBuilder) {
	var instanceMinKey MinKey
	var instancePrimitiveMinKey primitive.MinKey

	typeMinKey = reflect.TypeOf(instanceMinKey)
	typePrimitiveMinKey = reflect.TypeOf(instancePrimitiveMinKey)

	rb.RegisterTypeDecoder(typeMinKey, bsoncodec.ValueDecoderFunc(codecMinKeyDecodeValue))
	rb.RegisterTypeEncoder(typeMinKey, bsoncodec.ValueEncoderFunc(codecMinKeyEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.MinKey, typeMinKey)
}

// codecMinKeyEncodeValue is the ValueDecoderFunc for primitive.MinKey.
func codecMinKeyEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeMinKey {
		return bsoncodec.ValueDecoderError{Name: "Mongo MinKeyEncodeValue", Types: []reflect.Type{typeMinKey}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveMinKey)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveMinKey))
}

// codecMinKeyDecodeValue is the ValueDecoderFunc for primitive.MinKey.
func codecMinKeyDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeMinKey {
		return bsoncodec.ValueDecoderError{Name: "Mongo MinKeyDecodeValue", Types: []reflect.Type{typeMinKey}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveMinKey)
	if err != nil {
		return err
	}

	var new primitive.MinKey
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeMinKey))

	return nil
}

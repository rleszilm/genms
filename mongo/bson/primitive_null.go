package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Null exports the driver specific Null so it can be used externally.
type Null primitive.Null

var (
	typeNull          reflect.Type
	typePrimitiveNull reflect.Type
)

func registerNull(rb *bsoncodec.RegistryBuilder) {
	var instanceNull Null
	var instancePrimitiveNull primitive.Null

	typeNull = reflect.TypeOf(instanceNull)
	typePrimitiveNull = reflect.TypeOf(instancePrimitiveNull)

	rb.RegisterTypeDecoder(typeNull, bsoncodec.ValueDecoderFunc(codecNullDecodeValue))
	rb.RegisterTypeEncoder(typeNull, bsoncodec.ValueEncoderFunc(codecNullEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.Null, typeNull)
}

// codecNullEncodeValue is the ValueDecoderFunc for primitive.Null.
func codecNullEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeNull {
		return bsoncodec.ValueDecoderError{Name: "Mongo NullEncodeValue", Types: []reflect.Type{typeNull}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveNull)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveNull))
}

// codecNullDecodeValue is the ValueDecoderFunc for primitive.Null.
func codecNullDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeNull {
		return bsoncodec.ValueDecoderError{Name: "Mongo NullDecodeValue", Types: []reflect.Type{typeNull}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveNull)
	if err != nil {
		return err
	}

	var new primitive.Null
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeNull))

	return nil
}

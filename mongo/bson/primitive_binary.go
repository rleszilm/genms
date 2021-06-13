package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Binary exports the driver specific Binary so it can be used externally.
type Binary primitive.Binary

var (
	typeBinary          reflect.Type
	typePrimitiveBinary reflect.Type
)

func registerBinary(rb *bsoncodec.RegistryBuilder) {
	var instanceBinary Binary
	var instancePrimitiveBinary primitive.Binary

	typeBinary = reflect.TypeOf(instanceBinary)
	typePrimitiveBinary = reflect.TypeOf(instancePrimitiveBinary)

	rb.RegisterTypeDecoder(typeBinary, bsoncodec.ValueDecoderFunc(codecBinaryDecodeValue))
	rb.RegisterTypeEncoder(typeBinary, bsoncodec.ValueEncoderFunc(codecBinaryEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.Binary, typeBinary)
}

// codecBinaryEncodeValue is the ValueDecoderFunc for primitive.Binary.
func codecBinaryEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeBinary {
		return bsoncodec.ValueDecoderError{Name: "Mongo BinaryEncodeValue", Types: []reflect.Type{typeBinary}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveBinary)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveBinary))
}

// codecBinaryDecodeValue is the ValueDecoderFunc for primitive.Binary.
func codecBinaryDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeBinary {
		return bsoncodec.ValueDecoderError{Name: "Mongo BinaryDecodeValue", Types: []reflect.Type{typeBinary}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveBinary)
	if err != nil {
		return err
	}

	var new primitive.Binary
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeBinary))

	return nil
}

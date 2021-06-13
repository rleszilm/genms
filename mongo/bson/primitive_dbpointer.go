package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DBPointer exports the driver specific DBPointer so it can be used externally.
type DBPointer primitive.DBPointer

var (
	typeDBPointer          reflect.Type
	typePrimitiveDBPointer reflect.Type
)

func registerDBPointer(rb *bsoncodec.RegistryBuilder) {
	var instanceDBPointer DBPointer
	var instancePrimitiveDBPointer primitive.DBPointer

	typeDBPointer = reflect.TypeOf(instanceDBPointer)
	typePrimitiveDBPointer = reflect.TypeOf(instancePrimitiveDBPointer)

	rb.RegisterTypeDecoder(typeDBPointer, bsoncodec.ValueDecoderFunc(codecDBPointerDecodeValue))
	rb.RegisterTypeEncoder(typeDBPointer, bsoncodec.ValueEncoderFunc(codecDBPointerEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.DBPointer, typeDBPointer)
}

// codecDBPointerEncodeValue is the ValueDecoderFunc for primitive.DBPointer.
func codecDBPointerEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeDBPointer {
		return bsoncodec.ValueDecoderError{Name: "Mongo DBPointerEncodeValue", Types: []reflect.Type{typeDBPointer}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveDBPointer)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveDBPointer))
}

// codecDBPointerDecodeValue is the ValueDecoderFunc for primitive.DBPointer.
func codecDBPointerDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeDBPointer {
		return bsoncodec.ValueDecoderError{Name: "Mongo DBPointerDecodeValue", Types: []reflect.Type{typeDBPointer}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveDBPointer)
	if err != nil {
		return err
	}

	var new primitive.DBPointer
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeDBPointer))

	return nil
}

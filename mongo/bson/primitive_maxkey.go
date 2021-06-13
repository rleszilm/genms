package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MaxKey exports the driver specific MaxKey so it can be used externally.
type MaxKey primitive.MaxKey

var (
	typeMaxKey          reflect.Type
	typePrimitiveMaxKey reflect.Type
)

func registerMaxKey(rb *bsoncodec.RegistryBuilder) {
	var instanceMaxKey MaxKey
	var instancePrimitiveMaxKey primitive.MaxKey

	typeMaxKey = reflect.TypeOf(instanceMaxKey)
	typePrimitiveMaxKey = reflect.TypeOf(instancePrimitiveMaxKey)

	rb.RegisterTypeDecoder(typeMaxKey, bsoncodec.ValueDecoderFunc(codecMaxKeyDecodeValue))
	rb.RegisterTypeEncoder(typeMaxKey, bsoncodec.ValueEncoderFunc(codecMaxKeyEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.MaxKey, typeMaxKey)
}

// codecMaxKeyEncodeValue is the ValueDecoderFunc for primitive.MaxKey.
func codecMaxKeyEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeMaxKey {
		return bsoncodec.ValueDecoderError{Name: "Mongo MaxKeyEncodeValue", Types: []reflect.Type{typeMaxKey}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveMaxKey)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveMaxKey))
}

// codecMaxKeyDecodeValue is the ValueDecoderFunc for primitive.MaxKey.
func codecMaxKeyDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeMaxKey {
		return bsoncodec.ValueDecoderError{Name: "Mongo MaxKeyDecodeValue", Types: []reflect.Type{typeMaxKey}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveMaxKey)
	if err != nil {
		return err
	}

	var new primitive.MaxKey
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeMaxKey))

	return nil
}

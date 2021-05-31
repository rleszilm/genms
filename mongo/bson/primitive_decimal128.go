
package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Decimal128 exports the driver specific Decimal128 so it can be used externally.
type Decimal128 primitive.Decimal128

var (
	typeDecimal128 reflect.Type
	typePrimitiveDecimal128 reflect.Type
)

func registerDecimal128(rb *bsoncodec.RegistryBuilder) {
	var instanceDecimal128 Decimal128
	var instancePrimitiveDecimal128 primitive.Decimal128

	typeDecimal128 = reflect.TypeOf(instanceDecimal128)
	typePrimitiveDecimal128 = reflect.TypeOf(instancePrimitiveDecimal128)

	rb.RegisterTypeDecoder(typeDecimal128, bsoncodec.ValueDecoderFunc(codecDecimal128DecodeValue))
	rb.RegisterTypeEncoder(typeDecimal128, bsoncodec.ValueEncoderFunc(codecDecimal128EncodeValue))
	rb.RegisterTypeMapEntry(bsontype.Decimal128, typeDecimal128)
}

// codecDecimal128EncodeValue is the ValueDecoderFunc for primitive.Decimal128.
func codecDecimal128EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeDecimal128 {
		return bsoncodec.ValueDecoderError{Name: "Mongo Decimal128EncodeValue", Types: []reflect.Type{ typeDecimal128 }, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveDecimal128)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveDecimal128))
}

// codecDecimal128DecodeValue is the ValueDecoderFunc for primitive.Decimal128.
func codecDecimal128DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeDecimal128 {
		return bsoncodec.ValueDecoderError{Name: "Mongo Decimal128DecodeValue", Types: []reflect.Type{ typeDecimal128 }, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveDecimal128)
	if err != nil {
		return err
	}

	var new primitive.Decimal128
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeDecimal128))

	return nil
}

// MarshalJSON returns the Decimal128 as a json encoded string.
func (id Decimal128) MarshalJSON() ([]byte, error) {
	return primitive.Decimal128(id).MarshalJSON()
}

// UnmarshalJSON populates the Decimal128.
func (id *Decimal128) UnmarshalJSON(b []byte) error {
	return (*primitive.Decimal128)(id).UnmarshalJSON(b)
}


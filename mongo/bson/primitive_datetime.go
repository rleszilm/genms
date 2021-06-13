
package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DateTime exports the driver specific DateTime so it can be used externally.
type DateTime primitive.DateTime

var (
	typeDateTime reflect.Type
	typePrimitiveDateTime reflect.Type
)

func registerDateTime(rb *bsoncodec.RegistryBuilder) {
	var instanceDateTime DateTime
	var instancePrimitiveDateTime primitive.DateTime

	typeDateTime = reflect.TypeOf(instanceDateTime)
	typePrimitiveDateTime = reflect.TypeOf(instancePrimitiveDateTime)

	rb.RegisterTypeDecoder(typeDateTime, bsoncodec.ValueDecoderFunc(codecDateTimeDecodeValue))
	rb.RegisterTypeEncoder(typeDateTime, bsoncodec.ValueEncoderFunc(codecDateTimeEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.DateTime, typeDateTime)
}

// codecDateTimeEncodeValue is the ValueDecoderFunc for primitive.DateTime.
func codecDateTimeEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeDateTime {
		return bsoncodec.ValueDecoderError{Name: "Mongo DateTimeEncodeValue", Types: []reflect.Type{ typeDateTime }, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveDateTime)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveDateTime))
}

// codecDateTimeDecodeValue is the ValueDecoderFunc for primitive.DateTime.
func codecDateTimeDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeDateTime {
		return bsoncodec.ValueDecoderError{Name: "Mongo DateTimeDecodeValue", Types: []reflect.Type{ typeDateTime }, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveDateTime)
	if err != nil {
		return err
	}

	var new primitive.DateTime
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeDateTime))

	return nil
}

// MarshalJSON returns the DateTime as a json encoded string.
func (id DateTime) MarshalJSON() ([]byte, error) {
	return primitive.DateTime(id).MarshalJSON()
}

// UnmarshalJSON populates the DateTime.
func (id *DateTime) UnmarshalJSON(b []byte) error {
	return (*primitive.DateTime)(id).UnmarshalJSON(b)
}


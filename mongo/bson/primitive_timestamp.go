package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Timestamp exports the driver specific Timestamp so it can be used externally.
type Timestamp primitive.Timestamp

var (
	typeTimestamp          reflect.Type
	typePrimitiveTimestamp reflect.Type
)

func registerTimestamp(rb *bsoncodec.RegistryBuilder) {
	var instanceTimestamp Timestamp
	var instancePrimitiveTimestamp primitive.Timestamp

	typeTimestamp = reflect.TypeOf(instanceTimestamp)
	typePrimitiveTimestamp = reflect.TypeOf(instancePrimitiveTimestamp)

	rb.RegisterTypeDecoder(typeTimestamp, bsoncodec.ValueDecoderFunc(codecTimestampDecodeValue))
	rb.RegisterTypeEncoder(typeTimestamp, bsoncodec.ValueEncoderFunc(codecTimestampEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.Timestamp, typeTimestamp)
}

// codecTimestampEncodeValue is the ValueDecoderFunc for primitive.Timestamp.
func codecTimestampEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeTimestamp {
		return bsoncodec.ValueDecoderError{Name: "Mongo TimestampEncodeValue", Types: []reflect.Type{typeTimestamp}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveTimestamp)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveTimestamp))
}

// codecTimestampDecodeValue is the ValueDecoderFunc for primitive.Timestamp.
func codecTimestampDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeTimestamp {
		return bsoncodec.ValueDecoderError{Name: "Mongo TimestampDecodeValue", Types: []reflect.Type{typeTimestamp}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveTimestamp)
	if err != nil {
		return err
	}

	var new primitive.Timestamp
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeTimestamp))

	return nil
}

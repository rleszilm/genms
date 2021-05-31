package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Regex exports the driver specific Regex so it can be used externally.
type Regex primitive.Regex

var (
	typeRegex          reflect.Type
	typePrimitiveRegex reflect.Type
)

func registerRegex(rb *bsoncodec.RegistryBuilder) {
	var instanceRegex Regex
	var instancePrimitiveRegex primitive.Regex

	typeRegex = reflect.TypeOf(instanceRegex)
	typePrimitiveRegex = reflect.TypeOf(instancePrimitiveRegex)

	rb.RegisterTypeDecoder(typeRegex, bsoncodec.ValueDecoderFunc(codecRegexDecodeValue))
	rb.RegisterTypeEncoder(typeRegex, bsoncodec.ValueEncoderFunc(codecRegexEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.Regex, typeRegex)
}

// codecRegexEncodeValue is the ValueDecoderFunc for primitive.Regex.
func codecRegexEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeRegex {
		return bsoncodec.ValueDecoderError{Name: "Mongo RegexEncodeValue", Types: []reflect.Type{typeRegex}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveRegex)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveRegex))
}

// codecRegexDecodeValue is the ValueDecoderFunc for primitive.Regex.
func codecRegexDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeRegex {
		return bsoncodec.ValueDecoderError{Name: "Mongo RegexDecodeValue", Types: []reflect.Type{typeRegex}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveRegex)
	if err != nil {
		return err
	}

	var new primitive.Regex
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeRegex))

	return nil
}

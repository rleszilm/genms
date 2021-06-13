package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CodeWithScope exports the driver specific CodeWithScope so it can be used externally.
type CodeWithScope primitive.CodeWithScope

var (
	typeCodeWithScope          reflect.Type
	typePrimitiveCodeWithScope reflect.Type
)

func registerCodeWithScope(rb *bsoncodec.RegistryBuilder) {
	var instanceCodeWithScope CodeWithScope
	var instancePrimitiveCodeWithScope primitive.CodeWithScope

	typeCodeWithScope = reflect.TypeOf(instanceCodeWithScope)
	typePrimitiveCodeWithScope = reflect.TypeOf(instancePrimitiveCodeWithScope)

	rb.RegisterTypeDecoder(typeCodeWithScope, bsoncodec.ValueDecoderFunc(codecCodeWithScopeDecodeValue))
	rb.RegisterTypeEncoder(typeCodeWithScope, bsoncodec.ValueEncoderFunc(codecCodeWithScopeEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.CodeWithScope, typeCodeWithScope)
}

// codecCodeWithScopeEncodeValue is the ValueDecoderFunc for primitive.CodeWithScope.
func codecCodeWithScopeEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeCodeWithScope {
		return bsoncodec.ValueDecoderError{Name: "Mongo CodeWithScopeEncodeValue", Types: []reflect.Type{typeCodeWithScope}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveCodeWithScope)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveCodeWithScope))
}

// codecCodeWithScopeDecodeValue is the ValueDecoderFunc for primitive.CodeWithScope.
func codecCodeWithScopeDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeCodeWithScope {
		return bsoncodec.ValueDecoderError{Name: "Mongo CodeWithScopeDecodeValue", Types: []reflect.Type{typeCodeWithScope}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveCodeWithScope)
	if err != nil {
		return err
	}

	var new primitive.CodeWithScope
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeCodeWithScope))

	return nil
}

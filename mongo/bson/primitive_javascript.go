package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JavaScript exports the driver specific JavaScript so it can be used externally.
type JavaScript primitive.JavaScript

var (
	typeJavaScript          reflect.Type
	typePrimitiveJavaScript reflect.Type
)

func registerJavaScript(rb *bsoncodec.RegistryBuilder) {
	var instanceJavaScript JavaScript
	var instancePrimitiveJavaScript primitive.JavaScript

	typeJavaScript = reflect.TypeOf(instanceJavaScript)
	typePrimitiveJavaScript = reflect.TypeOf(instancePrimitiveJavaScript)

	rb.RegisterTypeDecoder(typeJavaScript, bsoncodec.ValueDecoderFunc(codecJavaScriptDecodeValue))
	rb.RegisterTypeEncoder(typeJavaScript, bsoncodec.ValueEncoderFunc(codecJavaScriptEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.JavaScript, typeJavaScript)
}

// codecJavaScriptEncodeValue is the ValueDecoderFunc for primitive.JavaScript.
func codecJavaScriptEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeJavaScript {
		return bsoncodec.ValueDecoderError{Name: "Mongo JavaScriptEncodeValue", Types: []reflect.Type{typeJavaScript}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveJavaScript)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveJavaScript))
}

// codecJavaScriptDecodeValue is the ValueDecoderFunc for primitive.JavaScript.
func codecJavaScriptDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeJavaScript {
		return bsoncodec.ValueDecoderError{Name: "Mongo JavaScriptDecodeValue", Types: []reflect.Type{typeJavaScript}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveJavaScript)
	if err != nil {
		return err
	}

	var new primitive.JavaScript
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeJavaScript))

	return nil
}

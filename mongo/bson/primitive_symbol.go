package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Symbol exports the driver specific Symbol so it can be used externally.
type Symbol primitive.Symbol

var (
	typeSymbol          reflect.Type
	typePrimitiveSymbol reflect.Type
)

func registerSymbol(rb *bsoncodec.RegistryBuilder) {
	var instanceSymbol Symbol
	var instancePrimitiveSymbol primitive.Symbol

	typeSymbol = reflect.TypeOf(instanceSymbol)
	typePrimitiveSymbol = reflect.TypeOf(instancePrimitiveSymbol)

	rb.RegisterTypeDecoder(typeSymbol, bsoncodec.ValueDecoderFunc(codecSymbolDecodeValue))
	rb.RegisterTypeEncoder(typeSymbol, bsoncodec.ValueEncoderFunc(codecSymbolEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.Symbol, typeSymbol)
}

// codecSymbolEncodeValue is the ValueDecoderFunc for primitive.Symbol.
func codecSymbolEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeSymbol {
		return bsoncodec.ValueDecoderError{Name: "Mongo SymbolEncodeValue", Types: []reflect.Type{typeSymbol}, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveSymbol)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveSymbol))
}

// codecSymbolDecodeValue is the ValueDecoderFunc for primitive.Symbol.
func codecSymbolDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeSymbol {
		return bsoncodec.ValueDecoderError{Name: "Mongo SymbolDecodeValue", Types: []reflect.Type{typeSymbol}, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveSymbol)
	if err != nil {
		return err
	}

	var new primitive.Symbol
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeSymbol))

	return nil
}

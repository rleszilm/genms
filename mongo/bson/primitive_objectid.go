
package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectID exports the driver specific ObjectID so it can be used externally.
type ObjectID primitive.ObjectID

var (
	typeObjectID reflect.Type
	typePrimitiveObjectID reflect.Type
)

func registerObjectID(rb *bsoncodec.RegistryBuilder) {
	var instanceObjectID ObjectID
	var instancePrimitiveObjectID primitive.ObjectID

	typeObjectID = reflect.TypeOf(instanceObjectID)
	typePrimitiveObjectID = reflect.TypeOf(instancePrimitiveObjectID)

	rb.RegisterTypeDecoder(typeObjectID, bsoncodec.ValueDecoderFunc(codecObjectIDDecodeValue))
	rb.RegisterTypeEncoder(typeObjectID, bsoncodec.ValueEncoderFunc(codecObjectIDEncodeValue))
	rb.RegisterTypeMapEntry(bsontype.ObjectID, typeObjectID)
}

// codecObjectIDEncodeValue is the ValueDecoderFunc for primitive.ObjectID.
func codecObjectIDEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != typeObjectID {
		return bsoncodec.ValueDecoderError{Name: "Mongo ObjectIDEncodeValue", Types: []reflect.Type{ typeObjectID }, Received: val}
	}

	enc, err := ec.LookupEncoder(typePrimitiveObjectID)
	if err != nil {
		return err
	}

	return enc.EncodeValue(ec, vw, val.Convert(typePrimitiveObjectID))
}

// codecObjectIDDecodeValue is the ValueDecoderFunc for primitive.ObjectID.
func codecObjectIDDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != typeObjectID {
		return bsoncodec.ValueDecoderError{Name: "Mongo ObjectIDDecodeValue", Types: []reflect.Type{ typeObjectID }, Received: val}
	}

	dec, err := dc.LookupDecoder(typePrimitiveObjectID)
	if err != nil {
		return err
	}

	var new primitive.ObjectID
	v := reflect.ValueOf(&new)
	if err := dec.DecodeValue(dc, vr, v.Elem()); err != nil {
		return err
	}
	val.Set(v.Elem().Convert(typeObjectID))

	return nil
}

// MarshalJSON returns the ObjectID as a json encoded string.
func (id ObjectID) MarshalJSON() ([]byte, error) {
	return primitive.ObjectID(id).MarshalJSON()
}

// UnmarshalJSON populates the ObjectID.
func (id *ObjectID) UnmarshalJSON(b []byte) error {
	return (*primitive.ObjectID)(id).UnmarshalJSON(b)
}


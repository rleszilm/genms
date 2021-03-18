package generator

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// Field adds functionality to the underlying field.
type Field struct {
	*protocgenlib.Field
	dalOptions  *annotations.DalFieldOptions
	typeOptions *annotations.DalFieldOptions_BackendFieldOptions
}

// NewField returns a new Field.
func NewField(msg *Message, field *protogen.Field) *Field {
	return AsField(protocgenlib.NewField(msg.Message, field))
}

// AsField wraps a Field.
func AsField(f *protocgenlib.Field) *Field {
	field := &Field{
		Field: f,
	}

	opts := f.Proto().Desc.Options()
	if proto.HasExtension(f.Proto().Desc.Options(), annotations.E_FieldOptions) {
		dalOptions := proto.GetExtension(opts, annotations.E_FieldOptions).(*annotations.DalFieldOptions)
		field.dalOptions = dalOptions
		field.typeOptions = dalOptions.GetRest()
	}

	return field
}

// QueryName returns the name of the field as it should appear in database queries.
func (f *Field) QueryName() string {
	if f.typeOptions != nil {
		if name := f.typeOptions.GetField(); name != "" {
			return name
		}
	}

	if f.dalOptions != nil {
		if name := f.dalOptions.GetField(); name != "" {
			return name
		}
	}

	return f.Name()
}

// Ignore returns whether the field should be skipped during processing.
func (f *Field) Ignore() bool {
	if f.dalOptions != nil {
		return f.dalOptions.Ignore
	}
	return false
}

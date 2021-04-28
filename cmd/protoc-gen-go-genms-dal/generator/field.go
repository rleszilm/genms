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
	options *annotations.DalFieldOptions
}

// NewField returns a new Field.
func NewField(msg *Message, field *protogen.Field) *Field {
	return AsField(protocgenlib.NewField(msg.Message, field))
}

// AsField wraps a Field.
func AsField(f *protocgenlib.Field) *Field {
	opts := f.Proto().Desc.Options()
	var dalOptions *annotations.DalFieldOptions
	if proto.HasExtension(opts, annotations.E_FieldOptions) {
		if topts, ok := proto.GetExtension(opts, annotations.E_FieldOptions).(*annotations.DalFieldOptions); ok {
			dalOptions = topts
		}
	}

	field := &Field{
		Field:   f,
		options: dalOptions,
	}

	return field
}

// Generator returns the Generator level Field.
func (f *Field) Generator() *Field {
	return f
}

// Options returns the DalFieldOptions
func (f *Field) Options() *annotations.DalFieldOptions {
	return f.options
}

// QueryName returns the name of the field as it should appear in database queries.
func (f *Field) QueryName() string {
	if f.options != nil {
		if name := f.options.GetField(); name != "" {
			return name
		}
	}

	return f.Name()
}

// Ignore returns whether the field should be skipped during processing.
func (f *Field) Ignore() bool {
	if f.options != nil {
		return f.options.GetIgnore()
	}
	return false
}

package rest

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// Field adds functionality to the underlying field.
type Field struct {
	*generator.Field
}

// NewField returns a new Field.
func NewField(msg *Message, field *protogen.Field) *Field {
	return AsField(protocgenlib.NewField(msg.ProtocGenLib(), field))
}

// AsField wraps a Field.
func AsField(f *protocgenlib.Field) *Field {
	field := &Field{
		Field: generator.AsField(f),
	}

	return field
}

// QueryName returns the name of the field as it should appear in database queries.
func (f *Field) QueryName() string {
	opts := f.Options()
	if opts != nil && opts.GetRest() != nil {
		if name := opts.GetRest().GetField(); name != "" {
			return name
		}
	}

	return f.Generator().QueryName()
}

// Ignore returns the name of the field as it should appear in database queries.
func (f *Field) Ignore() bool {
	opts := f.Options()
	if opts != nil && opts.GetRest() != nil {
		return opts.GetRest().GetIgnore()
	}

	return f.Generator().Ignore()
}

package protocgenlib

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Field adds functionality to the underlying Field.
type Field struct {
	Outfile *protogen.GeneratedFile
	Message *protogen.Message
	Field   *protogen.Field
}

// NewField returns a new Field.
func NewField(outfile *protogen.GeneratedFile, msg *protogen.Message, field *protogen.Field) *Field {
	return &Field{
		Outfile: outfile,
		Message: msg,
		Field:   field,
	}
}

// Name returns the name of the field.
func (f *Field) Name() string {
	return string(f.Field.Desc.Name())
}

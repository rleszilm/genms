package protocgenlib

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Field adds functionality to the underlying Field.
type Field struct {
	message *Message
	field   *protogen.Field
}

// NewField returns a new Field.
func NewField(msg *Message, field *protogen.Field) *Field {
	return &Field{
		message: msg,
		field:   field,
	}
}

// Proto returns the base protogen object.
func (f *Field) Proto() *protogen.Field {
	return f.field
}

// ProtocGenLib returns the base protogen object.
func (f *Field) ProtocGenLib() *Field {
	return f
}

// Message returns the underlying Message.
func (f *Field) Message() *Message {
	return f.message
}

// Outfile returns the file to which this field would be written.
func (f *Field) Outfile() *protogen.GeneratedFile {
	return f.message.Outfile()
}

// Name returns the name of the field.
func (f *Field) Name() string {
	return string(f.field.Desc.Name())
}

// Kind returns the fields go type.
func (f *Field) Kind() string {
	prefix := f.Prefix()
	if f.field.Message != nil {
		return prefix + f.field.Message.GoIdent.GoName
	}
	if f.field.Enum != nil {
		return prefix + f.field.Enum.GoIdent.GoName
	}
	return ToGoKind(f.field.Desc.Kind(), prefix)
}

// QualifiedKind returns the fully qualified kind.
func (f *Field) QualifiedKind() string {
	prefix := f.Prefix()
	if f.field.Message != nil {
		return prefix + f.message.File().QualifiedKind(f.field.Message.GoIdent)
	}
	if f.field.Enum != nil {
		return prefix + f.message.File().QualifiedKind(f.field.Enum.GoIdent)
	}
	return ToGoKind(f.field.Desc.Kind(), prefix)
}

// Prefix returns the prefix.
func (f *Field) Prefix() string {
	var prefix string
	slice := f.IsSlice()
	message := f.IsMessage()
	optional := f.IsOptional()
	if slice && message {
		prefix = "[]*"
	} else if slice {
		prefix = "[]"
	} else if optional || message {
		prefix = "*"
	}
	return prefix
}

// ToRef returns the string needed to make a reference of the field.
func (f *Field) ToRef() string {
	if f.IsMessage() || f.IsOptional() {
		return "&"
	}
	return ""
}

// IsRef returns whether the internal type is a reference.
func (f *Field) IsRef() bool {
	return f.IsMessage() || f.IsOptional()
}

// IsMessage returns whether the field is a reference.
func (f *Field) IsMessage() bool {
	return f.field.Message != nil
}

// IsSlice returns whether the field is a slice of values.
func (f *Field) IsSlice() bool {
	return f.field.Desc.IsList()
}

// IsOptional returns whether the field is marked as optional.
func (f *Field) IsOptional() bool {
	return f.field.Desc.HasOptionalKeyword()
}

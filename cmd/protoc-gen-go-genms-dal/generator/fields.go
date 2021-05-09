package generator

import (
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
)

// Fields is a struct that contains data about the messages fields.
type Fields struct {
	*protocgenlib.Fields
}

// NewFields returns a new Fields
func NewFields(msg *Message) *Fields {
	return AsFields(protocgenlib.NewFields(msg.Message))
}

// AsFields wraps Fields.
func AsFields(fields *protocgenlib.Fields) *Fields {
	return &Fields{
		Fields: fields,
	}
}

// Generator returns the generator level Fields
func (f *Fields) Generator() *Fields {
	return f
}

// ByName returns the specified field.
func (f *Fields) ByName(n string) *Field {
	field := f.Fields.ByName(n)
	if field == nil {
		return nil
	}

	return AsField(field.ProtocGenLib())
}

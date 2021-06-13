package mongo

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
)

// Fields is a struct that contains data about the messages fields.
type Fields struct {
	id *Field
	*generator.Fields
}

// NewFields returns a new Fields
func NewFields(msg *Message) *Fields {
	return AsFields(protocgenlib.NewFields(msg.ProtocGenLib()))
}

// AsFields wraps Fields.
func AsFields(fields *protocgenlib.Fields) *Fields {
	return &Fields{
		Fields: generator.AsFields(fields),
	}
}

// ByName returns the specified field.
func (f *Fields) ByName(n string) *Field {
	field := f.Fields.ByName(n)
	if field == nil {
		return nil
	}

	return AsField(field.ProtocGenLib())
}

// ID returns the field that contains the _id of hte collection.
func (f *Fields) ID() *Field {
	if f.id != nil {
		return f.id
	}

	for _, n := range f.Names() {
		field := f.ByName(n)
		if field.QueryName() == "_id" {
			f.id = field
		}
	}

	return f.id
}

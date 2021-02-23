package generator

import (
	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/gen_microservice/internal/protoc-gen-lib"
	"google.golang.org/protobuf/proto"
)

// Field adds functionality to the underlying field.
type Field struct {
	*protocgenlib.Field
	dalOptions  *annotations.DalFieldOptions
	typeOptions *annotations.DalFieldOptions_BackendFieldOptions
}

// AsRestField returns a new Field.
func AsRestField(f *protocgenlib.Field) *Field {
	field := &Field{
		Field: f,
	}

	opts := f.Field.Desc.Options()
	if proto.HasExtension(f.Field.Desc.Options(), annotations.E_FieldOptions) {
		dalOptions := proto.GetExtension(opts, annotations.E_FieldOptions).(*annotations.DalFieldOptions)
		field.dalOptions = dalOptions
		field.typeOptions = dalOptions.GetRest()
	}

	return field
}

// QueryName returns the name of the field as it should appear in database queries.
func (f *Field) QueryName() string {
	if name := f.typeOptions.GetField(); name != "" {
		return name
	}

	if name := f.dalOptions.GetField(); name != "" {
		return name
	}

	return f.Name()
}

// Ignore returns whether the field should be skipped during processing.
func (f *Field) Ignore() bool {
	return f.dalOptions.Ignore
}

package postgres

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// Field adds functionality to the underlying field.
type Field struct {
	*generator.Field
	dalOptions  *annotations.DalFieldOptions
	typeOptions *annotations.DalFieldOptions_BackendFieldOptions
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

// SQLNilKind returns the SQL nil kind for this field.
func (f *Field) SQLNilKind() string {
	pkg := f.Message().File().QualifiedPackageName("database/sql")

	switch f.QualifiedKind() {
	case "bool":
		return pkg + ".NullBool"
	case "float64":
		return pkg + ".NullFloat64"
	case "float32":
		return pkg + ".NullFloat64"
	case "int32":
		return pkg + ".NullInt32"
	case "int64":
		return pkg + ".NullInt64"
	case "string":
		return pkg + ".NullString"
	default:
		if f.Proto().Enum != nil {
			return pkg + ".NullInt32"
		}
		return f.QualifiedKind()
	}
}

// ValueFromSQLNil returns the go value from the given nil value.
func (f *Field) ValueFromSQLNil() string {
	pkg := f.Message().File().QualifiedPackageName("database/sql")

	switch f.QualifiedKind() {
	case "bool":
		return pkg + ".NullBool"
	case "float64":
		return pkg + ".NullFloat64"
	case "float32":
		return pkg + ".NullFloat64"
	case "int32":
		return pkg + ".NullInt32"
	case "int64":
		return pkg + ".NullInt64"
	case "string":
		return pkg + ".NullString"
	default:
		if f.Proto().Enum != nil {
			return pkg + ".NullInt32"
		}
		return f.QualifiedKind()
	}
}

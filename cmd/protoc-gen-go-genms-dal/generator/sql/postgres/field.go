package postgres

import (
	"fmt"

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
	if opts != nil && opts.GetPostgres() != nil {
		if name := opts.GetPostgres().GetField(); name != "" {
			return name
		}
	}

	return f.Generator().QueryName()
}

// Ignore returns the name of the field as it should appear in database queries.
func (f *Field) Ignore() bool {
	opts := f.Options()
	if opts != nil && opts.GetPostgres() != nil {
		return opts.GetPostgres().GetIgnore()
	}

	return f.Generator().Ignore()
}

// HasSQLNil returns whether the sql scanner would use a nil type.
func (f *Field) HasSQLNil() bool {
	if f.Proto().Message != nil {
		return false
	}
	return true
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

// ValueFromSQLValue returns the go logic to convert the scanned value to the go value.
func (f *Field) ValueFromSQLValue(obj string) string {
	switch f.QualifiedKind() {
	case "bool":
		return fmt.Sprintf("%s.%s.Bool", obj, protocgenlib.ToTitleCase(f.Name()))
	case "float32":
		return fmt.Sprintf("float32(%s.%s.Float64)", obj, protocgenlib.ToTitleCase(f.Name()))
	case "float64":
		return fmt.Sprintf("%s.%s.Float64", obj, protocgenlib.ToTitleCase(f.Name()))
	case "int32":
		return fmt.Sprintf("%s.%s.Int32", obj, protocgenlib.ToTitleCase(f.Name()))
	case "int64":
		return fmt.Sprintf("%s.%s.Int64", obj, protocgenlib.ToTitleCase(f.Name()))
	case "string":
		return fmt.Sprintf("%s.%s.String", obj, protocgenlib.ToTitleCase(f.Name()))
	default:
		if f.Proto().Enum != nil {
			return fmt.Sprintf("%s(%s.%s.Int32)", f.QualifiedKind(), obj, protocgenlib.ToTitleCase(f.Name()))
		}
		return fmt.Sprintf("%s.%s", obj, protocgenlib.ToTitleCase(f.Name()))
	}
}

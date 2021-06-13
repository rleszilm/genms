package mongo

import (
	"fmt"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
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
	if opts != nil && opts.GetMongo() != nil {
		if name := opts.GetMongo().GetName(); name != "" {
			return name
		}
	}

	return f.Generator().QueryName()
}

// QueryKind returns the fields go type.
func (f *Field) QueryKind() (string, error) {
	if f.Options().GetMongo().GetBson() == annotations.BSONPrimitive_NoBSONPrimitive {
		return f.Kind(), nil
	}

	switch f.Options().GetMongo().GetBson() {
	case annotations.BSONPrimitive_NoBSONPrimitive:
		return f.Kind(), nil
	case annotations.BSONPrimitive_ObjectID:
		return "ObjectID", nil
	default:
		return "", fmt.Errorf("invalid bson primitive: %v", f.Options().GetMongo().GetBson())
	}
}

// QualifiedQueryKind returns the fully qualified kind.
func (f *Field) QualifiedQueryKind() (string, error) {
	if f.Options().GetMongo().GetBson() == annotations.BSONPrimitive_NoBSONPrimitive {
		return f.QualifiedKind(), nil
	}

	switch f.Options().GetMongo().GetBson() {
	case annotations.BSONPrimitive_NoBSONPrimitive:
		return f.Kind(), nil
	case annotations.BSONPrimitive_ObjectID:
		return f.Outfile().QualifiedGoIdent(protogen.GoIdent{GoName: "ObjectID", GoImportPath: "github.com/rleszilm/genms/mongo/bson"}), nil
	default:
		return "", fmt.Errorf("invalid bson primitive: %v", f.Options().GetMongo().GetBson())
	}
}

// Ignore returns the name of the field as it should appear in database queries.
func (f *Field) Ignore() bool {
	opts := f.Options()
	if opts != nil && opts.GetMongo() != nil {
		return opts.GetMongo().GetIgnore()
	}

	return f.Generator().Ignore()
}

// ToMongo indicates what type the field value should be converted to.
func (f *Field) ToMongo() string {
	kind := f.Options().GetMongo().GetBson()
	if kind == annotations.BSONPrimitive_NoBSONPrimitive {
		return ""
	}
	return annotations.BSONPrimitive_name[int32(kind)]
}

// ToGo indicates what type the mongo field should be converted to.
func (f *Field) ToGo() (string, error) {
	kind := f.Options().GetMongo().GetBson()
	switch kind {
	case annotations.BSONPrimitive_NoBSONPrimitive:
		return "", nil
	}
	return f.Kind(), nil
}

// IsExtRef returns whether the external type is a reference.
func (f *Field) IsExtRef() bool {
	return f.IsExtMessage() || f.IsExtSlice()
}

// IsExtMessage returns whether the field is a reference.
func (f *Field) IsExtMessage() bool {
	return f.IsMessage()
}

// IsExtSlice returns whether the field is a slice of values.
func (f *Field) IsExtSlice() bool {
	qKind, err := f.QualifiedQueryKind()
	if err != nil {
		return false
	}

	if len(qKind) < 2 {
		return false
	}

	return qKind[:2] == "[]"
}

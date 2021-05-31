package generator

import (
	"errors"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// Query adds functionality to the query options.
type Query struct {
	*annotations.Query
	File   *File
	Fields *Fields
}

// NewQuery returns a new Query
func NewQuery(file *File, fields *Fields, q *annotations.Query) *Query {
	return &Query{
		Query:  q,
		File:   file,
		Fields: fields,
	}
}

// Method returns the interface definition,
func (q *Query) Method() string {
	return protocgenlib.ToTitleCase(q.Name)
}

// QueryProvided returns whether a query should be formatted and stored.
func (q *Query) QueryProvided() bool {
	switch q.Mode {
	case annotations.Query_Auto, annotations.Query_ProviderStub:
		return true
	default:
		return false
	}
}

// QueryImplemented returns whether a query should be formatted and stored.
func (q *Query) QueryImplemented() bool {
	switch q.Mode {
	case annotations.Query_Auto:
		return true
	default:
		return false
	}
}

// ArgKind returns the kind of the specified arg.
func (q *Query) ArgKind(a *annotations.Arg) (string, error) {
	if a.GetName() != "" {
		f := q.Fields.ByName(a.GetName())
		if f == nil {
			return "", errors.New("no field")
		}
		return f.QualifiedKind(), nil
	}
	return q.File.QualifiedKind(protogen.GoIdent{GoName: a.GetKind()}), nil
}

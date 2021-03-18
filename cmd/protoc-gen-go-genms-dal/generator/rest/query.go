package rest

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// Queries is a set of queries.
type Queries []*Query

// NewQueries generates a set of queries from a message.
func NewQueries(msg *protogen.Message, fields *protocgenlib.Fields) Queries {
	queries := Queries{}

	//fOpts := f.Desc.Options()
	//	if proto.HasExtension(fOpts, annotations.E_FieldOptions) {
	//		ext := proto.GetExtension(fOpts, annotations.E_FieldOptions).(*annotations.DalFieldOptions)

	mOpts := msg.Desc.Options()
	if proto.HasExtension(mOpts, annotations.E_MessageOptions) {
		dOpts, ok := proto.GetExtension(mOpts, annotations.E_MessageOptions).(*annotations.DalOptions)
		if !ok {
			return queries
		}

		for _, q := range dOpts.Queries {
			queries = append(queries, NewQuery(msg, fields, q))
		}
	}

	return queries
}

// Query adds functionality to the query options.
type Query struct {
	*generator.Query
	opts *annotations.DalOptions_Query
}

// NewQuery returns a new Query
func NewQuery(msg *protogen.Message, fields *protocgenlib.Fields, query *annotations.DalOptions_Query) *Query {
	q := &Query{
		//	Query: generator.NewQuery(msg, fields, query),
		opts: query,
	}

	return q
}

// Implementation returns the queries implementation.
func (q *Query) Implementation() string {
	return ""
}

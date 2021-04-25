package rest

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
)

// Query adds functionality to the query options.
type Query struct {
	*generator.Query
}

// NewQuery returns a new Query
func NewQuery(file *File, fields *Fields, q *annotations.DalOptions_Query) *Query {
	return AsQuery(generator.NewQuery(file.Generator(), fields.Generator(), q))
}

// AsQuery returns the a query.
func AsQuery(q *generator.Query) *Query {
	return &Query{Query: q}
}

// Method returns the method of the rest query.
func (q *Query) Method() string {
	return q.GetRest().GetMethod().String()
}

package generator

import "github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator"

// Query adds functionality to the query options.
type Query struct {
	*generator.Query
}

// NewQuery returns a new Query
func NewQuery() *Query {
	return &Query{
		Query: generator.NewQuery(),
	}
}

// Implementation returns the queries implementation.
func (q *Query) Implementation() string {
	return ""
}

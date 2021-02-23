package generator

import protocgenlib "github.com/rleszilm/gen_microservice/internal/protoc-gen-lib"

// Query adds functionality to the query options.
type Query struct {
	message *Message
	fields  *protocgenlib.Fields
}

// NewQuery returns a new Query
func NewQuery() *Query {
	return &Query{}
}

// Method returns the interface definition,
func (q *Query) Method() string {
	return ""
}

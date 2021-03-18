package generator

import (
	"fmt"
	"strings"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
)

// Query adds functionality to the query options.
type Query struct {
	Fields *Fields
	Query  *annotations.DalOptions_Query
}

// NewQuery returns a new Query
func NewQuery(fields *Fields, query *annotations.DalOptions_Query) *Query {
	return &Query{
		Fields: fields,
		Query:  query,
	}
}

// Method returns the interface definition,
func (q *Query) Method() string {
	return protocgenlib.ToTitleCase(q.Query.Name)
}

// Args returns the method arguments.
func (q *Query) Args() string {
	tokens := []string{}

	for _, a := range q.Query.Args {
		field := q.Fields.ByName(a)
		tokens = append(tokens, fmt.Sprintf("%s %s", field.Name(), field.Kind()))
	}

	return strings.Join([]string{}, ", ")
}

package postgres

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
)

// Query adds functionality to the query options.
type Query struct {
	*annotations.DalOptions_Query
}

// NewQuery returns a new Query
func NewQuery(q *annotations.DalOptions_Query) *Query {
	return &Query{
		DalOptions_Query: q,
	}
}

// Method returns the interface definition,
func (q *Query) Method() string {
	return protocgenlib.ToTitleCase(q.Name)
}

// QueryProvided returns whether a query should be formatted and stored.
func (q *Query) QueryProvided() bool {
	switch q.Mode {
	case annotations.DalOptions_Query_QueryMode_Auto, annotations.DalOptions_Query_QueryMode_ProviderStub:
		return true
	default:
		return false
	}
}

// QueryImplemented returns whether a query should be formatted and stored.
func (q *Query) QueryImplemented() bool {
	switch q.Mode {
	case annotations.DalOptions_Query_QueryMode_Auto:
		return true
	default:
		return false
	}
}

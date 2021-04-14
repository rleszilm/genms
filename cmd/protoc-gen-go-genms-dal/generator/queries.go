package generator

import "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"

// Queries is a struct that contains data about the messages queries.
type Queries struct {
	queryNames    []string
	queriesByName map[string]*Query
}

// Queries returns the query names.
func (q *Queries) Names() []string {
	return q.queryNames
}

// ByName returns the specified query.
func (q *Queries) ByName(n string) *Query {
	return q.queriesByName[n]
}

// NewQueries returns a new queries.
func NewQueries(file *File, opts *annotations.DalOptions) *Queries {
	queryNames := []string{}
	queriesByName := map[string]*Query{}

	for _, q := range opts.Queries {
		query := NewQuery(file, q)
		queryNames = append(queryNames, query.Name)
		queriesByName[query.Name] = query
	}

	return &Queries{
		queryNames:    queryNames,
		queriesByName: queriesByName,
	}
}

package postgres

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
)

// Queries is a struct that contains data about the messages queries.
type Queries struct {
	*generator.Queries
}

// AsQueries wraps a Queries.
func NewQueries(file *generator.File, opts *annotations.DalOptions) *Queries {
	return &Queries{
		Queries: generator.NewQueries(file, opts),
	}
}

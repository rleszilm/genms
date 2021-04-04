package sql

import (
	"context"

	"github.com/rleszilm/genms/service"
)

// DB defines the interface to use when getting database connections.
type DB interface {
	service.Service

	Bind(string, interface{}) (string, []interface{}, error)
	Rebind(string) string
	Query(context.Context, string, ...interface{}) (Rows, error)
	QueryWithReplacements(context.Context, string, interface{}) (Rows, error)
	Exec(context.Context, string, ...interface{}) (Result, error)
	ExecWithReplacements(context.Context, string, interface{}) (Result, error)
}

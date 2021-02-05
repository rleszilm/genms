package sqlx

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rleszilm/gen_microservice/service"
	"github.com/rleszilm/gen_microservice/sql"
)

// DB is a sql.DB that uses sqlx under the hood.
type DB struct {
	service.Deps

	db  *sqlx.DB
	cfg sql.Config
}

// Initialize implements service.Service.Initialize.
func (d *DB) Initialize(_ context.Context) error {
	db, err := sqlx.Connect(d.cfg.Driver(), d.cfg.ConnectionString())
	if err != nil {
		return err
	}
	d.db = db

	return nil
}

// Shutdown implements service.Service.Shutdown.
func (d *DB) Shutdown(_ context.Context) error {
	return d.db.Close()
}

// NameOf implements service.Service.NameOf.
func (d *DB) NameOf() string {
	return "sqlx"
}

// String implements service.Service.String
func (d *DB) String() string {
	return d.NameOf()
}

// Bind implements sql.DB.Bind
func (d *DB) Bind(query string, arg interface{}) (string, []interface{}, error) {
	return d.db.BindNamed(query, arg)
}

// Rebind implements sql.DB.Rebind
func (d *DB) Rebind(query string) string {
	return d.db.Rebind(query)
}

// Query implements sql.DB.Query
func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (sql.Rows, error) {
	return d.db.QueryxContext(ctx, query, args...)
}

// QueryWithReplacements implements sql.DB.QueryWithReplacements
func (d *DB) QueryWithReplacements(ctx context.Context, query string, arg interface{}) (sql.Rows, error) {
	return d.db.NamedQueryContext(ctx, query, arg)
}

// Exec implements sql.DB.Exec
func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

// ExecWithReplacements implements sql.DB.ExecWithReplacements
func (d *DB) ExecWithReplacements(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return d.db.NamedExecContext(ctx, query, arg)
}

// NewDB instantiates a DB with an exporter to report metrics.
func NewDB(cfg sql.Config) *DB {
	return &DB{
		cfg: cfg,
	}
}

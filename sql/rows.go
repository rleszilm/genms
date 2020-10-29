package sql

// Rows defines the result interface for a SQL query.
type Rows interface {
	Row

	Next() bool
	NextResultSet() bool
	Close() error
}

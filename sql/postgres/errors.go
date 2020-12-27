package postgres_sql

import "errors"

var (
	// ErrBadDriverValue is returned when the value given by a driver cannot be
	// scanned into the object.
	ErrBadDriverValue = errors.New("bad driver value")
)

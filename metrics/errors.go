package metrics

import "errors"

var (
	// ErrInvalidExporter is returned when creating a metrics service with a non-supported
	// exporter.
	ErrInvalidExporter = errors.New("invalid exporter")
)

package exporter

// Exporter interface is a wrapper for OpenCensus Exporter.
type Exporter interface {
	// Start starts the exporter
	Start() error
	// Stop stops the exporter
	Stop() error
}

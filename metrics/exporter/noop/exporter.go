package noop_exporter

// Exporter is the Datadog metric exporter.
type Exporter struct {
}

// Start initialize registers the Datadog exporter and starts metric collection.
func (e *Exporter) Start() error {
	// exporter does not require a start.
	return nil
}

// Stop stops the exporter from publishing metrics.
func (e *Exporter) Stop() error {
	// exporter does not require being stopped
	return nil
}

// New instantiates and returns a noop Exporter.
func New() *Exporter {
	return &Exporter{}
}

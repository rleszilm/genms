package prometheus_exporter

import (
	"contrib.go.opencensus.io/exporter/prometheus"
	rest_service "github.com/rleszilm/gen_microservice/service/rest"
)

// Exporter is the prometheus metric exporter.
type Exporter struct {
	*prometheus.Exporter
	config *Config
	server *rest_service.Server
}

// Start initialize registers the prometheus exporter and starts metric collection.
func (e *Exporter) Start() error {
	e.server.WithRoute(e.config.RequestPath, e)
	return nil
}

// Stop stops the exporter from publishing metrics.
func (e *Exporter) Stop() error {
	// exporter does not require being stopped
	return nil
}

// NewExporter instantiates and returns a prometheus Exporter.
func NewExporter(config *Config, server *rest_service.Server) (*Exporter, error) {
	pe, err := prometheus.NewExporter(
		prometheus.Options{
			Namespace:   config.Namespace,
			ConstLabels: config.Labels,
		},
	)
	if err != nil {
		return nil, err
	}

	return &Exporter{
		Exporter: pe,
		config:   config,
		server:   server,
	}, nil
}

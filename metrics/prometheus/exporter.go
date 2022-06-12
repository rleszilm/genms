package prometheus_exporter

import (
	"context"

	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/rleszilm/genms/service"
	http_service "github.com/rleszilm/genms/service/http"
)

// Exporter is the prometheus metric exporter.
type Exporter struct {
	service.UnimplementedService

	*prometheus.Exporter
	config *Config
	server *http_service.Server
}

// Initialize implements service.Service.Initialize.
func (e *Exporter) Initialize(ctx context.Context) error {
	return nil
}

// Shutdown implements service.Service.Shutdown.
func (e *Exporter) Shutdown(ctx context.Context) error {
	return nil
}

// String implements service.Service.String
func (e *Exporter) String() string {
	return "genms-metrics-prometheus"
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
func NewExporter(config *Config, server *http_service.Server) (*Exporter, error) {
	pe, err := prometheus.NewExporter(
		prometheus.Options{
			Namespace:   config.Namespace,
			ConstLabels: config.Labels,
		},
	)
	if err != nil {
		return nil, err
	}

	exp := &Exporter{
		Exporter: pe,
		config:   config,
		server:   server,
	}

	server.WithDependencies(exp)
	return exp, nil
}

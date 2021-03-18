package metrics

import (
	"context"

	"github.com/rleszilm/genms/metrics/exporter"
)

// Service is the object which manages the metrics exporter.
type Service struct {
	exporter exporter.Exporter
}

// Initialize implements service.Service.Initialize.
func (m *Service) Initialize(ctx context.Context) error {
	return m.exporter.Start()
}

// Shutdown implements service.Service.Shutdown.
func (m *Service) Shutdown(ctx context.Context) error {
	return m.exporter.Stop()
}

// NewService instantiates a Service with an exporter to report metrics.
func NewService(exporter exporter.Exporter) *Service {
	return &Service{
		exporter: exporter,
	}
}

package metrics

import (
	"context"

	"github.com/rleszilm/genms/service"
)

// Service is the object which manages the metrics exporter.
type Service struct {
	service.Dependencies
}

// Initialize implements service.Service.Initialize.
func (s *Service) Initialize(ctx context.Context) error {
	return nil
}

// Shutdown implements service.Service.Shutdown.
func (s *Service) Shutdown(ctx context.Context) error {
	return nil
}

// String implements service.Service.String
func (s *Service) String() string {
	return "metrics"
}

// NewService instantiates a Service with an exporter to report metrics.
func NewService(exporters ...service.Service) *Service {
	svc := &Service{}
	svc.WithDependencies(exporters...)
	return svc
}

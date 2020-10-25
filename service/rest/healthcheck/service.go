package healthcheck

import (
	"context"
	"net/http"

	rest_service "github.com/rleszilm/gen_microservice/service/rest"
)

// Service function returns an http.Handler that handles system status request.
type Service struct {
	config *Config
	server *rest_service.Server
}

// Initialize implements the service.Initialize interface for Service.
func (s *Service) Initialize(_ context.Context) error {
	s.server.WithRoute(s.config.RequestPath, s)
	return nil
}

// Shutdown implements the service.Shutdown interface for Service.
func (s *Service) Shutdown(_ context.Context) error {
	return nil
}

// Name implements service.Name
func (s *Service) Name() string {
	if s.config.Name != "" {
		return s.config.Name
	}
	return "healthcheck"
}

// Healthy is the handler that checks whether the service is ready to service.
func (s *Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("system is available"))
}

// NewService instantitates a Service server.
func NewService(config *Config, server *rest_service.Server) *Service {
	return &Service{
		config: config,
		server: server,
	}
}

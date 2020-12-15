package healthcheck

import (
	"context"
	"net/http"

	"github.com/rleszilm/gen_microservice/service"
	rest_service "github.com/rleszilm/gen_microservice/service/rest"
)

// Service function returns an http.Handler that handles system status request.
type Service struct {
	service.Deps

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

// NameOf implements service.NameOf
func (s *Service) NameOf() string {
	if s.config.Name != "" {
		return s.config.Name
	}
	return "healthcheck"
}

// String returns a sting identifier
func (s *Service) String() string {
	return s.NameOf()
}

// Healthy is the handler that checks whether the service is ready to service.
func (s *Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("system is available"))
}

// NewService instantitates a Service server.
func NewService(config *Config, server *rest_service.Server) *Service {
	svc := &Service{
		config: config,
		server: server,
	}

	server.WithDependencies(svc)

	return svc
}

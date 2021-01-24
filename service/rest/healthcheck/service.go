package healthcheck

import (
	"context"
	"net/http"
	"path"

	"github.com/rleszilm/gen_microservice/service"
	rest_service "github.com/rleszilm/gen_microservice/service/rest"
)

// Service function returns an http.Handler that handles system status request.
type Service struct {
	service.Deps

	healthyFunc http.HandlerFunc
	readyFunc   http.HandlerFunc

	config *Config
	server *rest_service.Server
}

// Initialize implements the service.Initialize interface for Service.
func (s *Service) Initialize(_ context.Context) error {
	s.server.WithRouteFunc(s.config.RequestPrefix, s.ServeHTTP)
	s.server.WithRouteFunc(path.Join(s.config.RequestPrefix, "ready"), s.ServeReady)
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

// ServeHTTP is the handler that checks whether the service is ready to service.
func (s *Service) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.healthyFunc(w, req)
}

// ServeHealthy runs the system health check.
func (s *Service) ServeHealthy(w http.ResponseWriter, req *http.Request) {
	if s.healthyFunc != nil {
		s.healthyFunc(w, req)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("system is healthy"))
}

// ServeReady runs the system ready check.
func (s *Service) ServeReady(w http.ResponseWriter, req *http.Request) {
	if s.readyFunc != nil {
		s.readyFunc(w, req)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("system is ready"))
}

// NewService instantitates a Service server.
func NewService(config *Config, server *rest_service.Server) *Service {
	svc := &Service{
		config: config,
		server: server,
	}

	if config != nil {
		svc.healthyFunc = config.HealthyFunc
		svc.readyFunc = config.ReadyFunc
	}

	server.WithDependencies(svc)

	return svc
}

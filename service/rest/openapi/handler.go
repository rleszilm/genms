package openapi

import (
	"context"
	"net/http"

	"github.com/rleszilm/genms/service"
	rest_service "github.com/rleszilm/genms/service/rest"
)

// Service function returns an http.Handler that handles system status request.
type Service struct {
	service.Dependencies

	config  *Config
	server  *rest_service.Server
	handler http.Handler
}

// Initialize implements the service.Initialize interface for Service.
func (s *Service) Initialize(_ context.Context) error {
	if s.config.Enabled {
		s.server.WithRoute(s.config.RequestPrefix, s.handler)
	}
	return nil
}

// Shutdown implements the service.Shutdown interface for Service.
func (s *Service) Shutdown(_ context.Context) error {
	return nil
}

// ID implements service.ID
func (s *Service) ID() string {
	if s.config.Name != "" {
		return s.config.Name
	}
	return "genms-openapi"
}

// NewService returns a new Service.
func NewService(config *Config, server *rest_service.Server) *Service {
	svc := &Service{
		config:  config,
		server:  server,
		handler: http.StripPrefix(config.RequestPrefix, http.FileServer(http.Dir(config.Dir))),
	}

	server.WithDependencies(svc)
	return svc
}

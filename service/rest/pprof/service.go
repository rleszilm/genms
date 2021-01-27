package pprof

import (
	"context"
	"net/http/pprof"
	"runtime"

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
	if s.config.Enabled {
		runtime.SetBlockProfileRate(1)
		s.server.WithRouteFunc("/debug/pprof/", pprof.Index)
		s.server.WithRouteFunc("/debug/pprof/cmdline", pprof.Cmdline)
		s.server.WithRouteFunc("/debug/pprof/profile", pprof.Profile)
		s.server.WithRouteFunc("/debug/pprof/symbol", pprof.Symbol)
		s.server.WithRouteFunc("/debug/pprof/trace", pprof.Trace)
	}
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
	return "pprof"
}

// String returns a sting identifier
func (s *Service) String() string {
	return s.NameOf()
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

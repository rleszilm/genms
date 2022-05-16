package pprof

import (
	"context"
	"net/http/pprof"
	"runtime"

	"github.com/rleszilm/genms/log"
	"github.com/rleszilm/genms/service"
	http_service "github.com/rleszilm/genms/service/http"
)

var (
	logs = log.NewChannel("pprof")
)

// Service function returns an http.Handler that handles system status request.
type Service struct {
	service.UnimplementedService

	config *Config
	server *http_service.Server
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

// ID implements service.ID
func (s *Service) ID() string {
	if s.config.Name != "" {
		return s.config.Name
	}
	return "genms-pprof"
}

// NewService instantitates a Service server.
func NewService(config *Config, server *http_service.Server) *Service {
	svc := &Service{
		config: config,
		server: server,
	}

	server.WithDependencies(svc)
	return svc
}

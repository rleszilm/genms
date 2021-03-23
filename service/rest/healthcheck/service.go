package healthcheck

import (
	"context"
	"net/http"
	"path"

	"github.com/rleszilm/genms/log"
	"github.com/rleszilm/genms/service"
	rest_service "github.com/rleszilm/genms/service/rest"
)

var (
	logs = log.NewChannel("health")
)

// Service function returns an http.Handler that handles system status request.
type Service struct {
	service.Dependencies

	healthyFunc http.HandlerFunc
	readyFunc   http.HandlerFunc

	config *Config
	server *rest_service.Server
}

// Initialize implements the service.Initialize interface for Service.
func (s *Service) Initialize(_ context.Context) error {
	if s.config.Enabled {
		logs.Debug("health route:", s.config.RequestPrefix)
		s.server.WithRouteFunc(s.config.RequestPrefix, s.ServeHealthy)
		logs.Debug("ready route:", path.Join(s.config.RequestPrefix, "ready"))
		s.server.WithRouteFunc(path.Join(s.config.RequestPrefix, "ready"), s.ServeReady)
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
	return "healthcheck"
}

// String returns a sting identifier
func (s *Service) String() string {
	return s.NameOf()
}

// ServeHealthy runs the system health check.
func (s *Service) ServeHealthy(w http.ResponseWriter, req *http.Request) {
	logs.Debug("serving healthy status")
	if s.healthyFunc != nil {
		logs.Trace("using custom healthy logic")
		s.healthyFunc(w, req)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("system is healthy"))
}

// ServeReady runs the system ready check.
func (s *Service) ServeReady(w http.ResponseWriter, req *http.Request) {
	logs.Debug("serving ready status")
	if s.readyFunc != nil {
		logs.Trace("using custom status logic")
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

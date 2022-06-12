package healthcheck

import (
	"context"
	"net/http"
	"path"

	"github.com/rleszilm/genms/log"
	"github.com/rleszilm/genms/service"
	http_service "github.com/rleszilm/genms/service/http"
)

var (
	logs = log.NewChannel("healthcheck")
)

// StatusFunc is a function that returns true/false depending on whether the system
// is in the given status.
type StatusFunc func() (bool, error)

// Service function returns an http.Handler that handles system status request.
type Service struct {
	service.UnimplementedService

	healthyFunc StatusFunc
	readyFunc   StatusFunc

	config *Config
	server *http_service.Server
}

// Initialize implements the service.Initialize interface for Service.
func (s *Service) Initialize(_ context.Context) error {
	if s.config.Enabled {
		logs.Debug("health route:", s.config.RequestPrefix)
		s.server.WithRouteFunc(s.config.RequestPrefix, s.ServeHealthy)
		logs.Debug("ready route:", path.Join(s.config.RequestPrefix, "ready"))
		s.server.WithRouteFunc(path.Join(s.config.RequestPrefix, "ready"), s.ServeReady)
	} else {
		logs.Debug("health checks disabled")
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
	return "genms-healthcheck"
}

// WithHealthyFunc sets the health check function.
func (s *Service) WithHealthyFunc(f StatusFunc) *Service {
	s.healthyFunc = f
	return s
}

// WithReadyFunc sets the health check function.
func (s *Service) WithReadyFunc(f StatusFunc) *Service {
	s.readyFunc = f
	return s
}

// ServeHealthy runs the system health check.
func (s *Service) ServeHealthy(w http.ResponseWriter, req *http.Request) {
	logs.Debug("serving healthy status")
	if s.healthyFunc != nil {
		logs.Trace("using custom healthy logic")
		if ok, err := s.healthyFunc(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("system is not healthy:" + err.Error()))
			return
		} else if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("system is not healthy"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("system is healthy"))
}

// ServeReady runs the system ready check.
func (s *Service) ServeReady(w http.ResponseWriter, req *http.Request) {
	logs.Debug("serving ready status")
	if s.readyFunc != nil {
		logs.Trace("using custom status logic")
		if ok, err := s.readyFunc(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("system is not ready:" + err.Error()))
			return
		} else if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("system is not ready"))
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("system is ready"))
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

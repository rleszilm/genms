package healthcheck

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/rleszilm/genms/logging"
	"github.com/rleszilm/genms/service"
	http_service "github.com/rleszilm/genms/service/http"
)

var (
	logs = logging.NewChannel("healthcheck")
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
		if err := s.server.WithRouteFunc(s.config.RequestPrefix, s.ServeHealthy); err != nil {
			return err
		}
		logs.Debug("ready route:", path.Join(s.config.RequestPrefix, "ready"))
		if err := s.server.WithRouteFunc(path.Join(s.config.RequestPrefix, "ready"), s.ServeReady); err != nil {
			return err
		}
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
		if ok, err := s.healthyFunc(); err != nil || !ok {
			respond(w, "system is not healthy:", nil)
			return
		}
	}

	respond(w, "system is healthy", nil)
}

// ServeReady runs the system ready check.
func (s *Service) ServeReady(w http.ResponseWriter, req *http.Request) {
	logs.Debug("serving ready status")
	if s.readyFunc != nil {
		logs.Trace("using custom status logic")
		if ok, err := s.readyFunc(); err != nil || !ok {
			respond(w, "system is not ready:", nil)
			return
		}
	}

	respond(w, "system is ready", nil)
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

func respond(w http.ResponseWriter, resp string, err error) {
	if asStatusError, ok := err.(*http_service.StatusError); ok && err != nil {
		w.WriteHeader(asStatusError.StatusCode())
		if _, err := w.Write([]byte(fmt.Sprint(resp, err))); err != nil {
			logs.Error("could not write response:", err)
		}
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(fmt.Sprint(resp, err))); err != nil {
			logs.Error("could not write response:", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(resp)); err != nil {
		logs.Error("could not write response:", err)
	}
}

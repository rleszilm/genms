package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rleszilm/genms/logging"
)

var (
	logs = logging.NewChannel("service")
)

// Manager maintains a list of app.Service interfaces. The Manager is intended to Initialize, Run
// and Stop all services in the application. User should be aware of the order in which services are
// added.
type Manager struct {
	signals chan Signal
	svcs    Services
}

// Register adds an Service interface to its slice. Services will start and stop in the order which
// they are Registered.
func (m *Manager) Register(svcs ...Service) {
	l := len(m.svcs)
	for i := 0; i < len(svcs); i++ {
		svcs[i].withServiceID(ServiceID(i + l))
	}

	m.svcs = append(m.svcs, svcs...)
}

// Initialize iterates through the list of services registered and invokes their respective Initialize method.
func (m *Manager) Initialize(ctx context.Context) error {
	m.signals = make(chan Signal, len(m.svcs))

	done := make(chan error)
	go func() {
		defer close(done)

		if cycle, err := m.svcs.Sort(); err != nil {
			if cycle != nil {
				logs.Fatal("cannot start services due to cyclical dependencies:", cycle)
			}
			done <- err
			return
		}

		for i := 0; i < len(m.svcs); i++ {
			svc := m.svcs[i]
			logs.Infof("starting service <%d> %s(%T)", svc.ServiceID(), svc.String(), svc)
			if err := svc.Initialize(ctx); err != nil {
				done <- fmt.Errorf("cannot start service: <%d> %s - %w", svc.ServiceID(), svc.String(), err)
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}

// Shutdown iterates through the list of services registered and invokes their respective shutdown method
// and logs any errors returned.
func (m *Manager) Shutdown(ctx context.Context) error {
	defer close(m.signals)

	done := make(chan error)
	go func() {
		defer close(done)

		if cycle, err := m.svcs.Sort(); err != nil {
			if cycle != nil {
				logs.Fatal("cannot shut down services due to cyclical dependencies:", cycle)
			}
			done <- err
			return
		}

		for i := len(m.svcs); i > 0; i-- {
			svc := m.svcs[i-1]
			logs.Infof("shutting down service <%d> %s(%T)", svc.ServiceID(), svc.String(), svc)
			if err := svc.Shutdown(ctx); err != nil {
				done <- err
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}

// Signal accepts a signal from a service.
func (m *Manager) Signal(s Signal) {
	m.signals <- s
}

// Wait blocks until an os signal tells the manager to shutdown.
func (m *Manager) Wait() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)
	for {
		select {
		case <-c:
			return
		case s := <-m.signals:
			err := s.Error()
			if err != nil {
				logs.Error("manager shutting down due to component error: <%d> %s", s.ServiceID(), err)
				return
			}
		}
	}
}

// NewManager returns a new service Manager.
func NewManager() *Manager {
	return &Manager{
		svcs: []Service{},
	}
}

package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

// Manager maintains a list of app.Service interfaces. The Manager is intended to Initialize, Run
// and Stop all services in the application. User should be aware of the order in which services are
// added.
type Manager struct {
	svcs Services
}

// Register adds an Service interface to its slice. Services will start and stop in the order which
// they are Registered.
func (m *Manager) Register(svcs ...Service) {
	m.svcs = append(m.svcs, svcs...)
}

// Initialize iterates through the list of services registered and invokes their respective Initialize method.
func (m *Manager) Initialize(ctx context.Context) error {
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
			logs.Infof("starting service %s(%T)\n", svc.NameOf(), svc)
			if err := svc.Initialize(ctx); err != nil {
				done <- fmt.Errorf("cannot start service: %s - %w", svc.NameOf(), err)
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
			logs.Infof("shutting down service %s(%T)\n", svc.NameOf(), svc)
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

// Wait blocks until an os signal tells the manager to shutdown.
func (m *Manager) Wait() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	<-c
}

// NewManager returns a new service Manager.
func NewManager() *Manager {
	return &Manager{
		svcs: []Service{},
	}
}

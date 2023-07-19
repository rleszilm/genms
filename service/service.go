package service

import (
	"context"
	"fmt"
)

// ServiceID is a service identifier assigned to a service by the manager.
type ServiceID int64

// Service describes a long-running instance whose life-cycle should start with Initialize and end
// with a Shutdown call.
type Service interface {
	// Initialize function initializes and starts the service. Initialize should panic in case of
	// any errors. It is intended that Initialize be called only once during the service life-cycle.
	Initialize(context.Context) error

	// Shutdown closes the long-running instance, or service.
	Shutdown(context.Context) error

	// String returns the identifier of a service. This must be unique if there are multiple instances
	// of the same service.
	String() string

	// ServiceID returns the services service id.
	ServiceID() ServiceID

	// withSID assigns a SID to the service.
	withServiceID(ServiceID)

	// dependencies returns the servers dependencies.
	dependencies() Services

	// WithDependencies adds a dependency to the service.
	WithDependencies(...Service)

	// Signaller returns the signaller assigned to the service.
	Signaller() Signaller

	// WithSignaller specifies the signaller that the service should use to report statuses.
	withSignaller(Signaller)

	// mustEmbedUnimplementedService
	mustEmbedUnimplementedService()
}

// Signal describes a message sent from a service back to the manager.
type Signal interface {
	ServiceID() ServiceID
	Message() string
	Error() error
}

// Signaller describes an interface that transmits signals between a service and the manager.
type Signaller interface {
	Signal(Signal)
}

// UnimplementedService is an implementation of Service that provides base logic and returns errors when user
// implemented methods are called but not defined.
type UnimplementedService struct {
	Services
	sid       ServiceID
	signaller Signaller
}

// Initialize implements Service.Initialize.
func (u *UnimplementedService) Initialize(_ context.Context) error {
	return fmt.Errorf("Service.Initialize is not defined")
}

// Shutdown implements Service.Shutdown.
func (u *UnimplementedService) Shutdown(_ context.Context) error {
	return fmt.Errorf("Service.Shutdown is not defined")
}

// String implements Service.String.
func (u *UnimplementedService) String() string {
	return "Service.String is not defined"
}

// ServiceID implements Service.SID.
func (u *UnimplementedService) ServiceID() ServiceID {
	return u.sid
}

// withServiceID implements Service.withServiceID.
func (u *UnimplementedService) withServiceID(s ServiceID) {
	u.sid = s
}

// WithDependencies implements Service.WithDependencies.
func (u *UnimplementedService) WithDependencies(svcs ...Service) {
	u.Services = append(u.Services, svcs...)
}

// dependencies implements Service.dependencies.
func (u *UnimplementedService) dependencies() Services {
	return u.Services
}

// Signaller implements Service.Signaller.
func (u *UnimplementedService) Signaller() Signaller {
	return u.signaller
}

// withSignaller implements Service.withSignaller.
func (u *UnimplementedService) withSignaller(s Signaller) {
	u.signaller = s
}

func (u *UnimplementedService) mustEmbedUnimplementedService() {
}

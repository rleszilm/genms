package service

import (
	"context"

	"github.com/rleszilm/genms/log"
)

var (
	logs = log.NewChannel("service")
)

// Service describes a long-running instance whose life-cycle should start with Initialize and end
// with a Shutdown call.
type Service interface {
	// Initialize function initializes and starts the service. Initialize should panic in case of
	// any errors. It is intended that Initialize be called only once during the service life-cycle.
	Initialize(context.Context) error

	// Shutdown closes the long-running instance, or service.
	Shutdown(context.Context) error

	// ID returns the identifier of a service. This must be unique if there are multiple instances
	// of the same service.
	ID() string

	// Dependants returns the servers dependencies.
	Dependants() Services

	// WithDependency adds a dependency to the service.
	WithDependencies(...Service)
}

// Listener is a Service that accepts connections and does work based on the requests made.
type Listener interface {
	Service

	// Scheme returns the request scheme of the underlying server.
	Scheme() string

	// Addr returns the address that the Listener is listening on.
	Addr() string
}

// Package greeter is generated by protoc-gen-go-genms. *DO NOT EDIT*
package greeter

import (
	context "context"
	service "github.com/rleszilm/genms/service"
	grpc "github.com/rleszilm/genms/service/grpc"
	rest "github.com/rleszilm/genms/service/rest"
	grpc1 "google.golang.org/grpc"
)

// WithRestServerService implements WithRestService
type WithRestServerService struct {
	service.Dependencies

	impl       WithRestServer
	grpcServer *grpc.Server
	restServer *rest.Server
}

// Initialize implements service.Service.Initialize
func (s *WithRestServerService) Initialize(ctx context.Context) error {
	s.grpcServer.WithService(func(server *grpc1.Server) {
		RegisterWithRestServer(server, s.impl)
	})

	if err := s.restServer.WithGrpcProxyHandler(ctx, "WithRest", RegisterWithRestHandlerFromEndpoint); err != nil {
		return err
	}

	return nil
}

// Shutdown implements service.Service.Shutdown
func (s *WithRestServerService) Shutdown(_ context.Context) error {
	return nil
}

// NameOf returns the name of the service
func (s *WithRestServerService) NameOf() string {
	return "with-rest"
}

// String returns the string name of the service
func (s *WithRestServerService) String() string {
	return s.NameOf()
}

// NewWithRestServerService returns a new WithRestServerService
func NewWithRestServerService(impl WithRestServer, grpcServer *grpc.Server, restServer *rest.Server) *WithRestServerService {
	server := &WithRestServerService{
		impl:       impl,
		grpcServer: grpcServer,
		restServer: restServer,
	}

	if asService, ok := impl.(service.Service); ok {
		server.WithDependencies(asService)
	}

	grpcServer.WithDependencies(server)
	restServer.WithDependencies(server)

	return server
}

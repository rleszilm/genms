// Code generated by protoc-gen-microservice. DO NOT EDIT.
package greeter

import (
	"context"
	"github.com/rleszilm/gen_microservice/service"
	grpc_service "github.com/rleszilm/gen_microservice/service/grpc"
	"google.golang.org/grpc"
	rest_service "github.com/rleszilm/gen_microservice/service/rest"
	graphql_service "github.com/rleszilm/gen_microservice/service/graphql"
	// @@protoc_insertion_point(genms-imports)
)

var (
	// @@protoc_insertion_point(genms-variables)
)

const (
	// @@protoc_insertion_point(genms-constants)
)

// WithRestServerService implements WithRestService
type WithRestServerService struct {
	service.Deps
	WithRestServer

	grpcServer *grpc_service.Server
	restServer *rest_service.Server
	
}

// Initialize implements service.Service.Initialize
func (s *WithRestServerService) Initialize(ctx context.Context) error {
	s.grpcServer.WithService(func(server *grpc.Server) {
		RegisterWithRestServer(server, s)
	})

	if err := s.restServer.WithGrpcProxy(ctx, "WithRest", RegisterWithRestHandlerFromEndpoint); err != nil {
		return err
	}
	
	return nil
}

// Shutdown implements service.Service.Shutdown
func (s *WithRestServerService) Shutdown(_ context.Context) error {
	return nil
}

func (s *WithRestServerService) Name() string {
	return "WithRest"
}

func (s *WithRestServerService) String() string {
	return s.Name()
}

// NewWithRestServerService returns a new WithRestServerService
func NewWithRestServerService(impl WithRestServer, grpcServer *grpc_service.Server, restServer *rest_service.Server) *WithRestServerService {
	server := &WithRestServerService{
		WithRestServer: impl,
		grpcServer: grpcServer,
		restServer: restServer,
		
	}

	grpcServer.WithDependencies(server)
	restServer.WithDependencies(server)
	

	return server
}
// WithGraphQLServerService implements WithGraphQLService
type WithGraphQLServerService struct {
	service.Deps
	WithGraphQLServer

	grpcServer *grpc_service.Server
	
	graphqlServer *graphql_service.Server
}

// Initialize implements service.Service.Initialize
func (s *WithGraphQLServerService) Initialize(ctx context.Context) error {
	s.grpcServer.WithService(func(server *grpc.Server) {
		RegisterWithGraphQLServer(server, s)
	})

	
	if err := s.graphqlServer.WithGrpcProxy(ctx, "WithGraphQL", RegisterWithGraphQLGraphqlWithOptions); err != nil {
		return err
	}
	return nil
}

// Shutdown implements service.Service.Shutdown
func (s *WithGraphQLServerService) Shutdown(_ context.Context) error {
	return nil
}

func (s *WithGraphQLServerService) Name() string {
	return "WithGraphQL"
}

func (s *WithGraphQLServerService) String() string {
	return s.Name()
}

// NewWithGraphQLServerService returns a new WithGraphQLServerService
func NewWithGraphQLServerService(impl WithGraphQLServer, grpcServer *grpc_service.Server, graphqlServer *graphql_service.Server) *WithGraphQLServerService {
	server := &WithGraphQLServerService{
		WithGraphQLServer: impl,
		grpcServer: grpcServer,
		
		graphqlServer: graphqlServer,
	}

	grpcServer.WithDependencies(server)
	
	graphqlServer.WithDependencies(server)

	return server
}
// WithRestAndGraphQLServerService implements WithRestAndGraphQLService
type WithRestAndGraphQLServerService struct {
	service.Deps
	WithRestAndGraphQLServer

	grpcServer *grpc_service.Server
	restServer *rest_service.Server
	graphqlServer *graphql_service.Server
}

// Initialize implements service.Service.Initialize
func (s *WithRestAndGraphQLServerService) Initialize(ctx context.Context) error {
	s.grpcServer.WithService(func(server *grpc.Server) {
		RegisterWithRestAndGraphQLServer(server, s)
	})

	if err := s.restServer.WithGrpcProxy(ctx, "WithRestAndGraphQL", RegisterWithRestAndGraphQLHandlerFromEndpoint); err != nil {
		return err
	}
	if err := s.graphqlServer.WithGrpcProxy(ctx, "WithRestAndGraphQL", RegisterWithRestAndGraphQLGraphqlWithOptions); err != nil {
		return err
	}
	return nil
}

// Shutdown implements service.Service.Shutdown
func (s *WithRestAndGraphQLServerService) Shutdown(_ context.Context) error {
	return nil
}

func (s *WithRestAndGraphQLServerService) Name() string {
	return "WithRestAndGraphQL"
}

func (s *WithRestAndGraphQLServerService) String() string {
	return s.Name()
}

// NewWithRestAndGraphQLServerService returns a new WithRestAndGraphQLServerService
func NewWithRestAndGraphQLServerService(impl WithRestAndGraphQLServer, grpcServer *grpc_service.Server, restServer *rest_service.Server, graphqlServer *graphql_service.Server) *WithRestAndGraphQLServerService {
	server := &WithRestAndGraphQLServerService{
		WithRestAndGraphQLServer: impl,
		grpcServer: grpcServer,
		restServer: restServer,
		graphqlServer: graphqlServer,
	}

	grpcServer.WithDependencies(server)
	restServer.WithDependencies(server)
	graphqlServer.WithDependencies(server)

	return server
}
// @@protoc_insertion_point(genms-logic)

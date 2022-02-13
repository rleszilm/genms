package graphql_service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/rleszilm/genms/service"
	grpc_service "github.com/rleszilm/genms/service/grpc"
	rest_service "github.com/rleszilm/genms/service/rest"
	"github.com/rleszilm/grpc-graphql-gateway/options"
	"github.com/rleszilm/grpc-graphql-gateway/runtime"
)

// GrpcProxyRoutes is a function that registers http routes to a grpc server.
type GrpcProxyRoutes func(*runtime.ServeMux, *options.ServerOptions) error

// Server is a service.Service that handles rest requests.
type Server struct {
	service.Dependencies
	name       string
	config     *Config
	restServer *rest_service.Server
}

// Initialize implements the Server.Initialize interface for Server.
func (s *Server) Initialize(ctx context.Context) error {
	return nil
}

// Shutdown implements the Server.Shutdown interface for Server.
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// ID implements Server.ID()
func (s *Server) ID() string {
	return s.name
}

// Scheme implements service.Listen.Scheme
func (s *Server) Scheme() string {
	if s.config.TLS.Enabled {
		return "graphqls"
	}
	return "graphql"
}

// Addr implements service.Listen.Addr
func (s *Server) Addr() string {
	return s.config.Addr
}

// WithGrpcProxy adds rest methods that proxy to a grpc server.
func (s *Server) WithGrpcProxy(_ context.Context, proxy *grpc_service.Proxy, proxyFunc GrpcProxyRoutes) error {
	if !proxy.Enabled {
		return nil
	}

	proxyOpts := &options.ServerOptions{
		Host:         proxy.Addr,
		WithInsecure: proxy.Insecure,
	}

	proxyMux := runtime.NewServeMux()
	if err := s.restServer.WithRoute(proxy.Pattern, proxyMux); err != nil {
		return err
	}

	if err := proxyFunc(proxyMux, proxyOpts); err != nil {
		return err
	}

	return nil
}

// NewServer returns a new Server.
func NewServer(name string, config *Config) (*Server, error) {
	restServer := config.RestServer
	if restServer == nil {
		restConfig := &rest_service.Config{}
		copier.Copy(restConfig, config)

		rest, err := rest_service.NewServer(name+"-rest", restConfig)
		if err != nil {
			return nil, err
		}
		restServer = rest
	}

	server := &Server{
		name:       name,
		config:     config,
		restServer: restServer,
	}

	server.restServer.WithDependencies(server)

	return server, nil
}

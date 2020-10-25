package graphql_service

import (
	"context"

	"github.com/jinzhu/copier"
	rest_service "github.com/rleszilm/gen_microservice/service/rest"
	"github.com/rleszilm/grpc-graphql-gateway/options"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

// GraphqlProxy is a function that registers http routes to a grpc server.
type GraphqlProxy func(context.Context, *runtime.ServeMux, *options.ServerOptions) error

// Server is a service.Service that handles rest requests.
type Server struct {
	name         string
	config       *Config
	server       *rest_service.Server
	sharedServer bool
	proxyMux     *runtime.ServeMux
}

// Initialize implements the Server.Initialize interface for Server.
func (s *Server) Initialize(ctx context.Context) error {
	if !s.sharedServer {
		return s.server.Initialize(ctx)
	}

	return nil
}

// Shutdown implements the Server.Shutdown interface for Server.
func (s *Server) Shutdown(ctx context.Context) error {
	if !s.sharedServer {
		return s.server.Shutdown(ctx)
	}
	return nil
}

// Name implements Server.Name()
func (s *Server) Name() string {
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

// WithProxy adds rest methods that proxy to a grpc server.
func (s *Server) WithProxy(ctx context.Context, proxy GraphqlProxy) error {
	options := &options.ServerOptions{
		Host:         s.config.ProxyGrpc.Addr,
		WithInsecure: s.config.ProxyGrpc.Secure,
	}

	if err := proxy(ctx, s.proxyMux, options); err != nil {
		return err
	}

	return nil
}

// NewServer returns a new Server.
func NewServer(name string, config *Config) (*Server, error) {
	proxyMux := runtime.NewServeMux()

	shared := config.RestServer != nil
	if !shared {
		restConfig := &rest_service.Config{}
		copier.Copy(restConfig, config)

		rest, err := rest_service.NewServer(name+"-rest", restConfig)
		if err != nil {
			return nil, err
		}
		config.RestServer = rest
	}

	config.RestServer.WithRoute(config.ProxyGrpc.Pattern, proxyMux)

	return &Server{
		name:         name,
		config:       config,
		server:       config.RestServer,
		sharedServer: shared,
		proxyMux:     proxyMux,
	}, nil
}

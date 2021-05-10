package rest_service

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rleszilm/genms/service"
	grpc_service "github.com/rleszilm/genms/service/grpc"
	"google.golang.org/grpc"
)

// GrpcProxyRoutes is a function that registers http routes to a grpc server.
type GrpcProxyRoutes func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

// Server is a service.Service that handles rest requests.
type Server struct {
	service.Dependencies
	name    string
	config  *Config
	server  *http.Server
	mux     *http.ServeMux
	grpcMux *runtime.ServeMux
}

// Initialize implements the Server.Initialize interface for Server.
func (s *Server) Initialize(ctx context.Context) error {
	go func() {
		if s.config.TLS.Enabled {
			if err := s.server.ListenAndServeTLS(s.config.TLS.Cert, s.config.TLS.Key); err != nil {
				log.Fatalln("Error serving rest requests", err)
			}
		} else {
			if err := s.server.ListenAndServe(); err != nil {
				log.Fatalln("Error serving rest requests", err)
			}
		}
	}()

	return nil
}

// Shutdown implements the Server.Shutdown interface for Server.
func (s *Server) Shutdown(_ context.Context) error {
	return s.server.Close()
}

// NameOf implements Server.NameOf()
func (s *Server) NameOf() string {
	return s.name
}

// String implements Server.String()
func (s *Server) String() string {
	return s.name
}

// Scheme implements service.Listen.Scheme
func (s *Server) Scheme() string {
	if s.config.TLS.Enabled {
		return "https"
	}
	return "http"
}

// Addr implements service.Listen.Addr
func (s *Server) Addr() string {
	return s.config.Addr
}

// WithRoute adds a route to the rest service
func (s *Server) WithRoute(pattern string, handler http.Handler) error {
	s.mux.Handle(pattern, http.StripPrefix(pattern, handler))
	return nil
}

// WithRouteFunc adds a route to the rest service
func (s *Server) WithRouteFunc(pattern string, handler http.HandlerFunc) error {
	s.WithRoute(pattern, handler)
	return nil
}

// WithGrpcProxy adds rest methods that proxy to a grpc server.
func (s *Server) WithGrpcProxy(ctx context.Context, proxy *grpc_service.Proxy, proxyFunc GrpcProxyRoutes) error {
	if !proxy.Enabled {
		return nil
	}

	proxyOpts := []grpc.DialOption{}
	if proxy.Insecure {
		proxyOpts = append(proxyOpts, grpc.WithInsecure())
	}

	if s.grpcMux == nil {
		s.grpcMux = runtime.NewServeMux()
		s.mux.Handle("/", s.grpcMux)
	}

	if err := proxyFunc(ctx, s.grpcMux, proxy.Addr, proxyOpts); err != nil {
		return err
	}

	return nil
}

// NewServer returns a new Server.
func NewServer(name string, config *Config) (*Server, error) {
	mux := http.NewServeMux()

	s := &Server{
		name:   name,
		config: config,
		server: &http.Server{
			Addr:    config.Addr,
			Handler: mux,
		},
		mux: mux,
	}

	return s, nil
}

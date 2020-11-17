package grpc_service

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rleszilm/gen_microservice/service"
	"google.golang.org/grpc"
)

// GrpcService is a function that registers a grpc service against a server.
type GrpcService func(*grpc.Server)

// Server is a service.Service that handles grpc requests.
type Server struct {
	service.Deps
	name       string
	config     *Config
	grpc       *grpc.Server
	services   []GrpcService
	middleware []grpc.UnaryServerInterceptor
}

// Initialize implements the Server.Initialize interface for Server.
func (s *Server) Initialize(_ context.Context) error {
	// set up grpc handler
	grpcOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(s.middleware...)),
	}
	s.grpc = grpc.NewServer(grpcOpts...)

	for _, svc := range s.services {
		svc(s.grpc)
	}

	listener, err := net.Listen(s.config.Transport, s.config.Addr)
	if err != nil {
		return err
	}

	if s.config.TLS.Enabled {
		cer, err := tls.LoadX509KeyPair(s.config.TLS.Cert, s.config.TLS.Key)
		if err != nil {
			return err
		}
		config := &tls.Config{Certificates: []tls.Certificate{cer}}
		listener = tls.NewListener(listener, config)
	}

	go func() {
		if err := s.grpc.Serve(listener); err != nil {
			log.Fatalln("Error serving grpc requests", err)
		}
		log.Println("no longer serving grpc")
	}()

	return nil
}

// Shutdown implements the Server.Shutdown interface for Server.
func (s *Server) Shutdown(_ context.Context) error {
	s.grpc.GracefulStop()
	return nil
}

// Name implements Server.Name()
func (s *Server) Name() string {
	return s.name
}

// String implements Server.String()
func (s *Server) String() string {
	return s.name
}

// Scheme implements service.Listen.Scheme
func (s *Server) Scheme() string {
	if s.config.TLS.Enabled {
		return "grpcs"
	}
	return "grpc"
}

// Addr implements service.Listen.Addr
func (s *Server) Addr() string {
	return s.config.Addr
}

// WithService attaches an implementation of a grpc service to the server.
func (s *Server) WithService(svc GrpcService) {
	s.services = append(s.services, svc)
}

// WithMiddleware adds a middleware to the server. Middleware can be added until Initialize has been called.
func (s *Server) WithMiddleware(m grpc.UnaryServerInterceptor) {
	s.middleware = append(s.middleware, m)
}

// NewServer returns a new Server.
func NewServer(name string, config *Config, middleware ...grpc.UnaryServerInterceptor) (*Server, error) {
	return &Server{
		name:       name,
		config:     config,
		middleware: middleware,
	}, nil
}

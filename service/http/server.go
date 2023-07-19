package http_service

import (
	"context"
	"net/http"
	"reflect"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rleszilm/genms/logging"
	"github.com/rleszilm/genms/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	logs = logging.NewChannel("http")
)

// LocalGrpcProxy is a function that registers http routes against a grpc server implementation.
// Using this method will call the grpc methods directly and may result in middleware being excluded.
type LocalGrpcProxy func(context.Context, *runtime.ServeMux) error

// RemoteGrpcProxy is a function that registers http routes that will proxy
// requests to a remote grpc server.
type RemoteGrpcProxy func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error

// Server is a service.Service that handles rest requests.
type Server struct {
	service.UnimplementedService
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
				logs.Fatal("Error serving http requests", err)
			}
		} else {
			if err := s.server.ListenAndServe(); err != nil {
				logs.Fatal("Error serving http requests", err)
			}
		}
	}()

	return nil
}

// Shutdown implements the Server.Shutdown interface for Server.
func (s *Server) Shutdown(_ context.Context) error {
	return s.server.Close()
}

// String implements Server.String()
func (s *Server) String() string {
	if s.name != "" {
		return "genms-http-" + s.name
	}
	return "genms-http"
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

// Routes returns a mapping of bound routes to handlers.
func (s *Server) Routes() map[string]interface{} {
	out := map[string]interface{}{}

	elem := reflect.ValueOf(s.mux).Elem()
	muxEntry := elem.FieldByName("m")
	for _, key := range muxEntry.MapKeys() {
		handler := muxEntry.MapIndex(key)
		out[handler.FieldByName("pattern").String()] = handler.FieldByName("h").Elem().Type()
	}

	return out
}

// WithRoute adds a route to the http service
func (s *Server) WithRoute(pattern string, handler http.Handler) error {
	s.mux.Handle(pattern, http.StripPrefix(pattern, handler))
	return nil
}

// WithRouteFunc adds a route to the http service
func (s *Server) WithRouteFunc(pattern string, handler http.HandlerFunc) error {
	return s.WithRoute(pattern, handler)
}

// WithLocalGrpcProxy adds http methods that proxy to a grpc server.
func (s *Server) WithLocalGrpcProxy(ctx context.Context, proxy *service.Proxy, proxyFunc LocalGrpcProxy) error {
	if !proxy.Enabled {
		logs.Trace("not adding routes as proxy is disabled")
		return nil
	}

	if s.grpcMux == nil {
		s.grpcMux = runtime.NewServeMux()
		s.mux.Handle("/", s.grpcMux)
	}

	logs.Debugf("adding proxy: %+v", proxy)
	if err := proxyFunc(ctx, s.grpcMux); err != nil {
		return err
	}

	return nil
}

// WithRemoteGrpcProxy adds http methods that proxy to a grpc server.
func (s *Server) WithRemoteGrpcProxy(ctx context.Context, proxy *service.Proxy, proxyFunc RemoteGrpcProxy) error {
	if !proxy.Enabled {
		logs.Trace("not adding routes as proxy is disabled")
		return nil
	}

	proxyOpts := []grpc.DialOption{}
	if proxy.Insecure {
		proxyOpts = append(proxyOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if s.grpcMux == nil {
		s.grpcMux = runtime.NewServeMux()
		s.mux.Handle("/", s.grpcMux)
	}

	logs.Debugf("adding proxy: %+v %+v", proxy, proxyOpts)
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

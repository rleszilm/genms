package main

import (
	"context"
	"log"

	"github.com/rleszilm/genms/protoc-gen-genms/example/greeter"
	"github.com/rleszilm/genms/service"
	grpc_service "github.com/rleszilm/genms/service/grpc"
	http_service "github.com/rleszilm/genms/service/http"
)

type greets struct {
	http *http_service.Server
	greeter.UnimplementedWithHTTPServer
}

func (g *greets) HelloRest(_ context.Context, req *greeter.Message) (*greeter.Message, error) {
	log.Print(g.http.Routes())
	return &greeter.Message{Value: req.GetValue() + " received"}, nil
}

func main() {
	ctx := context.Background()

	// create service  manager
	manager := service.NewManager()

	httpServer, err := http_service.NewServer("http-api",
		&http_service.Config{
			Config: service.Config{
				Transport: "tcp",
				Addr:      ":8081",
			},
		},
	)
	if err != nil {
		log.Fatalln("Unable to instantiate http api server: ", err)
	}
	manager.Register(httpServer)

	grpcServer, err := grpc_service.NewServer("grpc-api",
		&grpc_service.Config{
			Config: service.Config{
				Transport: "tcp",
				Addr:      ":8080",
				TLS: service.TLS{
					Enabled: false,
				},
			},
		},
	)
	if err != nil {
		log.Fatalln("Unable to instantiate grpc api server: ", err)
	}
	manager.Register(grpcServer)

	impl := &greets{
		http: httpServer,
	}

	proxy := &service.Proxy{
		Enabled:  true,
		Prefix:   "/v1/grpc",
		Addr:     ":8080",
		Insecure: true,
	}

	service := greeter.NewWithHTTPService(impl, &greeter.WithHTTPServiceConfig{
		Proxy: proxy,
	})
	service.WithGrpcServer(grpcServer)
	service.WithHttpServer(httpServer)
	manager.Register(service)

	if err := manager.Initialize(ctx); err != nil {
		log.Fatal("Unable to start services: ", err)
	}
	manager.Wait()
	if err := manager.Shutdown(ctx); err != nil {
		log.Println("unable to shutdown services:", err)
	}
}

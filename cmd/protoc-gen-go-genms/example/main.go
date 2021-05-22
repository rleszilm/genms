package main

import (
	"context"
	"log"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms/example/greeter"
	"github.com/rleszilm/genms/service"
	graphql_service "github.com/rleszilm/genms/service/graphql"
	grpc_service "github.com/rleszilm/genms/service/grpc"
	rest_service "github.com/rleszilm/genms/service/rest"
	"github.com/rleszilm/genms/service/rest/healthcheck"
)

type greets struct {
	greeter.UnimplementedWithRestAndGraphQLServer
}

func (g *greets) HelloRestAndGraphQL(_ context.Context, req *greeter.Message) (*greeter.Message, error) {
	return &greeter.Message{Value: req.GetValue() + " received"}, nil
}

func main() {
	ctx := context.Background()

	// create service  manager
	manager := service.NewManager()

	restServer, err := rest_service.NewServer("rest-api",
		&rest_service.Config{
			Config: service.Config{
				Transport: "tcp",
				Addr:      ":8081",
			},
		},
	)
	if err != nil {
		log.Fatalln("Unable to instantiate rest api server: ", err)
	}
	manager.Register(restServer)

	health := healthcheck.NewService(&healthcheck.Config{RequestPrefix: "/health"}, restServer)
	manager.Register(health)

	graphqlServer, err := graphql_service.NewServer("graphql-api",
		&graphql_service.Config{
			RestServer: restServer,
			Proxies: map[string]*graphql_service.GrpcProxy{
				"WithRestAndGraphQL": {
					Enabled:  true,
					Pattern:  "/graphql",
					Addr:     ":8080",
					Insecure: true,
				},
			},
		},
	)
	if err != nil {
		log.Fatalln("Unable to instantiate graphql api server: ", err)
	}
	manager.Register(graphqlServer)

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

	impl := &greets{}

	proxy := &grpc_service.Proxy{
		Enabled:  true,
		Pattern:  "/v1/grpc",
		Addr:     ":8080",
		Insecure: true,
	}

	service := greeter.NewWithRestAndGraphQLServerService(impl, grpcServer, restServer, graphqlServer, proxy)
	manager.Register(service)

	if err := manager.Initialize(ctx); err != nil {
		log.Fatal("Unable to start services: ", err)
	}
	manager.Wait()
	manager.Shutdown(ctx)
}

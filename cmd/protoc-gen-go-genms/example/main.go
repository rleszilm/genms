package main

import (
	"context"
	"log"

	"github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms/example/greeter"
	"github.com/rleszilm/gen_microservice/service"
	graphql_service "github.com/rleszilm/gen_microservice/service/graphql"
	grpc_service "github.com/rleszilm/gen_microservice/service/grpc"
	rest_service "github.com/rleszilm/gen_microservice/service/rest"
	"github.com/rleszilm/gen_microservice/service/rest/healthcheck"
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
			Proxies: map[string]*rest_service.ProxyGrpc{
				"WithRestAndGraphQL": {
					Enabled:  true,
					Pattern:  "/rest/",
					Addr:     ":8080",
					Insecure: true,
				},
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
			Proxies: map[string]*graphql_service.ProxyGrpc{
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

	service := greeter.NewWithRestAndGraphQLServerService(impl, grpcServer, restServer, graphqlServer)
	manager.Register(service)

	if err := manager.Initialize(ctx); err != nil {
		log.Fatal("Unable to start services: ", err)
	}
	manager.Wait()
	manager.Shutdown(ctx)
}

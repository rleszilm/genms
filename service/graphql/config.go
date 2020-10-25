package graphql_service

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rleszilm/gen_microservice/service"
	rest_service "github.com/rleszilm/gen_microservice/service/rest"
)

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	service.Config

	RestServer *rest_service.Server
	ProxyGrpc  ProxyGrpc `envconfig:"proxy"`
}

// ProxyGrpc is configuration used when configuring a rest service as a grpc proxy.
type ProxyGrpc struct {
	Enabled bool   `envconfig:"enabled" default:"true"`
	Pattern string `envconfig:"pattern" default:"/"`
	Addr    string `envconfig:"addr" default:""`
	Secure  bool   `envconfig:"secure" default:"false"`
}

// NewFromEnv generates a new set of configuration data.
func NewFromEnv(namespace string) (*Config, error) {
	c := &Config{}
	err := envconfig.Process(namespace, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

package grpc_service

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rleszilm/gen_microservice/service"
)

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	service.Config
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

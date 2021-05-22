package graphql_service

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rleszilm/genms/config"
	"github.com/rleszilm/genms/service"
	rest_service "github.com/rleszilm/genms/service/rest"
)

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	service.Config

	RestServer *rest_service.Server
	ProxyNames []string `envconfig:"proxies" default:""`
	Proxies    map[string]*GrpcProxy
}

// GrpcProxy is configuration used when configuring a rest service as a grpc proxy.
type GrpcProxy struct {
	Enabled  bool   `envconfig:"enabled" default:"true"`
	Pattern  string `envconfig:"pattern" default:"/"`
	Addr     string `envconfig:"addr" default:""`
	Insecure bool   `envconfig:"insecure" default:"false"`
}

// NewFromEnv generates a new set of configuration data.
func NewFromEnv(namespace string) (*Config, error) {
	c := &Config{
		Proxies: map[string]*GrpcProxy{},
	}

	err := envconfig.Process(namespace, c)
	if err != nil {
		return nil, err
	}

	for _, proxyName := range c.ProxyNames {
		proxy := &GrpcProxy{}
		err := envconfig.Process(config.Join(namespace, proxyName), proxy)
		if err != nil {
			return nil, err
		}

		c.Proxies[proxyName] = proxy
	}

	return c, nil
}

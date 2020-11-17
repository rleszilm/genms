package rest_service

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rleszilm/gen_microservice/config"
	"github.com/rleszilm/gen_microservice/service"
)

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	service.Config

	ProxyNames []string `envconfig:"proxies" default:""`
	Proxies    map[string]*ProxyGrpc
}

// ProxyGrpc is configuration used when configuring a rest service as a grpc proxy.
type ProxyGrpc struct {
	Enabled  bool   `envconfig:"enabled" default:"true"`
	Pattern  string `envconfig:"pattern" default:"/"`
	Addr     string `envconfig:"addr" default:""`
	Insecure bool   `envconfig:"secure" default:"true"`
}

// NewFromEnv generates a new set of configuration data.
func NewFromEnv(namespace string) (*Config, error) {
	c := &Config{
		Proxies: map[string]*ProxyGrpc{},
	}

	err := envconfig.Process(namespace, c)
	if err != nil {
		return nil, err
	}

	for _, proxyName := range c.ProxyNames {
		proxy := &ProxyGrpc{}
		err := envconfig.Process(config.Join(namespace, proxyName), proxy)
		if err != nil {
			return nil, err
		}

		c.Proxies[proxyName] = proxy
	}

	return c, nil
}

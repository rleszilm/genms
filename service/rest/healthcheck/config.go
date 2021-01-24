package healthcheck

import (
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	Name          string           `envconfig:"name" default:""`
	RequestPrefix string           `envconfig:"request_path" default:"/health"`
	HealthyFunc   http.HandlerFunc `ignored:"true"`
	ReadyFunc     http.HandlerFunc `ignored:"true"`
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

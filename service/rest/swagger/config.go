package swagger

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	Name          string `envconfig:"name" default:""`
	Enabled       bool   `envconfig:"enabled"`
	RequestPrefix string `envconfig:"request_prefix" default:"/swagger/"`
	Dir           string `envconfig:"dir" default:"var/swagger"`
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

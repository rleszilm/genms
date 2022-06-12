package pprof

import "github.com/kelseyhightower/envconfig"

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	Enabled bool   `envconfig:"enabled" default:"false"`
	Name    string `envconfig:"name" default:""`
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

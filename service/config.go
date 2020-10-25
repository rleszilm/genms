package service

import "github.com/kelseyhightower/envconfig"

// Config is default config that can be used by any service type. This is meant
// to be embedded into specific config types rather than directly instantiated.
type Config struct {
	Transport string `envconfig:"transport" default:"tcp"`
	Addr      string `envconfig:"addr" default:":8080"`
	TLS       TLS
}

// TLS is configuration for use with tls connections
type TLS struct {
	Enabled bool   `envconfig:"enabled" default:"false"`
	Key     string `envconfig:"key" default:""`
	Cert    string `envconfig:"cert" default:""`
}

// NewFromEnv returns a new Gamehub configuration based on environment variables.
func NewFromEnv(prefix string) (*Config, error) {
	c := Config{}
	if err := envconfig.Process(prefix, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

package prometheus_exporter

import "github.com/kelseyhightower/envconfig"

// Config defines the configuration for a prometheus exporter.
type Config struct {
	RequestPath string            `envconfig:"request_path" default:"/metrics"`
	Namespace   string            `envconfig:"namespace" default:"gameday"`
	Labels      map[string]string `envconfig:"labels" default:""`
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

package mongo

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	URI             string        `envconfig:"uri" default:"mongodb://localhost:27017"`
	AppName         string        `envconfig:"appname" default:""`
	MaxPoolSize     uint64        `envconfig:"pool_size" default:"10"`
	MaxConnIdleTime time.Duration `envconfig:"conn_idle_time" default:"30s"`
	Database        string        `envconfig:"database" required:"true"`
	Timeout         time.Duration `envconfig:"timeout" default:"5s"`
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

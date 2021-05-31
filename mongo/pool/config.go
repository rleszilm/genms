package pool

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config defines the configuration for dal mongo interactions.
type Config struct {
	URI             string        `envconfig:"uri" default:"mongodb://localhost:27017"`
	AppName         string        `envconfig:"appname" default:""`
	MaxPoolSize     uint64        `envconfig:"pool_size" default:"25"`
	MaxConnIdleTime time.Duration `envconfig:"conn_idle_time" default:"30s"`
	Database        string        `envconfig:"database" default:"vvv-repl"`
	Timeout         time.Duration `envconfig:"timeout" default:"5s"`
	ReadPref        string        `envconfig:"read_pref" default:"primarypreferred"`
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

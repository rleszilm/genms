package kafka

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config is the struct used to parse configuration from environment variables.
type Config struct {
	Consumer ConsumerConfig `envconfig:"consumer"`
	Producer ProducerConfig `envconfig:"producer"`
}

// ConsumerConfig is the struct used to parse configuration from environment variables.
type ConsumerConfig struct {
	BrokerList []string `envconfig:"broker_list" default:"kafka:9092"`
	Group      string   `envconfig:"group" default:"gameday"`
}

// ProducerConfig is the struct used to parse configuration from environment variables.
type ProducerConfig struct {
	BrokerList     []string      `envconfig:"broker_list" default:"kafka:2181"`
	FlushFrequency time.Duration `envconfig:"flush_frequency" default:"10ms"`
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

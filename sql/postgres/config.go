package postgres_sql

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/rleszilm/gen_microservice/sql"
)

// Config implements sql.Config
type Config struct {
	sql.EnvConfig
}

// Driver implements sql.Config.SqlDriver
func (c *Config) Driver() string {
	return "postgres"
}

// ConnectionString implements sql.Config.ConnectionString
func (c *Config) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Database,
	)
}

// NewFromEnv returns a new Config from env vars.
func NewFromEnv(namespace string) (*Config, error) {
	c := &Config{}
	err := envconfig.Process(namespace, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

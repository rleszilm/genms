package sql

// Config is an interface that provides configuration values.
type Config interface {
	Driver() string
	ConnectionString() string
}

// EnvConfig represents sql connection data
type EnvConfig struct {
	User     string `envconfig:"user" required:"true"`
	Password string `envconfig:"password" required:"true"`
	Host     string `envconfig:"host" default:"localhost"`
	Port     int    `envconfig:"port" required:"true"`
	Database string `envconfig:"database" required:"true"`
}

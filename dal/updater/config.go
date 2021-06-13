package updater

import (
	"time"
)

type Config struct {
	Database   string        `envconfig:"database" required:"true"`
	Collection string        `envconfig:"collection" required:"true"`
	Interval   time.Duration `envconfig:"interval" default:"15m"`
	Timeout    time.Duration `envconfig:"timeout" default:"5s"`
}

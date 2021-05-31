package pool

import (
	"context"

	pkgMongo "github.com/rleszilm/genms/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Dialer is a structure that can be used to get a mongo connection.
type Dialer struct {
	config   *Config
	client   *mongo.Client
	options  *options.ClientOptions
	readpref readpref.Mode
}

// Initialize implements service.Initialize
func (d *Dialer) Initialize(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, d.config.Timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, d.options)
	if err != nil {
		return err
	}

	rp, err := readpref.New(d.readpref)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, rp); err != nil {
		return err
	}
	d.client = client

	return nil
}

// Shutdown implements service.Shutdown
func (d *Dialer) Shutdown(ctx context.Context) error {
	return d.client.Disconnect(ctx)
}

// Dial returns a new mongo connection.
func (d *Dialer) Dial(_ context.Context) (pkgMongo.Client, error) {
	session, err := d.client.StartSession()
	if err != nil {
		return nil, err
	}

	return pkgMongo.NewSimpleClient(session), nil
}

// NewDialer returns a new Dialer.
func NewDialer(config *Config) (*Dialer, error) {
	readpref, err := readpref.ModeFromString(config.ReadPref)
	if err != nil {
		return nil, err
	}

	opts := options.Client()
	opts.SetAppName(config.AppName)
	opts.SetMaxPoolSize(config.MaxPoolSize)
	opts.SetMaxConnIdleTime(config.MaxConnIdleTime)
	opts.ApplyURI(config.URI)

	return &Dialer{
		config:   config,
		options:  opts,
		readpref: readpref,
	}, nil
}

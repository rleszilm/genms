package mongo

import (
	"context"

	"github.com/rleszilm/gen_microservice/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client is a wrapper for a mongo.Client that allows it to be managed as a service.
type Client struct {
	service.Deps
	*mongo.Client

	config *Config
}

// Initialize implements the service.Initialize interface for Service.
func (c *Client) Initialize(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	// ensure mongo is available
	rp, err := readpref.New(readpref.PrimaryPreferredMode)
	if err != nil {
		return err
	}
	return c.Client.Ping(ctx, rp)
}

// Shutdown implements the service.Shutdown interface for Service.
func (c *Client) Shutdown(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}

// NameOf implements service.NameOf
func (c *Client) NameOf() string {
	return "mongo"
}

// String implements service.String
func (c *Client) String() string {
	return c.NameOf()
}

// NewClient returns a new Client
func NewClient(config *Config) (*Client, error) {
	opts := options.Client()
	opts.SetAppName(config.AppName)
	opts.SetMaxPoolSize(config.MaxPoolSize)
	opts.SetMaxConnIdleTime(config.MaxConnIdleTime)
	opts.ApplyURI(config.URI)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
		config: config,
	}, nil
}

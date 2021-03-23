package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rleszilm/genms/service"
)

// Client is a wrapper for a redis.Client that allows it to be manages as a service.
type Client struct {
	service.Dependencies
	*redis.Client
}

// Initialize implements the service.Initialize interface for Service.
func (c *Client) Initialize(ctx context.Context) error {
	// ensure redis is available
	_, err := c.Client.Ping(ctx).Result()
	return err
}

// Shutdown implements the service.Shutdown interface for Service.
func (c *Client) Shutdown(_ context.Context) error {
	return c.Client.Close()
}

// NameOf implements service.NameOf
func (c *Client) NameOf() string {
	return "redis"
}

// String implements service.String
func (c *Client) String() string {
	return c.NameOf()
}

// NewClient returns a new Client.
func NewClient(config *Config) *Client {
	return &Client{
		Client: redis.NewClient(ToV8Options(config)),
	}
}

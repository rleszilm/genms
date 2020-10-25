package kafka

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/rleszilm/gen_microservice/service"
)

// ConsumerFunc is a handler for a consumer claim.
type ConsumerFunc func(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim) error

// Consumer is an interface that describes a Kafka consumer.
type Consumer interface {
	service.Service

	Consume(context.Context) error
}

// ConsumerGroup is an alias for sarama.ConsumerGroup
type ConsumerGroup sarama.ConsumerGroup

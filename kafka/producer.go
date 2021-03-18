package kafka

import (
	"context"

	"github.com/rleszilm/genms/service"
)

// Producer is an interface that describes a Kafka producer.
type Producer interface {
	service.Service

	// Producer writes a message to the given topic.
	Produce(context.Context, *ProducerMessage) error
}

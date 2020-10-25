package kafka

import (
	"context"

	"github.com/rleszilm/gen_microservice/service"
)

// Producer is an interface that describes a Kafka producer.
type Producer interface {
	service.Service

	// Producer writes a message to the given topic.
	Produce(context.Context, *ProducerMessage) error
}

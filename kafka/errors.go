package kafka

import "errors"

var (
	// ErrProducerShutdown is returned when producing to a closed producer.
	ErrProducerShutdown = errors.New("kafka: producer is not active")
	// ErrConsumerNoHandler is returned when there is no consumer for a claim.
	ErrConsumerNoHandler = errors.New("kafka: no consumer for topic")
)

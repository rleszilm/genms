package kafka

import "github.com/Shopify/sarama"

// ProducerMessage is a wrapper for the underlying sarama producer message.
type ProducerMessage struct {
	*sarama.ProducerMessage
}

// ProducerMetadata contains metadata about the ProducerMessage
type ProducerMetadata struct {
	OnSuccess func(*ProducerMessage)
	OnError   func(*ProducerMessage, error)
}

// NewProducerMessage returns a new ProducerMessage
func NewProducerMessage(topic string, key sarama.Encoder, value sarama.Encoder) *ProducerMessage {
	return &ProducerMessage{
		ProducerMessage: &sarama.ProducerMessage{
			Topic:    topic,
			Key:      key,
			Value:    value,
			Metadata: &ProducerMetadata{},
		},
	}
}

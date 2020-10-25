package kafka

import (
	"context"

	"github.com/Shopify/sarama"
)

// RegistryConsumer is a consumer that has handlers registered per topic.
type RegistryConsumer struct {
	ConsumerGroup

	done     chan struct{}
	handlers map[string]ConsumerFunc
}

// Initialize implements service.Service.Initialize
func (r *RegistryConsumer) Initialize(ctx context.Context) error {
	return nil
}

// Shutdown implements service.Service.Shutdown
func (r *RegistryConsumer) Shutdown(ctx context.Context) error {
	close(r.done)
	return nil
}

// Consume implements Consumer.Consume
func (r *RegistryConsumer) Consume(ctx context.Context) error {
	topics := []string{}
	for topic := range r.handlers {
		topics = append(topics, topic)
	}

	for {
		select {
		case <-r.done:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := r.ConsumerGroup.Consume(ctx, topics, r); err != nil {
				return err
			}
		}
	}
}

// Setup implements sarama.ConsumerGroupHandler.Setup
func (r *RegistryConsumer) Setup(sess sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup implements sarama.ConsumerGroupHandler.Cleanup
func (r *RegistryConsumer) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim implements sarama.ConsumerGroupHandler.ConsumeClaim
func (r *RegistryConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	if handler, ok := r.handlers[claim.Topic()]; ok {
		return handler(sess, claim)
	}
	return ErrConsumerNoHandler
}

// NewRegistryConsumer returns a new RegistryConsumer
func NewRegistryConsumer(cfg *ConsumerConfig) (*RegistryConsumer, error) {
	scfg := sarama.NewConfig()
	scfg.Version = sarama.V2_6_0_0

	group, err := sarama.NewConsumerGroup(cfg.BrokerList, cfg.Group, scfg)
	if err != nil {
		return nil, err
	}

	return &RegistryConsumer{
		ConsumerGroup: group,

		done:     make(chan struct{}),
		handlers: map[string]ConsumerFunc{},
	}, nil
}

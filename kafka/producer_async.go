package kafka

import (
	"context"

	"github.com/Shopify/sarama"
)

// AsyncProducer is an asynchronous kafka producer.
type AsyncProducer struct {
	done     chan struct{}
	producer sarama.AsyncProducer
}

// Initialize implements service.Service.Initialize
func (a *AsyncProducer) Initialize(ctx context.Context) error {
	go a.worker(ctx)
	return nil
}

// Shutdown implements service.Service.Shutdown
func (a *AsyncProducer) Shutdown(ctx context.Context) error {
	close(a.done)
	return nil
}

// Produce implements Producer.Produce
func (a *AsyncProducer) Produce(ctx context.Context, m *ProducerMessage) error {
	select {
	case <-a.done:
		return ErrProducerShutdown
	case <-ctx.Done():
		return ctx.Err()
	default:
		a.producer.Input() <- m.ProducerMessage
	}

	return nil
}

func (a *AsyncProducer) worker(ctx context.Context) {
	for {
		select {
		case <-a.done:
			return
		case <-ctx.Done():
			return
		case msg := <-a.producer.Successes():
			if md, ok := msg.Metadata.(*ProducerMetadata); ok {
				if md.OnSuccess != nil {
					md.OnSuccess(&ProducerMessage{ProducerMessage: msg})
				}
			}
		case msg := <-a.producer.Errors():
			if md, ok := msg.Msg.Metadata.(*ProducerMetadata); ok {
				if md.OnError != nil {
					md.OnError(&ProducerMessage{ProducerMessage: msg.Msg}, msg.Err)
				}
			}
		}
	}
}

// NewAsyncProducer returns a new AsyncProducer
func NewAsyncProducer(config *ProducerConfig) (*AsyncProducer, error) {
	scfg := sarama.NewConfig()
	scfg.Producer.RequiredAcks = sarama.WaitForLocal
	scfg.Producer.Compression = sarama.CompressionSnappy
	scfg.Producer.Flush.Frequency = config.FlushFrequency
	scfg.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer(config.BrokerList, scfg)
	if err != nil {
		return nil, err
	}

	return &AsyncProducer{
		done:     make(chan struct{}),
		producer: producer,
	}, nil
}

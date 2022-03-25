package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/satioO/scheduler/scheduler/cqrs/message"
	"github.com/sirupsen/logrus"
)

type PublisherConfig struct {
	// Kafka brokers list.
	Brokers []string

	// Marshaler is used to marshal messages from Watermill format into Kafka format.
	Marshaler Marshaler

	// OverwriteSaramaConfig holds additional sarama settings.
	OverwriteSaramaConfig *sarama.Config
}

type Publisher struct {
	config   PublisherConfig
	producer sarama.SyncProducer

	closed bool
}

func (c *PublisherConfig) setDefaults() {
	if c.OverwriteSaramaConfig == nil {
		c.OverwriteSaramaConfig = DefaultSaramaSyncPublisherConfig()
	}
}

func DefaultSaramaSyncPublisherConfig() *sarama.Config {
	config := sarama.NewConfig()

	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Version = sarama.V1_0_0_0
	config.Metadata.Retry.Backoff = time.Second * 2
	config.ClientID = "watermill"

	return config
}

func NewPublisher(config PublisherConfig) (*Publisher, error) {
	producer, err := sarama.NewSyncProducer(config.Brokers, config.OverwriteSaramaConfig)

	if err != nil {
		return nil, errors.Wrap(err, "cannot create Kafka producer")
	}

	return &Publisher{
		config:   config,
		producer: producer,
	}, nil
}

func (p *Publisher) Publish(topic string, messages ...*message.Message) error {
	for _, msg := range messages {
		logrus.Printf("Topic:::%s, Message: %v", topic, msg)

		kafkaMsg, err := p.config.Marshaler.Marshal(topic, msg)
		if err != nil {
			return errors.Wrapf(err, "cannot marshal message %s", msg.UUID)
		}

		logrus.Info("Kafka Message:::", kafkaMsg)

		partition, offset, err := p.producer.SendMessage(kafkaMsg)
		if err != nil {
			return errors.Wrapf(err, "cannot produce message %s", msg.UUID)
		}

		logrus.Info("Message sent to Kafka", map[string]any{
			"partition": partition,
			"offset":    offset,
		})
	}

	return nil
}

// Close should flush unsent messages, if publisher is async.
func (*Publisher) Close() error {
	return nil
}

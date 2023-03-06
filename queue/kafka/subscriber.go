package kafka

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/depocket/support/queue"
	"go.uber.org/zap"
)

func DepocketSaramaSubscriberConfig() *sarama.Config {
	config := sarama.NewConfig()

	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true
	config.ClientID = "depocket"

	return config
}

func NewSubscriber(z *zap.Logger, consumerGroup string) queue.Subscriber {
	config := parseConfig()
	brokers := strings.Split(config.Servers, ",")

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: DepocketSaramaSubscriberConfig(),
			ConsumerGroup:         consumerGroup,
		},
		queue.NewLogger(z),
	)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

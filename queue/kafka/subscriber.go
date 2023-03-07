package kafka

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/depocket/support/queue"
	"go.uber.org/zap"
)

func DepocketSaramaSubscriberConfig() *sarama.Config {
	config := sarama.NewConfig()

	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true
	config.ClientID = "depocket"
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}

func NewSubscriber(z *zap.Logger, consumerGroup string) queue.Subscriber {
	config := parseConfig()
	brokers := strings.Split(config.Servers, ",")

	var logger watermill.LoggerAdapter = nil
	if z != nil {
		logger = queue.NewLogger(z)
	}
	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: DepocketSaramaSubscriberConfig(),
			ConsumerGroup:         consumerGroup,
		},
		logger,
	)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

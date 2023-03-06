package kafka

import (
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/depocket/support/queue"
	"go.uber.org/zap"
)

func DepocketSaramaPublisherConfig() *sarama.Config {
	config := sarama.NewConfig()

	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Version = sarama.V1_0_0_0
	config.Metadata.Retry.Backoff = time.Second * 2
	config.ClientID = "depocket"

	return config
}

func NewPublisher(z *zap.Logger) queue.Publisher {
	config := parseConfig()
	brokers := strings.Split(config.Servers, ",")

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:               brokers,
			Marshaler:             kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: DepocketSaramaPublisherConfig(),
		},
		queue.NewLogger(z),
	)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

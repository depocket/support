package kafka

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Servers string `env:"KAFKA_SERVERS" envDefault:"localhost:9092"`
}

func parseConfig() *Config {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}

package factory

import (
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	wmkafka "github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/cikupin/learn-firehose/config"
)

func NewKafkaPublisher() *wmkafka.Publisher {
	kafka0Host := fmt.Sprintf("%s:%d", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)
	pubConfig := wmkafka.PublisherConfig{
		Brokers:               []string{kafka0Host},
		Marshaler:             wmkafka.DefaultMarshaler{},
		OverwriteSaramaConfig: wmkafka.DefaultSaramaSyncPublisherConfig(),
	}

	publisher, err := wmkafka.NewPublisher(pubConfig, watermill.NewStdLogger(true, false))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return publisher
}

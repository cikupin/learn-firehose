package config

var Cfg config

type kafkaConfig struct {
	Host string `env:"KAFKA_HOST"`
	Port int    `env:"KAFKA_PORT"`
}

type config struct {
	Kafka kafkaConfig
}

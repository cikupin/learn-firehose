package config

var Cfg config

type appConfig struct {
	Port int `env:"APP_PORT" envDefault:"3000"`
}

type kafkaConfig struct {
	Host      string `env:"KAFKA_HOST"`
	Port      int    `env:"KAFKA_PORT"`
	TopicGRPC string `env:"KAFKA_TOPIC_GRPC"`
	TopicJSON string `env:"KAFKA_TOPIC_JSON"`
}

type config struct {
	App   appConfig
	Kafka kafkaConfig
}

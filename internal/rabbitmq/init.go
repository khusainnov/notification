package rabbitmq

import "github.com/khusainnov/notification/internal/config"

func prepareConsumer(cfg *config.Config) *Config {
	return &Config{
		RabbitHost:     cfg.RabbitHost,
		RabbitPort:     cfg.RabbitPort,
		RabbitUser:     cfg.RabbitUser,
		RabbitPassword: cfg.RabbitPassword,
		QueueName:      cfg.QueueName,
	}
}

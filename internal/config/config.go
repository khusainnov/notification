package config

import "go.uber.org/zap"

type Config struct {
	L *zap.Logger

	GRPCAddr string `env:"GRPC_ADDR" envDefault:":9000"`

	// Email credentials
	MailFrom     string `env:"MAIL_FROM,required"`
	MailPassword string `env:"MAIL_PASSWORD,required"`
	MailHost     string `env:"MAIL_HOST" envDefault:"smtp.gmail.com"`
	MailPort     string `env:"MAIL_PORT" envDefault:":587"`

	// RabbitMQ
	RabbitHost     string `env:"RABBIT_HOST" envDefault:"localhost"`
	RabbitPort     string `env:"RABBIT_PORT" envDefault:":15672"`
	RabbitUser     string `env:"RABBIT_USER" envDefault:"rabbitmq"`
	RabbitPassword string `env:"RABBIT_PASSWORD" envDefault:"rabbitmq"`

	QueueName string `env:"RABBIT_QUEUE_NAME" envDefault:"rsskhsnnv"`
}

package email

import "github.com/khusainnov/notification/internal/config"

func prepareConfig(cfg *config.Config) *Config {
	return &Config{
		MailFrom:     cfg.MailFrom,
		MailPassword: cfg.MailPassword,
		MailHost:     cfg.MailHost,
		MailPort:     cfg.MailPort,
	}
}

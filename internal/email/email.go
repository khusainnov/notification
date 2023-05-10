package email

import (
	"fmt"
	"net/smtp"

	"github.com/khusainnov/notification/internal/config"
	"github.com/khusainnov/notification/internal/model"
	"go.uber.org/zap"
)

// Client –
type Client struct {
	log *zap.Logger
	cfg *Config
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		log: cfg.L,
		cfg: prepareConfig(cfg),
	}
}

// SendEmail – send final emails
func (c *Client) SendEmail(email *model.SendEmail) error {
	auth := c.authMail()

	addr := c.cfg.MailHost + c.cfg.MailPort

	msg := fmt.Sprintf("Subject: %s\r\n\r\n%s", email.Subject, email.Email.Body)

	if err := smtp.SendMail(addr, auth, c.cfg.MailFrom, email.Emails, []byte(msg)); err != nil {
		return fmt.Errorf("failed to send email, %w", err)
	}

	return nil
}

func (c *Client) authMail() smtp.Auth {
	auth := smtp.PlainAuth(
		"",
		c.cfg.MailFrom,
		c.cfg.MailPassword,
		c.cfg.MailHost,
	)

	return auth
}

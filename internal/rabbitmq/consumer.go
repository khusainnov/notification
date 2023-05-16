package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/khusainnov/notification/internal/config"
	"github.com/khusainnov/notification/internal/email"
	"github.com/khusainnov/notification/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

// Client –
type Client struct {
	log   *zap.Logger
	conn  *amqp.Connection
	cfg   *Config
	email *email.Client
}

func NewConsumer(cfg *config.Config, email *email.Client) (*Client, error) {
	addr := fmt.Sprintf("amqp://%s:%s@%s%s/", cfg.RabbitUser, cfg.RabbitPassword, cfg.RabbitHost, cfg.RabbitPort)
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to rabbit, %w", err)
	}

	return &Client{
		log:   cfg.L,
		conn:  conn,
		cfg:   prepareConsumer(cfg),
		email: email,
	}, nil
}

// Run – listen incoming messages and send it to `SendEmail`
func (c *Client) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}

	defer c.conn.Close()
	messages, err := c.setup()
	if err != nil {
		c.log.Fatal("failed to get messages channel", zap.Error(err))
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				c.log.Info("called context done")
				return
			case msg, ok := <-messages:
				if !ok {
					c.log.Fatal("reading from a closed channel")
				}

				var req model.SendEmail

				if err = json.Unmarshal(msg.Body, &req); err != nil {
					c.log.Error("failed to unmarshal body", zap.Error(err))
				}

				if err = c.email.SendEmail(&req); err != nil {
					c.log.Error("failed to send email", zap.Error(err))
				}
			default:
				continue
			}
		}
	}()

	wg.Wait()
}

// setup – setting up rabbit delivery channel
func (c *Client) setup() (<-chan amqp.Delivery, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		c.log.Error("cannot get message channel", zap.Error(err))
		return nil, fmt.Errorf("cannot get message channel, %w", err)
	}

	q, err := ch.QueueDeclare(
		c.cfg.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.log.Error("failed declare queue", zap.Error(err))
		return nil, fmt.Errorf("failed declare queue, %w", err)
	}

	messages, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.log.Error("failed to register a consumer", zap.Error(err))
		return nil, fmt.Errorf("failed to register a consumer, %w", err)
	}

	return messages, nil
}

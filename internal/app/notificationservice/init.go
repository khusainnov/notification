package notificationservice

import (
	"github.com/khusainnov/notification/internal/config"
	"github.com/khusainnov/notification/internal/email"
	napi "github.com/khusainnov/notification/pkg/notificationapi/v1"
	"go.uber.org/zap"
)

type NotificationImpl struct {
	*napi.UnimplementedNotificationServiceServer
	l     *zap.Logger
	email *email.Client
}

func NewClient(cfg *config.Config) *NotificationImpl {
	return &NotificationImpl{
		l:     cfg.L,
		email: email.NewClient(cfg),
	}
}

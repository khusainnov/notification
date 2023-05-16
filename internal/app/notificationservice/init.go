package notificationservice

import (
	"github.com/khusainnov/notification/internal/email"
	napi "github.com/khusainnov/notification/pkg/notificationapi/v1"
)

type NotificationImpl struct {
	*napi.UnimplementedNotificationServiceServer
	email *email.Client
}

func NewClient(email *email.Client) *NotificationImpl {
	return &NotificationImpl{
		email: email,
	}
}

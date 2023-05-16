package notificationservice

import (
	"context"

	"github.com/khusainnov/notification/internal/app/notificationservice/adapters"
	napi "github.com/khusainnov/notification/pkg/notificationapi/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *NotificationImpl) SendEmail(ctx context.Context, req *napi.SendEmailRequest) (*napi.SendEmailResponse, error) {
	napiReq := adapters.SendEmailFromPb(req)
	if err := n.email.SendEmail(napiReq); err != nil {
		return &napi.SendEmailResponse{
			Status: napi.EmailDeliveryStatus_EMAIL_DELIVERY_STATUS_FAILED,
		}, status.Errorf(codes.Internal, "%v, %v", napi.EmailDeliveryStatus_EMAIL_DELIVERY_STATUS_FAILED, err)
	}

	return &napi.SendEmailResponse{
		Status: napi.EmailDeliveryStatus_EMAIL_DELIVERY_STATUS_OK,
	}, nil
}

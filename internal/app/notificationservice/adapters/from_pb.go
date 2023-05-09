package adapters

import (
	"net/mail"

	"github.com/khusainnov/notification/internal/model"
	napi "github.com/khusainnov/notification/pkg/notificationapi/v1"
)

func SendEmailFromPb(req *napi.SendEmailRequest) *model.SendEmail {
	r := &model.SendEmail{
		Emails:  validateEmail(req.Emails...),
		Subject: req.Subject,
		Email: struct {
			Body string `json:"body"`
		}{
			Body: req.Body.Body,
		},
	}

	return r
}

func validateEmail(emails ...string) []string {
	wrongMails := 0
	for i, email := range emails {
		if _, err := mail.ParseAddress(email); err != nil {
			wrongMails++
			copy(emails[i:], emails[i+1:])
		}
	}

	return emails[:len(emails)-wrongMails]
}

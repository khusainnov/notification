package model

type SendEmail struct {
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Email   struct {
		Body string `json:"body"`
	} `json:"email"`
}

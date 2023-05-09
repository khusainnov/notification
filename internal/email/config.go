package email

type Config struct {
	MailFrom     string `env:"MAIL_FROM,required"`
	MailPassword string `env:"MAIL_PASSWORD,required"`
	MailHost     string `env:"MAIL_HOST" envDefault:"smtp.gmail.com"`
	MailPort     string `env:"MAIL_PORT" envDefault:":587"`
}

package smtp

import "github.com/wolfsblu/go-chef/infra/env"

func NewSMTPMailer() *Mailer {
	return &Mailer{
		config: Config{
			Host:     env.MustGet("SMTP_HOST"),
			Port:     env.MustGetInt("SMTP_PORT"),
			User:     env.MustGet("SMTP_USERNAME"),
			Password: env.MustGet("SMTP_PASSWORD"),
		},
	}
}

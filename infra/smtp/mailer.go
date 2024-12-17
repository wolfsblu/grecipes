package smtp

import (
	"bytes"
	"fmt"
	"github.com/wolfsblu/go-chef/domain"
	"gopkg.in/gomail.v2"
	"html/template"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Mailer struct {
	config Config
}

func (s *Mailer) SendPasswordReset(token domain.PasswordResetToken) error {
	tpl, err := buildTemplate("password-reset.html", PasswordResetTemplate{
		ResetLink: "https://google.com",
	})
	if err != nil {
		return err
	}
	if err = s.sendMessage(s.config.User, token.User.Email, "Password Reset", tpl); err != nil {
		return err
	}
	return nil
}

func (s *Mailer) sendMessage(sender, recipient, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", recipient)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(s.config.Host, s.config.Port, s.config.User, s.config.Password)
	return dialer.DialAndSend(msg)
}

func buildTemplate(path string, data any) (string, error) {
	t := template.New(path)
	t, err := t.ParseFS(templateFS, fmt.Sprintf("templates/%s", path))
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

package smtp

import (
	"bytes"
	"github.com/wolfsblu/go-chef/domain"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/fs"
	"log"
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

type PasswordResetTemplate struct {
	ResetLink string
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
	subFS, err := fs.Sub(templateFS, "templates")
	items, _ := fs.ReadDir(subFS, ".")
	for _, item := range items {
		log.Println(item)
	}
	if err != nil {
		return "", err
	}
	t, err = t.ParseFS(subFS, path)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

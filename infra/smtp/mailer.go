package smtp

import (
	"github.com/wolfsblu/go-chef/domain"
	"gopkg.in/gomail.v2"
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
func (s *Mailer) SendUserRegistration(registration domain.UserRegistration) error {
	tpl, err := buildTemplate("user-registration.html", UserRegistrationTemplate{
		ConfirmLink: "https://google.com",
	})
	if err != nil {
		return err
	}
	if err = s.sendMessage(s.config.User, registration.User.Email, "Registration", tpl); err != nil {
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

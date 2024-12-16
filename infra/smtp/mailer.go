package smtp

import "github.com/wolfsblu/go-chef/domain"

type Mailer struct{}

func (s *Mailer) SendPasswordReset(u domain.User) error {
	return nil
}

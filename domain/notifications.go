package domain

type NotificationSender interface {
	SendPasswordReset(token PasswordResetToken) error
	SendUserRegistration(registration UserRegistration) error
}

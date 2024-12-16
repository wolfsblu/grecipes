package domain

type NotificationSender interface {
	SendPasswordReset(u User) error
}

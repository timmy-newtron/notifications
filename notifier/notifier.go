package notifier

import (
	"stackit/notifications/postnotification"
)

type NotificationResponse struct {
	StatusCode int
	Status     string
}

type Notifier interface {
	Notify(*postnotification.PostNotification) (*NotificationResponse, error)
}

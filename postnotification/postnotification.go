package postnotification

import "errors"

type Type string

// string enum for notification types
const (
	Warning Type = "Warning"
	Info    Type = "Info"
)

type PostNotification struct {
	Type        Type   `json:"Type" binding:"required" enum:"Warning,Info"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Channel     string `json:"Channel"`
}

// Validate validates the post notification
func (postNotification *PostNotification) Validate() error {
	if postNotification.Type == "" {
		return errors.New("Type is required")
	}
	if postNotification.Name == "" {
		return errors.New("Name is required")
	}
	if postNotification.Description == "" {
		return errors.New("Description is required")
	}
	if postNotification.Channel == "" {
		return errors.New("Channel is required")
	}
	return nil
}

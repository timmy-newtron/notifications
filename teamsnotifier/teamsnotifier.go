package teamsnotifier

import (
	"bytes"
	"encoding/json"
	"net/http"
	"stackit/notifications/notifier"
	"stackit/notifications/postnotification"
)

type TeamsNotifier struct {
	WebhookUrl string
}

// Notify sends a notification to teams
func (teamsNotifier *TeamsNotifier) Notify(postNotification *postnotification.PostNotification) (*notifier.NotificationResponse, error) {
	return teamsNotifier.sendToTeamsDependingOnType(postNotification)
}

const adaptiveCardSchema = "http://adaptivecards.io/schemas/adaptive-card.json"
const adaptiveCardVersion = "1.2"

// A Teams message looks like this:
type TeamsMessage struct {
	Type        string       `json:"type"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	ContentType string  `json:"contentType"`
	ContentUrl  string  `json:"contentUrl"`
	Content     Content `json:"content"`
}

type Content struct {
	Schema  string      `json:"$schema"`
	Type    string      `json:"type"`
	Version string      `json:"version"`
	Body    []TextBlock `json:"body"`
}

type TextBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// createTeamsMessage creates a teams message from a string
func createTeamsMessage(message string) TeamsMessage {
	teamsMessage := TeamsMessage{
		Type: "message",
		Attachments: []Attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				ContentUrl:  "",
				Content: Content{
					Schema:  adaptiveCardSchema,
					Type:    "AdaptiveCard",
					Version: adaptiveCardVersion,
					Body: []TextBlock{
						{
							Type: "TextBlock",
							Text: message,
						},
					},
				},
			},
		},
	}
	return teamsMessage
}

// sendToTeamsDependingOnType sends a notification to teams depending on the type
func (teamsNotifier *TeamsNotifier) sendToTeamsDependingOnType(notification *postnotification.PostNotification) (*notifier.NotificationResponse, error) {
	switch notification.Type {
	case postnotification.Warning:
		return teamsNotifier.sendToTeams(notification)
	default:
		return &notifier.NotificationResponse{
			Status:     "Message not dispatched",
			StatusCode: 200,
		}, nil
	}
}

// sendToTeams sends a notification to teams
func (teamsNotifier *TeamsNotifier) sendToTeams(notification *postnotification.PostNotification) (*notifier.NotificationResponse, error) {

	teamsMessage := createTeamsMessage(notification.Description)
	jsonData, err := json.Marshal(teamsMessage)
	if err != nil {
		return &notifier.NotificationResponse{
			Status:     "error",
			StatusCode: 500,
		}, err
	}

	res, err := http.Post(teamsNotifier.WebhookUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return &notifier.NotificationResponse{
			Status:     "error",
			StatusCode: 500,
		}, err
	}

	return &notifier.NotificationResponse{
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}, nil
}

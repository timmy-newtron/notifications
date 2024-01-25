package teamsnotifier_test

import (
	"net/http"
	"stackit/notifications/postnotification"
	"stackit/notifications/teamsnotifier"
	"testing"
)

var warningPayload = postnotification.PostNotification{
	Type:        "Warning",
	Name:        "Test Warning",
	Description: "This is a test warning",
}

var infoPayload = postnotification.PostNotification{
	Type:        "Info",
	Name:        "Test Info",
	Description: "This is a test info",
}

var garbagePayload = postnotification.PostNotification{
	Type:        "Garbage",
	Name:        "Test Garbage",
	Description: "This is a test garbage",
}

var testCases = []struct {
	payload postnotification.PostNotification
	code    int
	body    string
}{
	{warningPayload, http.StatusOK, "200 OK"},
	{infoPayload, http.StatusOK, "Message not dispatched"},
	{garbagePayload, http.StatusOK, "Message not dispatched"},
}

func TestPostNotifications(t *testing.T) {

	teamsNotifier := &teamsnotifier.TeamsNotifier{
		WebhookUrl: "YourSuperSecretWebhookUrlHere",
	}

	for _, tc := range testCases {
		res, err := teamsNotifier.Notify(&tc.payload)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if res.StatusCode != tc.code {
			t.Errorf("Expected %v, got %v", tc.code, res.StatusCode)
		}
		if res.Status != tc.body {
			t.Errorf("Expected %v, got %v", tc.body, res.Status)
		}
	}
}

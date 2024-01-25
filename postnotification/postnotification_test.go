package postnotification_test

import (
	"stackit/notifications/postnotification"
	"testing"
)

func TestValidate(t *testing.T) {
	// table driven tests
	var testCases = []struct {
		payload postnotification.PostNotification
		err     string
	}{
		{postnotification.PostNotification{
			Type:        "Warning",
			Name:        "Test Warning",
			Description: "This is a test warning",
			Channel:     "teams",
		}, ""},
		{postnotification.PostNotification{
			Type:        "Info",
			Name:        "Test Info",
			Description: "This is a test info",
			Channel:     "teams",
		}, ""},
		{postnotification.PostNotification{
			Type:        "Garbage",
			Name:        "Test Garbage",
			Description: "This is a test garbage",
			Channel:     "teams",
		}, ""},
		{postnotification.PostNotification{
			Type:        "Warning",
			Name:        "",
			Description: "This is a test warning",
			Channel:     "teams",
		}, "Name is required"},
		{postnotification.PostNotification{
			Type:        "Warning",
			Name:        "Test Warning",
			Description: "",
			Channel:     "teams",
		}, "Description is required"},
		{postnotification.PostNotification{
			Type:        "Warning",
			Name:        "Test Warning",
			Description: "This is a test warning",
		}, "Channel is required"},
	}

	for _, tc := range testCases {
		err := tc.payload.Validate()
		if err != nil && err.Error() != tc.err {
			t.Errorf("Expected %v, got %v", tc.err, err.Error())
		}
	}
}

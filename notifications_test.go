package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"stackit/notifications/postnotification"
	"testing"

	"github.com/gin-gonic/gin"
)

var warningPayload = postnotification.PostNotification{
	Type:        "Warning",
	Name:        "Test Warning",
	Description: "This is a test warning",
	Channel:     "teams",
}

var infoPayload = postnotification.PostNotification{
	Type:        "Info",
	Name:        "Test Info",
	Description: "This is a test info",
	Channel:     "teams",
}

var garbagePayload = postnotification.PostNotification{
	Type:        "Garbage",
	Name:        "Test Garbage",
	Description: "This is a test garbage",
	Channel:     "teams",
}

var noChannelPayload = postnotification.PostNotification{
	Type:        "Garbage",
	Name:        "Test Garbage",
	Description: "This is a test garbage",
}

// table driven tests
var testCases = []struct {
	payload postnotification.PostNotification
	code    int
	body    string
}{
	{warningPayload, http.StatusOK, `{"status":"200 OK"}`},
	{infoPayload, http.StatusOK, `{"status":"Message not dispatched"}`},
	{garbagePayload, http.StatusOK, `{"status":"Message not dispatched"}`},
	{noChannelPayload, http.StatusBadRequest, `{"error":"Channel is required"}`},
}

func TestPostNotifications(t *testing.T) {
	router := gin.Default()

	router.POST("/notifications", createNotification)

	for _, tc := range testCases {

		// payload as io.Reader
		payloadBytes, err := json.Marshal(tc.payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/notifications", bytes.NewReader(payloadBytes))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		// Check the response status code
		if recorder.Code != tc.code {
			t.Errorf("Expected status code %d, but got %d", tc.code, recorder.Code)
		}

		// Check the response body
		expected := tc.body
		if recorder.Body.String() != expected {
			t.Errorf("Expected the body to contain %q but got %q", expected, recorder.Body.String())
		}
	}
}

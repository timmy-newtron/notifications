package main

import (
	"net/http"
	"stackit/notifications/notifier"
	"stackit/notifications/postnotification"
	"stackit/notifications/teamsnotifier"

	"github.com/gin-gonic/gin"
)

var teamsNotifier = &teamsnotifier.TeamsNotifier{
	WebhookUrl: "YourSuperSecretWebhookUrlHere",
}

var notifiers = map[string]notifier.Notifier{
	"teams": teamsNotifier,
}

func main() {

	router := gin.Default()

	router.POST("/notifications", createNotification)

	router.Run(":8080")
}

func createNotification(c *gin.Context) {
	// bind json to struct
	var notification postnotification.PostNotification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate struct
	if err := notification.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// choose notifier
	nf := notifiers[notification.Channel]

	if nf == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No notifier found for channel \"" + notification.Channel + "\""})
		return
	}

	res, err := nf.Notify(&notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(res.StatusCode, gin.H{"status": res.Status})
}

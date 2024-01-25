# notifications

A simple notification microservice for sending notifications to users.

## Requirements

This project is written in Go. You can install Go from [here](https://golang.org/doc/install).

## Running

To run the project, you can use the following command:

```bash
go run notifications.go
```

This will start a web server listening on port 8080. You can send notifications to the server using the API described below.

## Testing

To run the tests, you can use the following command:

```bash
go test ./...
```

## API

### POST /notifications

Send a notification to a user.

#### Request

```json
{
  "type": "Warning" | "Info",
  "name": "string",
  "description": "string",
  "channel": "teams",
}
```

#### Success Response

```json
{
  "status": "200 OK"
}
```

#### Error Response

```json
{
  "status": "Message not dispatched"
}
```

## Extending the API

The notification service is designed to be extended. All notifications are sent by a Notifier. The Notifier interface is defined as follows:

```go
type Notifier interface {
	Notify(*postnotification.PostNotification) (*NotificationResponse, error)
}
```

where `PostNotification` is defined as follows:

```go
type NotificationResponse struct {
	StatusCode int
	Status     string
}
```

To add a new notifier, you can implement the `Notifier` interface and add it to the `notifiers` map in `main.go`. The notifiers map maps a string to a notifier. The string is the name of the notifier, and the notifier is the notifier itself. For example, to add a new notifier called `Slack`, you can do the following:

```go
notifiers["slack"] = &SlackNotifier{}
```

This would dispatch notifications with the `slack` channel to the `SlackNotifier` notifier.


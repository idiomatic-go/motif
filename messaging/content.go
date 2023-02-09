package messaging

import (
	"context"
	"github.com/idiomatic-go/motif/runtime"
)

// ContentMap - slice of any content to be included in a message
type ContentMap map[string][]any

// Credentials - type for a credentials function
type Credentials func() (username string, password string, err error)

// DatabaseUrl - struct for a database connection string
type DatabaseUrl struct {
	Url string
}

// ActuatorApply - type for applying an actuator
type ActuatorApply func(ctx context.Context, status **runtime.Status, uri, requestId, method string) (fn ActuatorComplete, newCtx context.Context, rateLimited bool)

// ActuatorComplete - return value from a call to ActuatorApply
type ActuatorComplete func()

// AccessCredentials - access function for Credentials in a message
func AccessCredentials(msg *Message) Credentials {
	if msg == nil || msg.Content == nil {
		return nil
	}
	for _, c := range msg.Content {
		if fn, ok := c.(Credentials); ok {
			return fn
		}
	}
	return nil
}

// AccessDatabaseUrl - access function for a DatabaseUrl in a message
func AccessDatabaseUrl(msg *Message) DatabaseUrl {
	if msg == nil || msg.Content == nil {
		return DatabaseUrl{}
	}
	for _, c := range msg.Content {
		if url, ok := c.(DatabaseUrl); ok {
			return url
		}
	}
	return DatabaseUrl{}
}

// AccessActuatorApply - access function for ActuatorApply in a message
func AccessActuatorApply(msg *Message) ActuatorApply {
	if msg == nil || msg.Content == nil {
		return nil
	}
	for _, c := range msg.Content {
		if fn, ok := c.(ActuatorApply); ok {
			return fn
		}
	}
	return nil
}

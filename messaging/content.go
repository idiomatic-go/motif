package messaging

import (
	"context"
	"github.com/idiomatic-go/motif/runtime"
)

type ContentMap map[string][]any

type Credentials func() (username string, password string, err error)

type DatabaseUrl struct {
	Url string
}

type ActuatorApply func(ctx context.Context, status **runtime.Status, uri, requestId, method string) (fn ActuatorComplete, newCtx context.Context, rateLimited bool)

type ActuatorComplete func()

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

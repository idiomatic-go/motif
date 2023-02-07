package messaging

import (
	"github.com/idiomatic-go/motif/runtime"
)

const (
	StartupEvent  = "event:startup"
	ShutdownEvent = "event:shutdown"
	PingEvent     = "event:ping"
	StatusEvent   = "event:status"
	FailoverEvent = "event:failover"
	HostName      = "host"
)

type MessageHandler func(msg Message)

type Message struct {
	To      string
	From    string
	Event   string
	Status  *runtime.Status
	Content []any
	ReplyTo MessageHandler
}

func ReplyTo(msg Message, status *runtime.Status) {
	if msg.ReplyTo == nil {
		return
	}
	msg.ReplyTo(Message{
		To:      msg.From,
		From:    msg.To,
		Event:   msg.Event,
		Status:  status,
		Content: nil,
		ReplyTo: nil,
	})
}

func NewMessageCacheHandler(cache *MessageCache) MessageHandler {
	return func(msg Message) {
		cache.Add(msg)
	}
}

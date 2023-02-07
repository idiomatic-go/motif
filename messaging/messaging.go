package messaging

import (
	"github.com/idiomatic-go/middleware/template"
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
	Status  *template.Status
	Content []any
	ReplyTo MessageHandler
}

/*
func NewStartupSuccessfulMessage(from Message) Message {
	return Message{To: from.From, From: from.To, Event: StartupEvent, Status: template.StatusOK}
}

func NewStartupFailureMessage(from Message) Message {
	return Message{To: from.From, From: from.To, Event: StartupEvent, Status: template.StatusInternal}
}

func StartupReplyTo(msg Message, successful bool, content ...any) {
	if msg.ReplyTo != nil && msg.Event == StartupEvent {
		var msg2 Message
		if successful {
			msg2 = NewStartupSuccessfulMessage(msg)
		} else {
			msg2 = NewStartupFailureMessage(msg)
		}
		msg2.Content = append(msg.Content, content...)
		msg.ReplyTo(msg2)
	}
}


*/

func ReplyTo(msg Message, status *template.Status) {
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

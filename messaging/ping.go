package messaging

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/middleware/template"
	"time"
)

const (
	maxWait = time.Second * 2
)

var pingLocation = pkgPath + "/ping"

func Ping[E template.ErrorHandler](ctx context.Context, uri string) (status *template.Status) {
	var e E

	if uri == "" {
		return e.HandleWithContext(ctx, pingLocation, errors.New("invalid argument: resource uri is empty"))
	}
	cache := NewMessageCache()
	msg := Message{To: uri, From: HostName, Event: PingEvent, Status: nil, ReplyTo: NewMessageCacheHandler(cache)}
	err := directory.Send(msg)
	if err != nil {
		return e.HandleWithContext(ctx, pingLocation, err)
	}
	duration := maxWait
	for wait := time.Duration(float64(duration) * 0.20); duration >= 0; duration -= wait {
		time.Sleep(wait)
		result, err1 := cache.Get(uri)
		if err1 != nil {
			continue //return e.HandleWithContext(ctx, pingLocation, err1)
		}
		if result.Status == nil {
			return e.HandleWithContext(ctx, pingLocation, errors.New(fmt.Sprintf("ping response status not available: [%v]", uri))).SetCode(template.StatusNotProvided)
		}
		return result.Status
	}
	return e.HandleWithContext(ctx, pingLocation, errors.New(fmt.Sprintf("ping response time out: [%v]", uri))).SetCode(template.StatusDeadlineExceeded)
}

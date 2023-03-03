package runtime

import (
	"context"
	"errors"
	http2 "github.com/idiomatic-go/motif/http"
	"net/http"
	"time"
)

type Context interface {
	context.Context
	http2.Exchange
	t() *context2
	withValue(key, val any) context.Context
}

type context2 struct {
	ctx  context.Context
	exec http2.Exchange
}

func NewContext(ctx context.Context, exec http2.Exchange) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return &context2{ctx: ctx, exec: exec}
}

func (c *context2) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *context2) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *context2) Err() error {
	return c.ctx.Err()
}

func (c *context2) Value(key any) any {
	return c.ctx.Value(key)
}

func (c *context2) Do(req *http.Request) (*http.Response, error) {
	if c.exec == nil {
		return nil, errors.New("invalid argument: Exchange interface is nil")
	}
	return c.exec.Do(req)
}

func (c *context2) t() *context2 {
	return c
}

func (c *context2) withValue(key, val any) context.Context {
	c.ctx = context.WithValue(c.ctx, key, val)
	return c
}

func ContextWithValue(ctx context.Context, key any, val any) context.Context {
	if ctx == nil {
		return nil
	}
	if curr, ok := any(ctx).(Context); ok {
		//ctx2 := curr.t()
		//ctx2.ctx = context.WithValue(ctx2.ctx, key, val)
		//ctx2.ctx = context.WithValue(ctx2.ctx, key, val)
		//return ctx2
		return curr.withValue(key, val)
	}
	return ctx
}

func IsContext(ctx context.Context) bool {
	if _, ok := any(ctx).(Context); ok {
		return true
	}
	return false
}

package runtime

import (
	"context"
	"errors"
	"github.com/google/uuid"
	http2 "github.com/idiomatic-go/motif/http"
	"net/http"
	"time"
)

const (
	xRequestIdName = "x-request-id"
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "context value " + k.name }

var (
	requestContextKey = &contextKey{"request-id"}
)

// ContextWithRequestId - creates a new Context with a request id
func ContextWithRequestId(ctx context.Context, requestId string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	} else {
		i := ctx.Value(requestContextKey)
		if i != nil {
			return ctx
		}
	}
	if requestId == "" {
		requestId = uuid.New().String()
	}
	return context.WithValue(ctx, requestContextKey, requestId)
}

// ContextWithRequest - creates a new Context with a request id from the request headers
func ContextWithRequest(req *http.Request) context.Context {
	if req == nil || req.Header == nil {
		return context.Background()
	}
	if req.Header.Get(xRequestIdName) == "" {
		req.Header.Add(xRequestIdName, uuid.New().String())
	}
	return ContextWithRequestId(req.Context(), req.Header.Get(xRequestIdName))
}

// ContextRequestId - return the requestId from a Context
func ContextRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	i := ctx.Value(requestContextKey)
	if requestId, ok := i.(string); ok {
		return requestId
	}
	return ""
}

type ExchangeContext interface {
	context.Context
	http2.Exchange
	withValue(key, val any) context.Context
}

type exchangeContext struct {
	ctx  context.Context
	exec http2.Exchange
}

func NewExchangeContext(ctx context.Context, exec http2.Exchange) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return &exchangeContext{ctx: ctx, exec: exec}
}

func (c *exchangeContext) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *exchangeContext) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *exchangeContext) Err() error {
	return c.ctx.Err()
}

func (c *exchangeContext) Value(key any) any {
	return c.ctx.Value(key)
}

func (c *exchangeContext) Do(req *http.Request) (*http.Response, error) {
	if c.exec == nil {
		return nil, errors.New("invalid argument: Exchange interface is nil")
	}
	return c.exec.Do(req)
}

//func (c *exchangeContext) t() *exchangeContext {
//	return c
//}

func (c *exchangeContext) withValue(key, val any) context.Context {
	c.ctx = context.WithValue(c.ctx, key, val)
	return c
}

func ContextWithValue(ctx context.Context, key any, val any) context.Context {
	if ctx == nil {
		return nil
	}
	if curr, ok := any(ctx).(ExchangeContext); ok {
		//ctx2 := curr.t()
		//ctx2.ctx = context.WithValue(ctx2.ctx, key, val)
		//ctx2.ctx = context.WithValue(ctx2.ctx, key, val)
		//return ctx2
		return curr.withValue(key, val)
	}
	return ctx
}

func IsExchangeContext(ctx context.Context) bool {
	if _, ok := any(ctx).(ExchangeContext); ok {
		return true
	}
	return false
}

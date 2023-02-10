package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return k.name }

var (
	doContextKey = &contextKey{"http-do"}
)

// DoProxy - Http client.Do proxy type
type DoProxy func(req *http.Request) (*http.Response, error)

// ContextWithDo - DoProxy context creation
func ContextWithDo(ctx context.Context, fn DoProxy) context.Context {
	if ctx == nil {
		ctx = context.Background()
	} else {
		if IsContextDo(ctx) {
			return ctx
		}
	}
	if fn == nil {
		return ctx
	}
	return &doContext{ctx, doContextKey, fn}
}

// ContextDo - call the DoProxy contained in the req.Context
func ContextDo(req *http.Request) (*http.Response, error) {
	if req == nil || req.Context() == nil {
		return nil, errors.New("context or request is nil")
	}
	i := req.Context().Value(doContextKey)
	if i == nil {
		return nil, errors.New(fmt.Sprintf("context value is nil for key: [%v]", doContextKey))
	}
	if do, ok := i.(DoProxy); ok && do != nil {
		return do(req)
	}
	return nil, errors.New("context proxy is not a valid type DoProxy(req *http.Request)")
}

// IsContextDoInRequest - determine if this request contains a DoProxy
func IsContextDoInRequest(req *http.Request) bool {
	if req == nil {
		return false
	}
	return IsContextDo(req.Context())
}

// IsContextDo - determine if this context contains a DoProxy
func IsContextDo(c context.Context) bool {
	if c == nil {
		return false
	}
	for {
		switch c.(type) {
		case *doContext:
			return true
		default:
			return false
		}
	}
}

type doContext struct {
	ctx      context.Context
	key, val any
}

func (*doContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*doContext) Done() <-chan struct{} {
	return nil
}

func (*doContext) Err() error {
	return nil
}

func (v *doContext) Value(any) any {
	return v.val
}

package runtime

import (
	"context"
	"github.com/google/uuid"
	"net/http"
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

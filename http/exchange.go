package http

import (
	"context"
	"net/http"
)

type Exchange interface {
	Do(req *http.Request) (*http.Response, error)
}

type exchange struct {
	exec func(*http.Request) (*http.Response, error)
}

func NewExchange(fn func(req *http.Request) (resp *http.Response, err error)) Exchange {
	return &exchange{fn}
}

func (e *exchange) Do(req *http.Request) (*http.Response, error) {
	return e.exec(req)
}

func exchangeCast(ctx context.Context) (Exchange, bool) {
	if ctx == nil {
		return nil, false
	}
	if e, ok := any(ctx).(Exchange); ok {
		return e, true
	}
	return nil, false
}

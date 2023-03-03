package http

import (
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
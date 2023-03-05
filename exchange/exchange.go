package exchange

import (
	"net/http"
)

type Exchange interface {
	Do(req *http.Request) (*http.Response, error)
}

type DefaultExchange struct{}

func (DefaultExchange) Do(req *http.Request) (*http.Response, error) {
	return Do(req)
}

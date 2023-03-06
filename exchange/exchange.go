package exchange

import (
	"net/http"
)

type Exchange interface {
	Do(req *http.Request) (*http.Response, error)
}

type Default struct{}

func (Default) Do(req *http.Request) (*http.Response, error) {
	return Do(req)
}

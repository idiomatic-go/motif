package exchange

import (
	"net/http"
)

// Exchange - interface for Http request/response interaction
type Exchange interface {
	Do(req *http.Request) (*http.Response, error)
}

type Default struct{}

func (Default) Do(req *http.Request) (*http.Response, error) {
	return Do(req)
}

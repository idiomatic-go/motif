package template

import (
	mhttp "github.com/idiomatic-go/motif/http"
	"net/http"
)

type HttpExchange interface {
	Do(req *http.Request) (*http.Response, error)
}

type DefaultClient struct{}

func (DefaultClient) Do(req *http.Request) (*http.Response, error) {
	return mhttp.Do(req)
}

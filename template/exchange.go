package template

import (
	http2 "github.com/idiomatic-go/motif/http"
	"net/http"
)

type DefaultExchange struct{}

func (DefaultExchange) Do(req *http.Request) (*http.Response, error) {
	return http2.Do(req)
}

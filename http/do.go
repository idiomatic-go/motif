package http

import (
	"errors"
	"net/http"
)

// Do - process a "client.Do" request with the http.DefaultClient
func Do(req *http.Request) (*http.Response, error) {
	return DoClient(req, http.DefaultClient)
}

// DoClient - process a "client.Do" request with the given client. Also, check the req.Context to determine
// if there is a "Do" proxy that should be called instead.
func DoClient(req *http.Request, client *http.Client) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("request is nil") //NewStatus(StatusInvalidArgument, doLocation, errors.New("request is nil"))
	}
	if client == nil {
		return nil, errors.New("client is nil")
	}
	//if e,ok := any(req.Context()).(Exchange); ok {
	//	e.Do()xc
	//}
	if IsContextDoInRequest(req) {
		return ContextDo(req)
	}
	return client.Do(req)
}

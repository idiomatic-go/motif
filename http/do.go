package http

import (
	"errors"
	"net/http"
)

func Do(req *http.Request) (*http.Response, error) {
	return DoClient(req, http.DefaultClient)
}

func DoClient(req *http.Request, client *http.Client) (*http.Response, error) {
	if req == nil {
		return nil, errors.New("request is nil") //NewStatus(StatusInvalidArgument, doLocation, errors.New("request is nil"))
	}
	if client == nil {
		return nil, errors.New("client is nil")
	}
	if IsContextDo(req.Context()) {
		return ContextDo(req)
	}
	return client.Do(req)
}

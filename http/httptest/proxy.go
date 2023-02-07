package httptest

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	HttpErrorUri   = "proxy://www.httperror.com"
	BodyIOErrorUri = "proxy://www.bodyioerror.com"
)

var bodyIOErrorResponse = NewIOErrorResponse()

// ErrorProxy - test proxy for http and body I/O errors
func ErrorProxy(req *http.Request) (*http.Response, error) {
	if req == nil || req.URL == nil {
		return nil, errors.New("request or URL is nil")
	}
	switch req.URL.String() {
	case HttpErrorUri:
		return HttpErrorResponse(req)
	case BodyIOErrorUri:
		return BodyIOErrorResponse(req)
	}
	return nil, errors.New(fmt.Sprintf("invalid request URL: %v", req.URL.String()))
}

// HttpErrorResponse - reusable http error response
func HttpErrorResponse(req *http.Request) (*http.Response, error) {
	return nil, http.ErrHijacked
}

// BodyIOErrorResponse - reusable body I/O error response
func BodyIOErrorResponse(req *http.Request) (*http.Response, error) {
	return bodyIOErrorResponse, nil
}

func Pattern(req *http.Request) string {
	if req == nil || req.URL == nil {
		return ""
	}
	pattern := req.URL.Scheme + "://" + req.URL.Host
	if req.URL.Path != "" {
		pattern += "/" + req.URL.Path
	}
	return pattern
}

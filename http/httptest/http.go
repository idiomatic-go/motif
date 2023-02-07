package httptest

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	HttpErrorUri   = "proxy://www.httperror.com"
	BodyIOErrorUri = "proxy://www.bodyioerror.com"
)

var BodyIOErrorResponse = NewIOErrorResponse()

func ErrorProxy(req *http.Request) (*http.Response, error) {
	if req == nil || req.URL == nil {
		return nil, errors.New("request or URL is nil")
	}
	switch req.URL.String() {
	case HttpErrorUri:
		return HttpErrorProxy(req)
	case BodyIOErrorUri:
		return BodyIOErrorProxy(req)
	}
	return nil, errors.New(fmt.Sprintf("invalid request URL: %v", req.URL.String()))
}

func HttpErrorProxy(req *http.Request) (*http.Response, error) {
	return nil, http.ErrHijacked
}

func BodyIOErrorProxy(req *http.Request) (*http.Response, error) {
	return BodyIOErrorResponse, nil
}

func NewResponse(httpStatus int, content []byte, kv ...string) *http.Response {
	if len(kv)&1 == 1 {
		kv = append(kv, "dummy header value")
	}
	resp := &http.Response{StatusCode: httpStatus, Header: make(http.Header), Request: nil}
	resp.Body = NewReaderCloser(bytes.NewReader(content), nil)
	for i := 0; i < len(kv); i += 2 {
		key := strings.ToLower(kv[i])
		resp.Header.Add(key, kv[i+1])
	}
	return resp
}

func NewIOErrorResponse() *http.Response {
	resp := &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Request: nil}
	resp.Body = NewReaderCloser(nil, io.ErrUnexpectedEOF)
	return resp
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

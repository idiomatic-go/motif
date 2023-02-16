package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// HttpLog - type for http logging
type HttpLog func(traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response)

// SetLogFn - allows setting an application configured logging function
func SetLogFn(fn HttpLog) {
	if fn != nil {
		logFn = fn
	}
}

var logFn = defaultLogFn

var defaultLogFn HttpLog = func(traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response) {
	s := fmt.Sprintf("start:%v ,"+
		"duration:%v ,"+
		"traffic:%v, "+
		"status-code:%v, "+
		"method:%v, "+
		"url:%v, "+
		"host:%v, "+
		"path:%v, "+
		"bytes:%v",
		start, duration, traffic, resp.StatusCode, req.Method, req.URL.String(), req.Host, req.URL.Path, resp.ContentLength)
	fmt.Printf("%v\n", s)
}

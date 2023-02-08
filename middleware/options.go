package middleware

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
	"github.com/idiomatic-go/motif/accesslog"
)

// SetLogFn - allows setting an application configured logging function
func SetLogFn(fn accesslog.LogFn) {
	if fn != nil {
		logFn = fn
	}
}

var logFn = defaultLog

var defaultLog accesslog.LogFn = func(entry *accessdata.Entry) {
	fmt.Printf("{%v}\n", entry)
}

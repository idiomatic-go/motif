package middleware

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
)

// SetLogFn - allows setting an application configured logging function
func SetLogFn(fn accessdata.Accessor) {
	if fn != nil {
		logFn = fn
	}
}

var logFn = defaultLog

var defaultLog accessdata.Accessor = func(entry *accessdata.Entry) {
	fmt.Printf("{%v}\n", entry)
}

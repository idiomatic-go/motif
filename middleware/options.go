package middleware

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
)

// LogAccessData - logging type
type LogAccessData func(entry *accessdata.Entry)

// SetLogFn - allows setting an application configured logging function
func SetLogFn(fn LogAccessData) {
	if fn != nil {
		logFn = fn
	}
}

var logFn LogAccessData = defaultAccess

var defaultAccess LogAccessData = func(entry *accessdata.Entry) {
	fmt.Printf("{%v}\n", entry)
}

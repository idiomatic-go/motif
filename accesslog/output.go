package accesslog

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
	"log"
)

// OutputHandler - interface used by Log to write the operators and data
type OutputHandler interface {
	Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter)
}

// NilOutputHandler - no output
type NilOutputHandler struct{}

func (NilOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter) {
}

// DebugOutputHandler - output to stdio
type DebugOutputHandler struct{}

func (DebugOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter) {
	fmt.Printf("%v\n", formatter.Format(items, data))
}

// TestOutputHandler - special use case of DebugOutputHandler for testing examples
type TestOutputHandler struct{}

func (TestOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter) {
	fmt.Printf("test: Write() -> [%v]\n", formatter.Format(items, data))
}

// LogOutputHandler - output to log
type LogOutputHandler struct{}

func (LogOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter) {
	log.Println(formatter.Format(items, data))
}

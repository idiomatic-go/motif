package accesslog

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
	"log"
)

// OutputHandler - interface used by Log to write the operators and data
type OutputHandler interface {
	Write(items []accessdata.Operator, data *accessdata.Entry)
}

// NilOutputHandler - no output
type NilOutputHandler struct{}

func (NilOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry) {}

// DebugOutputHandler - output to stdio
type DebugOutputHandler struct{}

func (DebugOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry) {
	fmt.Printf("%v\n", accessdata.WriteJson(items, data))
}

// LogOutputHandler - output to log
type LogOutputHandler struct{}

func (LogOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry) {
	log.Println(accessdata.WriteJson(items, data))
}

// TestOutputHandler - special use case of DebugOutputHandler for testing examples
type TestOutputHandler struct{}

func (TestOutputHandler) Write(items []accessdata.Operator, data *accessdata.Entry) {
	fmt.Printf("test: Write() -> [%v]\n", accessdata.WriteJson(items, data))
}

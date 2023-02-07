package accesslog

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
)

type OutputHandlerFn func(s string)

func (OutputHandlerFn) Write(s string) {

}

func ExampleOutputHandler() {
	fmt.Printf("test: Output[NilOutputHandler](s)\n")
	Output[NilOutputHandler](nil, nil)

	fmt.Printf("test: Output[DebugOutputHandler](s)\n")
	Output[DebugOutputHandler]([]accessdata.Operator{{"error", "message"}}, accessdata.NewEntry())

	fmt.Printf("test: Output[TestOutputHandler](s)\n")
	Output[TestOutputHandler](nil, nil)

	fmt.Printf("test: Output[LogOutputHandler](s)\n")
	Output[LogOutputHandler](nil, nil)

	//Output:
	//test: Output[NilOutputHandler](s)
	//test: Output[DebugOutputHandler](s)
	//{"error":"message"}
	//test: Output[TestOutputHandler](s)
	//test: Write() -> [{}]
	//test: Output[LogOutputHandler](s)

}

func Output[O OutputHandler](items []accessdata.Operator, data *accessdata.Entry) {
	var o O
	o.Write(items, data)
}

package accesslog

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
)

func ExampleOutputHandler() {
	fmt.Printf("test: Output[NilOutputHandler,accessdata.TextFormatter](nil,nil)\n")
	logTest[NilOutputHandler, accessdata.TextFormatter](nil, nil)

	fmt.Printf("test: Output[DebugOutputHandler,accessdata.JsonFormatter](operators,data)\n")
	ops := []accessdata.Operator{{"error", "message"}}
	logTest[DebugOutputHandler, accessdata.JsonFormatter](ops, accessdata.NewEntry())

	fmt.Printf("test: Output[TestOutputHandler,accessdata.JsonFormatter](nil,nil)\n")
	logTest[TestOutputHandler, accessdata.JsonFormatter](nil, nil)

	fmt.Printf("test: Output[TestOutputHandler,accessdata.JsonFormatter](ops,data)\n")
	logTest[TestOutputHandler, accessdata.JsonFormatter](ops, accessdata.NewEntry())

	fmt.Printf("test: Output[LogOutputHandler,accessdata.JsonFormatter](ops,data)\n")
	logTest[LogOutputHandler, accessdata.JsonFormatter](ops, accessdata.NewEntry())

	//Output:
	//test: Output[NilOutputHandler,accessdata.TextFormatter](nil,nil)
	//test: Output[DebugOutputHandler,accessdata.JsonFormatter](operators,data)
	//{"error":"message"}
	//test: Output[TestOutputHandler,accessdata.JsonFormatter](nil,nil)
	//test: Write() -> [{}]
	//test: Output[TestOutputHandler,accessdata.JsonFormatter](ops,data)
	//test: Write() -> [{"error":"message"}]
	//test: Output[LogOutputHandler,accessdata.JsonFormatter](ops,data)

}

func logTest[O OutputHandler, F accessdata.Formatter](items []accessdata.Operator, data *accessdata.Entry) {
	var o O
	var f F
	o.Write(items, data, f)
}

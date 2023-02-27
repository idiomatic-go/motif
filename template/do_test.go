package template

import (
	"fmt"
	"net/http"
)

func ExampleDo() {
	req, _ := http.NewRequest("GET", "https://www.google.com/search?q=test", nil)
	resp, buf, status := Do[DebugError, []byte, DefaultClient](req)
	fmt.Printf("test: Do[DebugError,[]byte,DefaultClient](req) -> [status:%v] [buf:%v] [resp:%v]\n", status, len(buf) > 0, resp != nil)

	//Output:
	//test: Do[DebugError,[]byte,DefaultClient](req) -> [status:OK] [buf:true] [resp:true]
}

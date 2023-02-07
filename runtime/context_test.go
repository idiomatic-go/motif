package runtime

import (
	"context"
	"fmt"
	"net/http"
)

func ExampleRequestId() {
	ctx := ContextWithRequestId(context.Background(), "123-456-abc")
	fmt.Printf("test: ContextRequestId(ctx) -> %v\n", ContextRequestId(ctx))

	//ctx = ContextWithRequestId(nil, "")
	//fmt.Printf("test: ContextRequestId(ctx) -> %v\n", ContextRequestId(ctx))

	//Output:
	//test: ContextRequestId(ctx) -> 123-456-abc

}

func _ExampleRequest() {
	ctx := ContextWithRequest(nil)
	fmt.Printf("test: ContextRequestId(ctx) -> %v\n", ContextRequestId(ctx))

	req, _ := http.NewRequest("", "https.www.google.com", nil)
	ctx = ContextWithRequest(req)
	fmt.Printf("test: ContextRequestId(ctx) -> %v [req:%v]\n", ContextRequestId(ctx), req.Header.Get(xRequestIdName))

	req.Header.Add(xRequestIdName, "x-request-id-value")
	ctx = ContextWithRequest(req)
	fmt.Printf("test: ContextRequestId(ctx) -> %v\n", ContextRequestId(ctx))

	//Output:
	//fail

}

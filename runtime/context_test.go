package runtime

import (
	"context"
	"fmt"
	"net/http"
)

func ExampleContextWithRequest() {
	ctx := ContextWithRequestId(context.Background(), "123-456-abc")
	fmt.Printf("test: ContextWithRequestId(ctx,id) -> %v\n", ContextRequestId(ctx))

	ctx = ContextWithRequest(nil)
	fmt.Printf("test: ContextWithRequest(nil) -> %v\n", ContextRequestId(ctx) != "")

	req, _ := http.NewRequest("", "https.www.google.com", nil)
	ctx = ContextWithRequest(req)
	fmt.Printf("test: ContextWithRequest(req) -> %v\n", ContextRequestId(ctx) != "")

	req, _ = http.NewRequest("", "https.www.google.com", nil)
	req.Header.Add(xRequestIdName, "x-request-id-value")
	ctx = ContextWithRequest(req)
	fmt.Printf("test: ContextWithRequest(req) -> %v\n", ContextRequestId(ctx))

	//Output:
	//test: ContextWithRequestId(ctx,id) -> 123-456-abc
	//test: ContextWithRequest(nil) -> false
	//test: ContextWithRequest(req) -> true
	//test: ContextWithRequest(req) -> x-request-id-value

}

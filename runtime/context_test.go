package runtime

import (
	"context"
	"errors"
	"fmt"
	http2 "github.com/idiomatic-go/motif/http"
	"net/http"
)

func testDo(req *http.Request) (*http.Response, error) {
	fmt.Printf("test: testDo() -> \n")
	return nil, errors.New("test error")
}

func ExampleContextWithRequestExisting() {
	ctx := ContextWithRequestId(context.Background(), "123-456-abc")
	fmt.Printf("test: ContextWithRequestId(context.Background(),id) -> %v [newContext:%v]\n", ContextRequestId(ctx), ctx != context.Background())

	ctxNew := ContextWithRequestId(ctx, "123-456-abc-xyz")
	fmt.Printf("test: ContextWithRequestId(ctx,id) -> %v [newContext:%v]\n", ContextRequestId(ctx), ctxNew != ctx)

	//Output:
	//test: ContextWithRequestId(context.Background(),id) -> 123-456-abc [newContext:true]
	//test: ContextWithRequestId(ctx,id) -> 123-456-abc [newContext:false]

}

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

func ExampleExchangeContext() {
	k1 := "1"
	k2 := "2"
	k3 := "3"
	v1 := "value 1"
	v2 := "value 2"
	v3 := "value 3"

	do1 := http2.NewExchange(testDo)
	ctx := NewExchangeContext(nil, do1)

	fmt.Printf("test: IsExchangeContext(ctx) -> %v\n", IsExchangeContext(ctx))
	fmt.Printf("test: Values() -> [key1:%v] [key2:%v] [key3:%v]\n", ctx.Value(k1), ctx.Value(k2), ctx.Value(k3))

	ctx1 := ContextWithValue(ctx, k1, v1)
	fmt.Printf("test: IsExchangeContext(ctx1) -> %v\n", IsExchangeContext(ctx1))
	fmt.Printf("test: Values() -> [key1:%v] [key2:%v] [key3:%v]\n", ctx1.Value(k1), ctx1.Value(k2), ctx1.Value(k3))

	ctx2 := ContextWithValue(ctx, k2, v2)
	fmt.Printf("test: IsExchangeContext(ctx2) -> %v\n", IsExchangeContext(ctx2))
	fmt.Printf("test: Values() -> [key1:%v] [key2:%v] [key3:%v]\n", ctx2.Value(k1), ctx2.Value(k2), ctx2.Value(k3))

	ctx3 := ContextWithValue(ctx, k3, v3)
	fmt.Printf("test: IsExchangeContext(ctx3) -> %v\n", IsExchangeContext(ctx3))
	fmt.Printf("test: Values() -> [key1:%v] [key2:%v] [key3:%v]\n", ctx3.Value(k1), ctx3.Value(k2), ctx3.Value(k3))

	//Output:
	//test: IsExchangeContext(ctx) -> true
	//test: Values() -> [key1:<nil>] [key2:<nil>] [key3:<nil>]
	//test: IsExchangeContext(ctx1) -> true
	//test: Values() -> [key1:value 1] [key2:<nil>] [key3:<nil>]
	//test: IsExchangeContext(ctx2) -> true
	//test: Values() -> [key1:value 1] [key2:value 2] [key3:<nil>]
	//test: IsExchangeContext(ctx3) -> true
	//test: Values() -> [key1:value 1] [key2:value 2] [key3:value 3]

}

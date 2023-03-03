package runtime

import (
	"errors"
	"fmt"
	http2 "github.com/idiomatic-go/motif/http"
	"net/http"
)

func testDo(req *http.Request) (*http.Response, error) {
	fmt.Printf("test: testDo() -> \n")
	return nil, errors.New("test error")
}

func ExampleNewContext_PrevValue() {
	ctx := ContextWithRequestId(nil, "1234-56-7890")
	do1 := http2.NewExchange(testDo)

	ctxNew := NewContext(ctx, do1)
	val := ctxNew.Value(requestContextKey)
	fmt.Printf("test: ctx.Value() -> [value:%v]\n", val)

	//Output:
	//test: ctx.Value() -> [value:1234-56-7890]

}

func ExampleNewContext_Fail() {
	do1 := http2.NewExchange(testDo)
	ctx := NewContext(nil, do1)

	fmt.Printf("test: NewContext() -> [ctx:%v] [newcontext:%v]\n", ctx, IsContext(ctx))

	ctx1 := ContextWithRequestId(ctx, "1234-56-7890")
	fmt.Printf("test: NewContext() -> [ctx:%v]\n", ctx1)

	val := ctx1.Value(requestContextKey)
	fmt.Printf("test: ctx.Value() -> [value:%v]\n", val)

	fmt.Printf("test: IsContext() -> [ctx:%v]\n", IsContext(ctx1))

	//Output:
	//fail
}

func ExampleNewContext_Success() {
	do1 := http2.NewExchange(testDo)
	ctx := NewContext(nil, do1)

	fmt.Printf("test: NewContext() -> [ctx:%v] [newcontext:%v]\n", ctx, IsContext(ctx))

	ctx1 := ContextWithValue(ctx, requestContextKey, "1234-56-7890")
	fmt.Printf("test: NewContext() -> [ctx:%v]\n", ctx1)

	val := ctx1.Value(requestContextKey)
	fmt.Printf("test: ctx.Value() -> [value:%v]\n", val)

	fmt.Printf("test: IsContext() -> [ctx:%v]\n", IsContext(ctx1))

	//Output:
	//fail
}

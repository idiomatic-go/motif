package http

import (
	"context"
	"fmt"
	"net/http"
)

func doDocxtProxy(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusOK
	return resp, nil
}

func ExampleContextWithDoExisting() {
	ctx := ContextWithDo(context.Background(), doDocxtProxy)
	fmt.Printf("test: ContextWithDo(context.Background(),doDocxtProxy) -> [ctxDo:%v] [newCtx:%v]\n", IsContextDo(ctx), ctx != context.Background())

	ctxNew := ContextWithDo(ctx, doDocxtProxy)
	fmt.Printf("test: ContextWithDo(ctx,doDocxtProxy) -> [ctxDo:%v] [newCtx:%v]\n", IsContextDo(ctx), ctxNew != ctx)

	//Output:
	//test: ContextWithDo(context.Background(),doDocxtProxy) -> [ctxDo:true] [newCtx:true]
	//test: ContextWithDo(ctx,doDocxtProxy) -> [ctxDo:true] [newCtx:false]

}

func ExampleContextWithDo() {
	ctx := ContextWithDo(nil, nil)
	fmt.Printf("test: ContextWithDo(nil,nil) -> [ctxDo:%v]\n", IsContextDo(ctx))

	ctx = ContextWithDo(context.Background(), nil)
	fmt.Printf("test: ContextWithDo(ctx,nil) -> [ctxDo:%v]\n", IsContextDo(ctx))

	ctx = ContextWithDo(context.Background(), doDocxtProxy)
	fmt.Printf("test: ContextWithDo(ctx,doDocxtProxy) -> [ctxDo:%v]\n", IsContextDo(ctx))

	//Output:
	//test: ContextWithDo(nil,nil) -> [ctxDo:false]
	//test: ContextWithDo(ctx,nil) -> [ctxDo:false]
	//test: ContextWithDo(ctx,doDocxtProxy) -> [ctxDo:true]

}

func ExampleContextDo() {
	resp, err := ContextDo(nil)
	fmt.Printf("test: ContextDo(nil) -> [error:%v] [resp:%v]\n", err, resp != nil)

	req, _ := http.NewRequest("", "https.www.google.com", nil)
	fmt.Printf("test: IsContextDoInRequest(req-nil ctx) -> %v\n", IsContextDoInRequest(req))

	resp, err = ContextDo(req)
	fmt.Printf("test: ContextDo(req-background ctx) -> [error:%v] [resp:%v]\n", err, resp != nil)

	req, _ = http.NewRequestWithContext(ContextWithDo(context.Background(), doDocxtProxy), "", "https.www.google.com", nil)
	fmt.Printf("test: IsContextDoInRequest(req) -> %v\n", IsContextDoInRequest(req))

	resp, err = ContextDo(req)
	fmt.Printf("test: ContextDo(req) -> [error:%v] [status:%v]\n", err, resp.StatusCode)

	//Output:
	//test: ContextDo(nil) -> [error:context or request is nil] [resp:false]
	//test: IsContextDoInRequest(req-nil ctx) -> false
	//test: ContextDo(req-background ctx) -> [error:context value is nil for key: [http-do]] [resp:false]
	//test: IsContextDoInRequest(req) -> true
	//test: ContextDo(req) -> [error:<nil>] [status:200]

}

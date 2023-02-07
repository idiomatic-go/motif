package http

import (
	"context"
	"fmt"
	"github.com/idiomatic-go/middleware/http/httptest"
	"io"
	"net/http"
)

const (
	httpContent = "proxy://www.somestupidname.come"
)

var doCtx = ContextWithDo(context.Background(), doProxy)

func doProxy(req *http.Request) (*http.Response, error) {
	if req == nil || req.URL == nil {
		return nil, nil
	}
	switch httptest.Pattern(req) {
	case httptest.HttpErrorUri, httptest.BodyIOErrorUri:
		return httptest.ErrorProxy(req)
	case httpContent:
		resp := httptest.NewResponse(http.StatusOK, []byte("<html><body><h1>Hello, World</h1></body></html>"), "content-type", "text/html", "content-length", "1234")
		return resp, nil
	default:
		fmt.Printf("test: doProxy(req) : unmatched pattern %v", httptest.Pattern(req))
	}
	return nil, nil
}

func ExampleDo_InvalidArgument() {
	_, s := Do(nil)
	fmt.Printf("test: Do(nil) -> [%v]\n", s)

	req, _ := http.NewRequest("", "http://www.google.com", nil)
	_, s = DoClient(req, nil)
	fmt.Printf("test: DoClient(req,nil) -> [%v]\n", s)

	//Output:
	//test: Do(nil) -> [request is nil]
	//test: DoClient(req,nil) -> [client is nil]

}

func ExampleDo_HttpError() {
	req, _ := http.NewRequestWithContext(doCtx, http.MethodGet, httptest.HttpErrorUri, nil)
	resp, status := Do(req)
	fmt.Printf("test: Do(req) -> [%v] [response:%v]\n", status, resp)

	//Output:
	//test: Do(req) -> [http: connection has been hijacked] [response:<nil>]

}

func ExampleDo_IOError() {
	req, _ := http.NewRequestWithContext(doCtx, http.MethodGet, httptest.BodyIOErrorUri, nil)
	resp, status := Do(req)
	fmt.Printf("test: Do(req) -> [%v] [resp:%v] [statusCode:%v] [body:%v]\n", status, resp != nil, resp.StatusCode, resp.Body != nil)

	defer resp.Body.Close()
	buf, s2 := io.ReadAll(resp.Body)
	fmt.Printf("test: io.ReadAll(resp.Body) -> [%v] [body:%v]\n", s2, string(buf))

	//Output:
	//test: Do(req) -> [<nil>] [resp:true] [statusCode:200] [body:true]
	//test: io.ReadAll(resp.Body) -> [unexpected EOF] [body:]

}

func ExampleDo_Content() {
	req, _ := http.NewRequestWithContext(doCtx, http.MethodGet, httpContent, nil)
	resp, status := Do(req)
	fmt.Printf("test: Do(req) -> [%v] [resp:%v] [statusCode:%v] [content-type:%v] [content-length:%v] [body:%v]\n",
		status, resp != nil, resp.StatusCode, resp.Header.Get("content-type"), resp.Header.Get("content-length"), resp.Body != nil)

	defer resp.Body.Close()
	buf, ioError := io.ReadAll(resp.Body)
	fmt.Printf("test: io.ReadAll(resp.Body) -> [err:%v] [body:%v]\n", ioError, string(buf))

	//Output:
	//test: Do(req) -> [<nil>] [resp:true] [statusCode:200] [content-type:text/html] [content-length:1234] [body:true]
	//test: io.ReadAll(resp.Body) -> [err:<nil>] [body:<html><body><h1>Hello, World</h1></body></html>]

}

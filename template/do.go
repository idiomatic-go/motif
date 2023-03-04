package template

import (
	http2 "github.com/idiomatic-go/motif/http"
	"github.com/idiomatic-go/motif/runtime"
	"net/http"
)

func Do[E ErrorHandler, T any, H http2.Exchange](req *http.Request) (resp *http.Response, t T, status *runtime.Status) {
	var e E
	var h H

	if req == nil {
		return nil, t, runtime.NewHttpStatusCode(http.StatusInternalServerError)
	}
	var err error
	resp, err = h.Do(req)
	if err != nil {
		return nil, t, e.HandleWithContext(req.Context(), "Do", err)
	}
	if resp == nil {
		return nil, t, runtime.NewHttpStatusCode(http.StatusInternalServerError)
	}
	if resp.StatusCode != http.StatusOK {
		return resp, t, runtime.NewHttpStatusCode(resp.StatusCode)
	}
	t, status = Deserialize[E, T](resp.Body)
	return
}

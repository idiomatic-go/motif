package middleware

import (
	"github.com/idiomatic-go/motif/runtime"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, buf []byte, status *runtime.Status, headers ...string) {
	if status != nil {
		w.WriteHeader(status.Http())
	}
	for _, name := range headers {
		status.AddMetadata(w.Header(), name)
	}
	if status.OK() {
		if buf != nil {
			w.Write(buf)
		}
	} else {
		if buf1 := status.Content(); buf1 != nil && status != nil {
			w.Write(buf1)
		}
	}
}

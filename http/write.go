package http

import (
	"github.com/idiomatic-go/motif/runtime"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, data []byte, status *runtime.Status, headers ...string) {
	if status != nil {
		w.WriteHeader(status.Http())
	}
	status.AddMetadata(w.Header(), headers...)
	if status.OK() {
		if data != nil {
			w.Write(data)
		}
	} else {
		if buf := status.Content(); buf != nil && status != nil {
			w.Write(buf)
		}
	}
}

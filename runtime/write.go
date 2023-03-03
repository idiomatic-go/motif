package runtime

import (
	"net/http"
)

func WriteResponse(w http.ResponseWriter, data []byte, status *Status, headers ...string) {
	if status == nil {
		w.WriteHeader(http.StatusOK)
		if data != nil {
			w.Write(data)
		}
		return
	}
	status.AddMetadata(w.Header(), headers...)
	w.WriteHeader(status.Http())
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

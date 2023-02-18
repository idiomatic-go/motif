package runtime

import (
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ContentTypeText = "text/plain" // charset=utf-8
	ContentTypeJson = "application/json"

	ContentLocation = "Content-Location"
)

func GetContentLocation(req *http.Request) string {
	if req != nil {
		return req.Header.Get(ContentLocation)
	}
	return ""
}

package middleware

import (
	"github.com/felixge/httpsnoop"
	"github.com/idiomatic-go/motif/accessdata"
	"net/http"
	"time"
)

// HttpHostMetricsHandler - http handler that captures metrics about an ingress request, also logs an access
// entry
func HttpHostMetricsHandler(appHandler http.Handler, msg string) http.Handler {
	wrappedH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		m := httpsnoop.CaptureMetrics(appHandler, w, r)
		//log.Printf("%s %s (code=%d dt=%s written=%d)", r.Method, r.URL, m.Code, m.Duration, m.Written)
		logFn(accessdata.NewHttpIngressEntry(start, time.Since(start), nil, r, m.Code, m.Written, ""))
	})
	return wrappedH
}

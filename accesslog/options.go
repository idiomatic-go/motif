package accesslog

import "github.com/idiomatic-go/motif/accessdata"

type LogFn func(entry *accessdata.Entry)

// SetIngressLogStatus - enable/disable ingress logging
func SetIngressLogStatus(enabled bool) {
	opt.ingress = enabled
}

// SetEgressLogStatus - enable/disable egress logging
func SetEgressLogStatus(enabled bool) {
	opt.egress = enabled
}

type options struct {
	ingress bool
	egress  bool
}

var opt options

func init() {
	opt.ingress = true
	opt.egress = true
}

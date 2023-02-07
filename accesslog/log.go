package accesslog

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
)

const (
	errorName     = "error"
	errorNilEntry = "access data entry is nil"
	errorEmptyFmt = "%v log entries are empty"
)

var ingressOperators []accessdata.Operator
var egressOperators []accessdata.Operator

// InitIngressOperators - allows configuration of access log attributes for ingress traffic
func InitIngressOperators(config []accessdata.Operator) error {
	var err error
	ingressOperators, err = accessdata.InitOperators(config)
	return err
}

// InitEgressOperators - allows configuration of access log attributes for egress traffic
func InitEgressOperators(config []accessdata.Operator) error {
	var err error
	egressOperators, err = accessdata.InitOperators(config)
	return err
}

// Log - handles writing the access log entry via the OutputHandler
func Log[O OutputHandler](entry *accessdata.Entry) {
	var o O

	if entry == nil {
		o.Write([]accessdata.Operator{{errorName, errorNilEntry}}, accessdata.NewEntry())
		return
	}
	if entry.IsIngress() {
		if !opt.ingress {
			return
		}
		if len(ingressOperators) == 0 {
			o.Write([]accessdata.Operator{{errorName, fmt.Sprintf(errorEmptyFmt, entry.Traffic)}}, accessdata.NewEntry())
			return
		}
		o.Write(ingressOperators, entry)
	} else {
		if !opt.egress {
			return
		}
		if len(egressOperators) == 0 {
			o.Write([]accessdata.Operator{{errorName, fmt.Sprintf(errorEmptyFmt, entry.Traffic)}}, accessdata.NewEntry())
			return
		}
		o.Write(egressOperators, entry)
	}
}

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

// Log - handles writing the access log entry utilizing the OutputHandler and Formatter
func Log[O OutputHandler, F accessdata.Formatter](data *accessdata.Entry) {
	var o O
	var f F
	if data == nil {
		o.Write([]accessdata.Operator{{errorName, errorNilEntry}}, accessdata.NewEntry(), f)
		return
	}
	if data.IsIngress() {
		if !opt.ingress {
			return
		}
		if len(ingressOperators) == 0 {
			o.Write(emptyOperators(data), accessdata.NewEntry(), f)
			return
		}
		o.Write(ingressOperators, data, f)
	} else {
		if !opt.egress {
			return
		}
		if len(egressOperators) == 0 {
			o.Write(emptyOperators(data), accessdata.NewEntry(), f)
			return
		}
		o.Write(egressOperators, data, f)
	}
}

func emptyOperators(data *accessdata.Entry) []accessdata.Operator {
	return []accessdata.Operator{{errorName, fmt.Sprintf(errorEmptyFmt, data.Traffic)}}
}

/*
func Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter) {
	var o O
	var f accessdata.Formatter

	if data == nil {
		o.Write([]accessdata.Operator{{errorName, errorNilEntry}}, accessdata.NewEntry(),f)
		return
	}
	if entry.IsIngress() {
		if !opt.ingress {
			return
		}
		if len(ingressOperators) == 0 {
			o.Write([]accessdata.Operator{{errorName, fmt.Sprintf(errorEmptyFmt, entry.Traffic)}}, accessdata.NewEntry(),f)
			return
		}
		o.Write(ingressOperators, entry,f)
	} else {
		if !opt.egress {
			return
		}
		if len(egressOperators) == 0 {
			o.Write([]accessdata.Operator{{errorName, fmt.Sprintf(errorEmptyFmt, entry.Traffic)}}, accessdata.NewEntry(),f)
			return
		}
		o.Write(egressOperators, entry,f)
	}
}

*/

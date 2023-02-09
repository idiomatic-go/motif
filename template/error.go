package template

import (
	"context"
	"fmt"
	"github.com/idiomatic-go/motif/runtime"
	"log"
)

// ErrorHandleFn - function type for error handling
type ErrorHandleFn func(location string, errs ...error) *runtime.Status

// ErrorStatusHandleFn - function type for error status handling
type ErrorStatusHandleFn func(s *runtime.Status) *runtime.Status

// ErrorHandler - error handler interface
type ErrorHandler interface {
	Handle(location string, errs ...error) *runtime.Status
	HandleWithContext(ctx context.Context, location string, errs ...error) *runtime.Status
	HandleStatus(s *runtime.Status) *runtime.Status
}

// NoOpError - no operation on errors
type NoOpError struct{}

func (NoOpError) Handle(location string, errs ...error) *runtime.Status {
	if !runtime.IsErrors(errs) {
		return runtime.NewStatusOK()
	}
	return runtime.NewStatusError(location, errs...)
}

func (NoOpError) HandleWithContext(ctx context.Context, location string, errs ...error) *runtime.Status {
	if !runtime.IsErrors(errs) {
		return runtime.NewStatusOK()
	}
	return runtime.NewStatusError(location, errs...).SetContext(ctx)
}

func (NoOpError) HandleStatus(s *runtime.Status) *runtime.Status {
	return s
}

// DebugError - debug error handler
type DebugError struct{}

func (h DebugError) Handle(location string, errs ...error) *runtime.Status {
	if !runtime.IsErrors(errs) {
		return runtime.NewStatusOK()
	}
	return h.HandleStatus(runtime.NewStatus(runtime.StatusInternal, location, errs...))
}

func (h DebugError) HandleWithContext(ctx context.Context, location string, errs ...error) *runtime.Status {
	if !runtime.IsErrors(errs) {
		return runtime.NewStatusOK()
	}
	return h.HandleStatus(runtime.NewStatus(runtime.StatusInternal, location, errs...).SetContext(ctx))
}

func (h DebugError) HandleStatus(s *runtime.Status) *runtime.Status {
	if s != nil && s.IsErrors() {
		loc := IfElse[string](s.Location() == "", "[]", s.Location())
		req := IfElse[string](s.RequestId() == "", "[]", s.RequestId())
		fmt.Printf("[%v %v %v]\n", req, loc, s.Errors())
		s.RemoveErrors()
	}
	return s
}

// LogError - logging error handler
type LogError struct{}

func (h LogError) Handle(location string, errs ...error) *runtime.Status {
	if !runtime.IsErrors(errs) {
		return runtime.NewStatusOK()
	}
	return h.HandleStatus(runtime.NewStatus(runtime.StatusInternal, location, errs...))
}

func (h LogError) HandleWithContext(ctx context.Context, location string, errs ...error) *runtime.Status {
	if !runtime.IsErrors(errs) {
		return runtime.NewStatusOK()
	}
	return h.HandleStatus(runtime.NewStatus(runtime.StatusInternal, location, errs...).SetContext(ctx))
}

func (h LogError) HandleStatus(s *runtime.Status) *runtime.Status {
	if s != nil && s.IsErrors() {
		loc := IfElse[string](s.Location() == "", "[]", s.Location())
		req := IfElse[string](s.RequestId() == "", "[]", s.RequestId())
		log.Println(req, loc, s.Errors())
		s.RemoveErrors()
	}
	return s
}

// NewErrorHandleFn - templated function providing an error handle function via a closure
func NewErrorHandleFn[E ErrorHandler]() ErrorHandleFn {
	var e E
	return func(location string, errs ...error) *runtime.Status {
		return e.Handle(location, errs...)
	}
}

// NewErrorStatusHandleFn - templated function providing an error status handle function via a closure
func NewErrorStatusHandleFn[E ErrorHandler]() ErrorStatusHandleFn {
	var e E
	return func(s *runtime.Status) *runtime.Status {
		return e.HandleStatus(s)
	}
}

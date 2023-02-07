package template

import (
	"errors"
	"fmt"
	"github.com/idiomatic-go/motif/runtime"
)

func ExampleNoOpErrorHandler_Handle() {
	location := "/test"
	err := errors.New("test error")
	var h NoOpError

	fmt.Printf("test: Handle(location,nil) -> [%v]\n", h.Handle(location, nil))
	fmt.Printf("test: Handle(location,err) -> [%v]\n", h.Handle(location, err))

	s := runtime.NewStatus(runtime.StatusInternal, location, nil)
	fmt.Printf("test: HandleStatus(s) -> [%v]\n", h.HandleStatus(s))

	s = runtime.NewStatus(runtime.StatusInternal, location, err)
	fmt.Printf("test: HandleStatus(s) -> [prev:%v] [curr:%v]\n", s, h.HandleStatus(s))

	//Output:
	//test: Handle(location,nil) -> [OK]
	//test: Handle(location,err) -> [Internal [test error]]
	//test: HandleStatus(s) -> [OK]
	//test: HandleStatus(s) -> [prev:Internal [test error]] [curr:Internal [test error]]

}

func ExampleDebugHandler_Handle() {
	location := "/DebugHandler"
	ctx := runtime.ContextWithRequestId(nil, "123-request-id")
	err := errors.New("test error")
	var h DebugError

	s := h.Handle(location, nil)
	fmt.Printf("test: Handle(location,nil) -> [%v] [errors:%v]\n", s, s.IsErrors())

	s = h.HandleWithContext(ctx, location, err)
	fmt.Printf("test: HandleWithContext(ctx,location,err) -> [%v] [errors:%v]\n", s, s.IsErrors())

	s = runtime.NewStatus(runtime.StatusInternal, location, nil).SetContext(ctx)
	fmt.Printf("test: HandleStatus(s) -> [%v] [errors:%v]\n", h.HandleStatus(s), s.IsErrors())

	s = runtime.NewStatus(runtime.StatusInternal, location, err).SetContext(ctx)
	errors := s.IsErrors()
	s1 := h.HandleStatus(s)
	fmt.Printf("test: HandleStatus(s) -> [prev:%v] [prev-errors:%v] [curr:%v] [curr-errors:%v]\n", s, errors, s1, s1.IsErrors())

	//Output:
	//test: Handle(location,nil) -> [OK] [errors:false]
	//[123-request-id /DebugHandler [test error]]
	//test: HandleWithContext(ctx,location,err) -> [Internal] [errors:false]
	//test: HandleStatus(s) -> [OK] [errors:false]
	//[123-request-id /DebugHandler [test error]]
	//test: HandleStatus(s) -> [prev:Internal] [prev-errors:true] [curr:Internal] [curr-errors:false]

}

func ExampleLogHandler_Handle() {
	location := "/LogHandler"
	ctx := runtime.ContextWithRequestId(nil, "")
	err := errors.New("test error")
	var h LogError

	s := h.Handle(location, nil)
	fmt.Printf("test: Handle(location,nil) -> [%v] [errors:%v]\n", s, s.IsErrors())

	s = h.HandleWithContext(ctx, location, err)
	fmt.Printf("test: HandleWithContext(ctx,location,err) -> [%v] [errors:%v]\n", s, s.IsErrors())

	s = runtime.NewStatus(runtime.StatusInternal, location, nil).SetContext(ctx)
	fmt.Printf("test: HandleStatus(s) -> [%v] [errors:%v]\n", h.HandleStatus(s), s.IsErrors())

	s = runtime.NewStatus(runtime.StatusInternal, location, err).SetContext(ctx)
	errors := s.IsErrors()
	s1 := h.HandleStatus(s)
	fmt.Printf("test: HandleStatus(s) -> [prev:%v] [prev-errors:%v] [curr:%v] [curr-errors:%v]\n", s, errors, s1, s1.IsErrors())

	//Output:
	//test: Handle(location,nil) -> [OK] [errors:false]
	//test: HandleWithContext(ctx,location,err) -> [Internal] [errors:false]
	//test: HandleStatus(s) -> [OK] [errors:false]
	//test: HandleStatus(s) -> [prev:Internal] [prev-errors:true] [curr:Internal] [curr-errors:false]

}

func ExampleErrorHandleFn() {
	loc := pkgPath + "/ErrorHandleFn"
	err := errors.New("debug - error message")

	fn := NewErrorHandleFn[DebugError]()
	fn(loc, err)
	fmt.Printf("test: NewErrorHandleFn[DebugErrorHandler]()\n")

	fn = NewErrorHandleFn[LogError]()
	fn(loc, errors.New("log - error message"))
	fmt.Printf("test: NewErrorHandleFn[LogErrorHandler]()\n")

	//Output:
	//[[] github.com/idiomatic-go/motif/template/ErrorHandleFn [debug - error message]]
	//test: NewErrorHandleFn[DebugErrorHandler]()
	//test: NewErrorHandleFn[LogErrorHandler]()

}

func ExampleErrorHandleStatus() {
	loc := pkgPath + "/ErrorHandleStatus"
	err := errors.New("debug - error message")

	fn := NewErrorStatusHandleFn[DebugError]()
	status := fn(runtime.NewStatusError(loc, err))
	fmt.Printf("test: NewErrorStatusHandleFn[DebugErrorHandler]() [status:%v] [errors:%v]\n", status, status.IsErrors())

	fn = NewErrorStatusHandleFn[LogError]()
	status = fn(runtime.NewStatusError(loc, errors.New("log - error message")))
	fmt.Printf("test: NewErrorStatusHandleFn[LogErrorHandler]() [status:%v] [errors:%v]\n", status, status.IsErrors())

	//Output:
	//[[] github.com/idiomatic-go/motif/template/ErrorHandleStatus [debug - error message]]
	//test: NewErrorStatusHandleFn[DebugErrorHandler]() [status:Internal] [errors:false]
	//test: NewErrorStatusHandleFn[LogErrorHandler]() [status:Internal] [errors:false]

}

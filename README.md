# motif 

Motif was inspired by the challenges of developing Go applications. Determining the patterns, or motifs, that need to be employed, is critical for writing clear idiomatic Go code. This YouTube video [Edward Muller - Go Anti-Patterns][emuller], does an excellent job of framing idiomatic go. 
[Robert Griesemer - The Evolution of Go][rgriesemer] also presents an important analogy between Go packages and LEGO® bricks. Reviewing the Go standard
library packaging structure provides a blueprint for an application architecture, and underscores how essential package design is for idiomatic Go. 

With the release of Go generics, a new paradigm has emerged: [templates][tutorialspoint]. Templates are not new, having been available in  C++ since 1991, and have become a standard through the work of teams like [boost][boost]. I prefer the term templates over generics, as templates are a paradigm, and generics connotes a class of implementations. What follows is a description of the packages in Motif, highlighting specific patterns and template implementations.  



## http
[Http][httppkg] introduces a new design pattern for testing http.Client.Do() calls: DoProxy. A DoProxy is added to a context.Context, and all client requests
are proxied,
~~~
// DoProxy - Http client.Do proxy type
type DoProxy func(req *http.Request) (*http.Response, error)

// ContextWithDo - DoProxy context creation
func ContextWithDo(ctx context.Context, fn DoProxy) context.Context {
	if ctx == nil {
		ctx = context.Background()
	} else {
		if IsContextDo(ctx) {
			return ctx
		}
	}
	if fn == nil {
		return ctx
	}
	return &doContext{ctx, doContextKey, fn} 
}
~~~

Http also includes a common http write response function:
~~~
func WriteResponse(w http.ResponseWriter, buf []byte, status *runtime.Status, headers ...string) {
    // implementation
}
~~~

## messaging
[Messaging][messagingpkg] provides a way for a hosting process to communicate with packages. Packages that register themselves can then be started and pinged by the 
host via the templated functions:
~~~
// Ping - templated function to "ping" a registered resource
func Ping[E template.ErrorHandler](ctx context.Context, uri string) (status *runtime.Status) {
    // Implementation details
}

// Startup - templated function to startup all registered resources.
func Startup[E template.ErrorHandler, O template.OutputHandler](duration time.Duration, content ContentMap) (status *runtime.Status) {
    // Implementation details
}
~~~

## resource
[Resource][resourcepkg] implements configuration map deserialization and validation.

## runtime
[Runtime][runtimepkg] implements environment, request context, and status types. The status type is used extensively as a function return value, and provides error,
http, and gRPC status codes.

## template
[Template][templatepkg] contains template functions for http.Request and http.Response body deserialization, and string expansion:
~~~
// Deserialize - templated function, providing deserialization of a request/response body
func Deserialize[E ErrorHandler, T any](body io.ReadCloser) (T, *runtime.Status) {
    // Implementation details
}
    
 // Resolver - template parameter name value lookup
type Resolver interface {
	Lookup(name string) (string, error)
}

// Expand - templated function to expand a template string, utilizing a resolver
func Expand[T Resolver](t string) (string, error) {   
   // Implementation details
}
~~~

Template parameters for output and error handling are also included:
~~~
// ErrorHandler - template parameter error handler interface
type ErrorHandler interface {
	Handle(location string, errs ...error) *runtime.Status
	HandleWithContext(ctx context.Context, location string, errs ...error) *runtime.Status
	HandleStatus(s *runtime.Status) *runtime.Status
}

// OutputHandler - template parameter output handler interface
type OutputHandler interface {
	Write(s string)
}
~~~

[emuller]: <https://www.youtube.com/watch?v=ltqV6pDKZD8>
[rgriesemer]: <https://www.youtube.com/watch?v=0ReKdcpNyQg>
[tutorialspoint]: <https://www.tutorialspoint.com/cplusplus/cpp_templates.htm>
[boost]: <https://www.boost.org/>
[httppkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/http>
[messagingpkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/messaging>
[resourcepkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/resource>
[runtimepkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/runtime>
[templatepkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/template>


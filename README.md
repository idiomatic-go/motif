# motif 

Motif was inspired by the challenges of developing Go applications. Determining the patterns, or motifs, that need to be employed, is critical for writing clear idiomatic Go code. This YouTube video [Edward Muller - Go Anti-Patterns][emuller], does an excellent job of framing idiomatic go. 
[Robert Griesemer - The Evolution of Go][rgriesemer] also presents an important analogy between Go packages and LEGO® bricks. Reviewing the Go standard
library packaging structure provides a blueprint for an application architecture, and underscores how essential package design is for idiomatic Go. 

With the release of Go generics, a new paradigm has emerged: [templates][tutorialspoint]. Templates are not new, having been available in  C++ since 1991, and have become a standard through the work of teams like [boost][boost]. I prefer the term templates over generics, as templates are a paradigm, and generics connotes a class of implementations. What follows is a description of the packages in Motif, highlighting specific patterns and template implementations.  


## accessdata 

[Accessdata][accessdatapkg] provides the Entry type, which contains all of the data needed for access logging. Also provided are functions and types that define command operators which 
allow the extraction and formatting of Entry data. The formatting of Entry data is implemented as a template parameter: 
~~~
// Formatter - template parameter for formatting
type Formatter interface {
	Format(items []Operator, data *Entry) string
}
~~~
Configurable items, specific to a package, are defined in an options.go file.

## accesslog

[Accesslog][accesslogpkg] encompasses access logging functionality. Seperate operators, and runtime initialization of those operators, are provided for ingress and egress traffic. An output template parameter allows redirection of the access logging: 
~~~
// OutputHandler - template parameter for log output
type OutputHandler interface {
	Write(items []accessdata.Operator, data *accessdata.Entry, formatter accessdata.Formatter)
}
~~~
The log function is a templated function, allowing for selection of output and formatting:
~~~
func Log[O OutputHandler, F accessdata.Formatter](entry *accessdata.Entry) {
    // implementation details
}
~~~

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

## middleware
[Middleware][middlewarepkg]

## runtime
[Runtime][runtimepkg]

## template
[Template][templatepkg]

[emuller]: <https://www.youtube.com/watch?v=ltqV6pDKZD8>
[rgriesemer]: <https://www.youtube.com/watch?v=0ReKdcpNyQg>
[tutorialspoint]: <https://www.tutorialspoint.com/cplusplus/cpp_templates.htm>
[boost]: <https://www.boost.org/>
[accessdatapkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/accessdata>
[accesslogpkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/accesslog>
[httppkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/http>
[messagingpkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/messaging>
[middlewarepkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/middleware>
[runtimepkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/runtime>
[templatepkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/template>
[accessdatapkg]: <https://pkg.go.dev/github.com/idiomatic-go/motif/accessdata>

# motif 

Motif was inspired by the challenges of developing Go applications. Determining the patterns, or motifs, that need to be employed, is critical for writing clear idiomatic Go code. This YouTube video [Edward Muller - Go Anti-Patterns][emuller], does an excellent job of framing idiomatic go. 
[Robert Griesemer - The Evolution of Go][rgriesemer] also presents an important analogy between Go packages and LEGO® bricks. Reviewing the Go standard
library packaging structure provides a blueprint for an application architecture, and underscores how essential package design is for idiomatic Go. 

With the release of Go generics, a new paradigm has emerged: [templates][tutorialspoint]. Templates are not new, having been available in  C++ since 1991, and have become a standard through the work of teams like [boost][boost]. I prefer the term templates over generics, as templates are a paradigm, and generics connotes a class of implementations. What follows is a description of the packages in Motif, highlighting specific patterns and template implementations.  


## accessdata

Accessdata provides the Entry type, which contains all of the data needed for access logging.
~~~
// Entry - struct for all access logging data
type Entry struct {
	Traffic  string
	Start    time.Time
	Duration time.Duration
	Origin   *Origin
	ActState map[string]string

	// Request
	Url       string
	Path      string
	Host      string
	Protocol  string
	Method    string
	Header    http.Header
	RequestId string

	// Response
	StatusCode    int
	BytesSent     int64 // ingress response
	BytesReceived int64 // handler response content length
	StatusFlags   string
}

~~~
Also provided are functions and types that define command operators which allow the extraction and formatting of Entry data. The formatting of Entry data is implemented as a template parameter: 
~~~
// Formatter - template parameter for formatting
type Formatter interface {
	Format(items []Operator, data *Entry) string
}

~~~

[emuller]: <https://www.youtube.com/watch?v=ltqV6pDKZD8>
[rgriesemer]: <https://www.youtube.com/watch?v=0ReKdcpNyQg>
[tutorialspoint]: <https://www.tutorialspoint.com/cplusplus/cpp_templates.htm>
[boost]: <https://www.boost.org/>

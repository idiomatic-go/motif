package accessdata

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	EgressTraffic      = "egress"
	IngressTraffic     = "ingress"
	PingTraffic        = "ping"
	TimeoutName        = "timeout"
	FailoverName       = "failover"
	RetryName          = "retry"
	RetryRateLimitName = "retryRateLimit"
	RetryRateBurstName = "retryBurst"
	RateLimitName      = "rateLimit"
	RateBurstName      = "burst"
	ActName            = "name"
)

// Origin - attributes that uniquely identify a service instance
type Origin struct {
	Region     string
	Zone       string
	Service    string
	InstanceId string
}

// Entry - struct for all access logging information
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

// NewEntry - create a new Entry
func NewEntry() *Entry {
	return new(Entry)
}

func newEntry(traffic string, start time.Time, duration time.Duration, actState map[string]string, req *http.Request, resp *http.Response, statusFlags string) *Entry {
	l := new(Entry)
	l.Traffic = traffic
	l.Start = start
	l.Duration = duration
	l.Origin = &opt.origin
	if actState == nil {
		actState = make(map[string]string, 1)
	}
	l.ActState = actState
	l.AddRequest(req)
	l.AddResponse(resp)
	l.StatusFlags = statusFlags
	return l
}

// NewHttpIngressEntry - create an Entry from Http ingress traffic
func NewHttpIngressEntry(start time.Time, duration time.Duration, actState map[string]string, req *http.Request, statusCode int, written int64, statusFlags string) *Entry {
	e := newEntry(IngressTraffic, start, duration, actState, req, nil, statusFlags)
	e.StatusCode = statusCode
	e.BytesSent = written
	return e
}

// NewHttpEgressEntry - create an Entry from Http egress traffic
func NewHttpEgressEntry(start time.Time, duration time.Duration, actState map[string]string, req *http.Request, resp *http.Response, statusFlags string) *Entry {
	return newEntry(EgressTraffic, start, duration, actState, req, resp, statusFlags)
}

// NewEgressEntry - create an Entry from non-http egress traffic
func NewEgressEntry(start time.Time, duration time.Duration, actState map[string]string, statusCode int, uri, requestId, method, statusFlags string) *Entry {
	e := newEntry(EgressTraffic, start, duration, actState, nil, nil, statusFlags)
	e.StatusCode = statusCode
	e.AddUrl(uri)
	e.RequestId = requestId
	e.Method = method
	return e
}

func (l *Entry) IsIngress() bool {
	return l.Traffic == IngressTraffic
}

func (l *Entry) IsEgress() bool {
	return l.Traffic == EgressTraffic
}

func (l *Entry) IsPing() bool {
	return IsPingRoute(l.Traffic, l.Path)
}

func (l *Entry) AddResponse(resp *http.Response) {
	if resp == nil {
		return
	}
	l.StatusCode = resp.StatusCode
	l.BytesReceived = resp.ContentLength
}

func (l *Entry) AddUrl(uri string) {
	if uri == "" {
		return
	}
	u, err := url.Parse(uri)
	if err != nil {
		l.Url = err.Error()
		return
	}
	if u.Scheme == "urn" && u.Host == "" {
		l.Url = uri
		l.Protocol = u.Scheme
		t := strings.Split(u.Opaque, ":")
		if len(t) == 1 {
			l.Host = t[0]
		} else {
			l.Host = t[0]
			l.Path = t[1]
		}
	} else {
		l.Protocol = u.Scheme
		l.Url = u.String()
		l.Path = u.Path
		l.Host = u.Host
	}
}

func (l *Entry) AddRequest(req *http.Request) {
	if req == nil {
		return
	}
	l.Protocol = req.Proto
	l.Method = req.Method
	if req.Header != nil {
		l.Header = req.Header.Clone()
	}
	if req.URL != nil {
		l.Url = req.URL.String()
		l.Path = req.URL.Path
		if req.Host == "" {
			l.Host = req.URL.Host
		} else {
			l.Host = req.Host
		}
	}
}

func (l *Entry) Value(value string) string {
	switch value {
	case TrafficOperator:
		if l.IsPing() {
			return PingTraffic
		}
		return l.Traffic
	case StartTimeOperator:
		return FmtTimestamp(l.Start)
	case DurationOperator:
		d := int(l.Duration / time.Duration(1e6))
		return strconv.Itoa(d)
	case DurationStringOperator:
		return l.Duration.String()

		// Origin
	case OriginRegionOperator:
		if l.Origin != nil {
			return l.Origin.Region
		}
		//return opt.origin.Region
	case OriginZoneOperator:
		if l.Origin != nil {
			return l.Origin.Zone
		}
		//return opt.origin.Zone
	//case OriginSubZoneOperator:
	//	if l.Origin != nil {
	//		return l.Origin.SubZone
	//	}
	//return opt.origin.SubZone
	case OriginServiceOperator:
		if l.Origin != nil {
			return l.Origin.Service
		}
		//return opt.origin.Service
	case OriginInstanceIdOperator:
		if l.Origin != nil {
			return l.Origin.InstanceId
		}
		//return opt.origin.InstanceId

		// Request
	case RequestMethodOperator:
		return l.Method
	case RequestProtocolOperator:
		return l.Protocol
	case RequestPathOperator:
		return l.Path
	case RequestUrlOperator:
		return l.Url
	case RequestHostOperator:
		return l.Host
	case RequestIdOperator:
		if l.RequestId != "" {
			return l.RequestId
		}
		return l.Header.Get(RequestIdHeaderName)
	case RequestFromRouteOperator:
		return l.Header.Get(FromRouteHeaderName)
	case RequestUserAgentOperator:
		return l.Header.Get(UserAgentHeaderName)
	case RequestAuthorityOperator:
		return ""
	case RequestForwardedForOperator:
		return l.Header.Get(FordwardedForHeaderName)

		// Response
	case StatusFlagsOperator:
		return l.StatusFlags
	case ResponseBytesReceivedOperator:
		return strconv.Itoa(int(l.BytesReceived))
	case ResponseBytesSentOperator:
		return fmt.Sprintf("%v", l.BytesSent)
	case ResponseStatusCodeOperator:
		return strconv.Itoa(l.StatusCode)

	// Actuator State
	case RouteNameOperator:
		return l.ActState[ActName]
	case TimeoutDurationOperator:
		return l.ActState[TimeoutName]
	case RateLimitOperator:
		return l.ActState[RateLimitName]
	case RateBurstOperator:
		return l.ActState[RateBurstName]
	case FailoverOperator:
		return l.ActState[FailoverName]
	case RetryOperator:
		return l.ActState[RetryName]
	case RetryRateLimitOperator:
		return l.ActState[RetryRateLimitName]
	case RetryRateBurstOperator:
		return l.ActState[RetryRateBurstName]
	}
	if strings.HasPrefix(value, RequestReferencePrefix) {
		name := requestOperatorHeaderName(value)
		return l.Header.Get(name)
	}
	if !strings.HasPrefix(value, OperatorPrefix) {
		return value
	}

	return ""
}

func (l *Entry) String() string {
	return fmt.Sprintf( //"start:%v ,"+
		//"duration:%v ,"+
		"traffic:%v, "+
			"route:%v, "+
			"request-id:%v, "+
			"status-code:%v, "+
			"method:%v, "+
			"url:%v, "+
			"host:%v, "+
			"path:%v, "+
			"timeout:%v, "+
			"rate-limit:%v, "+
			"rate-burst:%v, "+
			"retry:%v, "+
			"retry-rate-limit:%v, "+
			"retry-rate-burst:%v, "+
			"status-flags:%v",
		//l.Value(StartTimeOperator),
		//l.Value(DurationOperator),
		l.Value(TrafficOperator),
		l.Value(RouteNameOperator),

		l.Value(RequestIdOperator),
		l.Value(ResponseStatusCodeOperator),
		l.Value(RequestMethodOperator),
		l.Value(RequestUrlOperator),
		l.Value(RequestHostOperator),
		l.Value(RequestPathOperator),

		l.Value(TimeoutDurationOperator),
		l.Value(RateLimitOperator),
		l.Value(RateBurstOperator),

		l.Value(RetryOperator),
		l.Value(RetryRateLimitOperator),
		l.Value(RetryRateBurstOperator),

		l.Value(StatusFlagsOperator),
	)
}

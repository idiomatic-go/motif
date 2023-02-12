package accessdata

import (
	"errors"
	"fmt"
	"strings"
)

const (
	OperatorPrefix         = "%"
	RequestReferencePrefix = "%REQ("

	RequestIdHeaderName     = "X-REQUEST-ID"
	FromRouteHeaderName     = "FROM-ROUTE"
	UserAgentHeaderName     = "USER-AGENT"
	FordwardedForHeaderName = "X-FORWARDED-FOR"

	TrafficOperator        = "%TRAFFIC%"      //  ingress, handler, ping
	StartTimeOperator      = "%START_TIME%"   // start time
	DurationOperator       = "%DURATION%"     // Total duration in milliseconds of the request from the start time to the last byte out.
	DurationStringOperator = "%DURATION_STR%" // Time package formatted

	OriginRegionOperator     = "%REGION%"      // origin region
	OriginZoneOperator       = "%ZONE%"        // origin zone
	OriginSubZoneOperator    = "%SUB_ZONE%"    // origin sub zone
	OriginServiceOperator    = "%SERVICE%"     // origin service
	OriginInstanceIdOperator = "%INSTANCE_ID%" // origin instance id

	RouteNameOperator       = "%ROUTE_NAME%"
	TimeoutDurationOperator = "%TIMEOUT_DURATION%"
	RateLimitOperator       = "%RATE_LIMIT%"
	RateBurstOperator       = "%RATE_BURST%"
	RetryOperator           = "%RETRY"
	RetryRateLimitOperator  = "%RETRY_RATE_LIMIT%"
	RetryRateBurstOperator  = "%RETRY_RATE_BURST%"
	FailoverOperator        = "%FAILOVER%"

	ResponseStatusCodeOperator    = "%STATUS_CODE%"    // HTTP status code
	ResponseBytesReceivedOperator = "%BYTES_RECEIVED%" // bytes received
	ResponseBytesSentOperator     = "%BYTES_SENT%"     // bytes sent
	StatusFlagsOperator           = "%STATUS_FLAGS%"   // status flags
	//UpstreamHostOperator  = "%UPSTREAM_HOST%"  // Upstream host URL (e.g., tcp://ip:port for TCP connections).

	RequestProtocolOperator = "%PROTOCOL%" // HTTP Protocol
	RequestMethodOperator   = "%METHOD%"   // HTTP method
	RequestUrlOperator      = "%URL%"
	RequestPathOperator     = "%PATH%"
	RequestHostOperator     = "%HOST%"

	RequestIdOperator           = "%X-REQUEST-ID%"    // X-REQUEST-ID request header value
	RequestFromRouteOperator    = "%FROM-ROUTE%"      // request from route name
	RequestUserAgentOperator    = "%USER-AGENT%"      // user agent request header value
	RequestAuthorityOperator    = "%AUTHORITY%"       // authority request header value
	RequestForwardedForOperator = "%X-FORWARDED-FOR%" // client IP address (X-FORWARDED-FOR request header value)

	GRPCStatusOperator       = "%GRPC_STATUS(X)%"     // gRPC status code formatted according to the optional parameter X, which can be CAMEL_STRING, SNAKE_STRING and NUMBER. X-REQUEST-ID request header value
	GRPCStatusNumberOperator = "%GRPC_STATUS_NUMBER%" // gRPC status code.

)

// Operator - configuration of logging entries
type Operator struct {
	Name  string
	Value string
}

func IsDirectOperator(op Operator) bool {
	return !strings.HasPrefix(op.Value, OperatorPrefix)
}

func IsRequestOperator(op Operator) bool {
	if !strings.HasPrefix(op.Value, RequestReferencePrefix) {
		return false
	}
	if len(op.Value) < (len(RequestReferencePrefix) + 2) {
		return false
	}
	return op.Value[len(op.Value)-2:] == ")%"
}

func RequestOperatorHeaderName(op Operator) string {
	if op.Name != "" {
		return op.Name
	}
	return requestOperatorHeaderName(op.Value)
}

func requestOperatorHeaderName(value string) string {
	if len(value) < (len(RequestReferencePrefix) + 2) {
		return ""
	}
	return value[len(RequestReferencePrefix) : len(value)-2]
}

func IsStringValue(op Operator) bool {
	switch op.Value {
	case DurationOperator, TimeoutDurationOperator, RateBurstOperator,
		RateLimitOperator, RetryOperator, RetryRateLimitOperator, RetryRateBurstOperator,
		FailoverOperator, ResponseStatusCodeOperator,
		ResponseBytesSentOperator, ResponseBytesReceivedOperator:
		return false
	}
	return true
}

var Operators = map[string]*Operator{
	TrafficOperator:        {"traffic", TrafficOperator},
	StartTimeOperator:      {"start_time", StartTimeOperator},
	DurationOperator:       {"duration_ms", DurationOperator},
	DurationStringOperator: {"duration", DurationStringOperator},

	OriginRegionOperator:     {"region", OriginRegionOperator},
	OriginZoneOperator:       {"zone", OriginZoneOperator},
	OriginSubZoneOperator:    {"sub_zone", OriginSubZoneOperator},
	OriginServiceOperator:    {"service", OriginServiceOperator},
	OriginInstanceIdOperator: {"instance_id", OriginInstanceIdOperator},

	// Route
	RouteNameOperator:       {"route_name", RouteNameOperator},
	TimeoutDurationOperator: {"timeout_ms", TimeoutDurationOperator},
	RateLimitOperator:       {"rate_limit", RateLimitOperator},
	RateBurstOperator:       {"rate_burst", RateBurstOperator},
	RetryOperator:           {"retry", RetryOperator},
	RetryRateLimitOperator:  {"retry_rate_limit", RetryRateLimitOperator},
	RetryRateBurstOperator:  {"retry_rate_burst", RetryRateBurstOperator},
	FailoverOperator:        {"failover", FailoverOperator},

	// Response
	ResponseStatusCodeOperator:    {"status_code", ResponseStatusCodeOperator},
	ResponseBytesReceivedOperator: {"bytes_received", ResponseBytesReceivedOperator},
	ResponseBytesSentOperator:     {"bytes_sent", ResponseBytesSentOperator},
	StatusFlagsOperator:           {"status_flags", StatusFlagsOperator},
	//UpstreamHostOperator:  {"upstream_host", UpstreamHostOperator},

	// Request
	RequestProtocolOperator: {"protocol", RequestProtocolOperator},
	RequestUrlOperator:      {"url", RequestUrlOperator},
	RequestMethodOperator:   {"method", RequestMethodOperator},
	RequestPathOperator:     {"path", RequestPathOperator},
	RequestHostOperator:     {"host", RequestHostOperator},

	RequestIdOperator:           {"request_id", RequestIdOperator},
	RequestFromRouteOperator:    {"request_id", RequestIdOperator},
	RequestUserAgentOperator:    {"user_agent", RequestUserAgentOperator},
	RequestAuthorityOperator:    {"authority", RequestAuthorityOperator},
	RequestForwardedForOperator: {"forwarded", RequestForwardedForOperator},

	// gRPC
	GRPCStatusOperator:       {"grpc_status", GRPCStatusOperator},
	GRPCStatusNumberOperator: {"grpc_number", GRPCStatusNumberOperator},
}

func CreateOperators(operators []string) ([]Operator, error) {
	var items []Operator
	for _, op := range operators {
		items = append(items, Operator{
			Name:  "",
			Value: op,
		})
	}
	return InitOperators(items)
}

func InitOperators(operators []Operator) ([]Operator, error) {
	var items []Operator

	//if items == nil {
	//	return nil,errors.New("invalid configuration: operators slice is nil")
	//}
	if len(operators) == 0 {
		return nil, errors.New("invalid configuration: configuration slice is empty")
	}
	dup := make(map[string]string)
	for _, op := range operators {
		op2, err := createOperator(op)
		if err != nil {
			return nil, err
		}
		if _, ok := dup[op2.Name]; ok {
			return nil, errors.New(fmt.Sprintf("invalid operator: name is a duplicate [%v]", op2.Name))
		}
		dup[op2.Name] = op2.Name
		items = append(items, op2)
	}
	return items, nil
}

func createOperator(op Operator) (Operator, error) {
	if IsEmpty(op.Value) {
		return Operator{}, errors.New(fmt.Sprintf("invalid operator: value is empty %v", op.Name))
	}
	if IsDirectOperator(op) {
		if IsEmpty(op.Name) {
			return Operator{}, errors.New(fmt.Sprintf("invalid operator: name is empty [%v]", op.Value))
		}
		return Operator{Name: op.Name, Value: op.Value}, nil
	}
	if op2, ok := Operators[op.Value]; ok {
		newOp := Operator{Name: op2.Name, Value: op.Value}
		if !IsEmpty(op.Name) {
			newOp.Name = op.Name
		}
		return newOp, nil
	}
	if IsRequestOperator(op) {
		return Operator{Name: RequestOperatorHeaderName(op), Value: op.Value}, nil
	}
	return Operator{}, errors.New(fmt.Sprintf("invalid operator: value not found or invalid %v", op.Value))
}

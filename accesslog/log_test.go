package accesslog

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
	"net/http"
	"reflect"
	"time"
)

func ExampleLog_Error() {
	start := time.Now()

	Log[TestOutputHandler](nil)
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start, time.Since(start), map[string]string{accessdata.ActName: "egress-route"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"error":"access data entry is nil"}]
	//test: Write() -> [{"error":"egress log entries are empty"}]

}

func ExampleLog_Origin() {
	name := "ingress-origin-route"
	start := time.Now()

	accessdata.SetOrigin(accessdata.Origin{Region: "us-west", Zone: "dfw", SubZone: "cluster", Service: "test-service", InstanceId: "123456-7890-1234"})
	err := InitIngressOperators([]accessdata.Operator{{Value: accessdata.StartTimeOperator}, {Value: accessdata.DurationOperator, Name: "duration_ms"},
		{Value: accessdata.TrafficOperator}, {Value: accessdata.RouteNameOperator}, {Value: accessdata.OriginRegionOperator}, {Value: accessdata.OriginZoneOperator}, {Value: accessdata.OriginSubZoneOperator}, {Value: accessdata.OriginServiceOperator}, {Value: accessdata.OriginInstanceIdOperator},
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Log[TestOutputHandler](accessdata.NewHttpIngressEntry(start1, time.Since(start), map[string]string{accessdata.ActName: name}, nil, 0, 0, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":"ingress","route_name":"ingress-origin-route","region":"us-west","zone":"dfw","sub_zone":"cluster","service":"test-service","instance_id":"123456-7890-1234"}]

}

func _ExampleLog_Ping() {
	name := "ingress-ping-route"

	accessdata.SetPingRoutes([]accessdata.PingRoute{{Traffic: "ingress", Pattern: name}})
	start := time.Now()
	err := InitIngressOperators([]accessdata.Operator{{Value: accessdata.StartTimeOperator}, {Value: accessdata.DurationOperator, Name: "duration_ms"},
		{Value: accessdata.TrafficOperator}, {Value: accessdata.RouteNameOperator}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Log[TestOutputHandler](accessdata.NewHttpIngressEntry(start1, time.Since(start), map[string]string{accessdata.ActName: name}, nil, 0, 0, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":"ping","route_name":"ingress-ping-route"}]

}

func ExampleLog_Timeout() {
	start := time.Now()

	err := InitEgressOperators([]accessdata.Operator{{Value: accessdata.StartTimeOperator}, {Name: "duration_ms", Value: accessdata.DurationOperator},
		{Value: accessdata.TrafficOperator}, {Value: accessdata.RouteNameOperator}, {Value: accessdata.TimeoutDurationOperator}, {Name: "static", Value: "value"}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start1, time.Since(start), map[string]string{accessdata.ActName: "handler-route", accessdata.TimeoutName: "5000"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":"egress","route_name":"handler-route","timeout_ms":5000,"static":"value"}]

}

func ExampleLog_RateLimiter_500() {
	start := time.Now()

	err := InitEgressOperators([]accessdata.Operator{{Value: accessdata.StartTimeOperator}, {Name: "duration", Value: accessdata.DurationOperator},
		{Value: accessdata.TrafficOperator}, {Value: accessdata.RouteNameOperator}, {Value: accessdata.RateLimitOperator}, {Value: accessdata.RateBurstOperator}, {Name: "static2", Value: "value2"}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start1, time.Since(start), map[string]string{accessdata.ActName: "handler-route", accessdata.RateLimitName: "500", accessdata.RateBurstName: "10"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration":0,"traffic":"egress","route_name":"handler-route","rate_limit":500,"rate_burst":10,"static2":"value2"}]

}

func ExampleLog_Failover() {
	start := time.Now()

	err := InitEgressOperators([]accessdata.Operator{{Value: accessdata.StartTimeOperator}, {Name: "duration", Value: accessdata.DurationOperator},
		{Value: accessdata.TrafficOperator}, {Value: accessdata.RouteNameOperator}, {Value: accessdata.FailoverOperator}, {Name: "static2", Value: "value2"}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start1, time.Since(start), map[string]string{accessdata.ActName: "handler-route", accessdata.FailoverName: "true"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration":0,"traffic":"egress","route_name":"handler-route","failover":true,"static2":"value2"}]

}

func ExampleLog_Retry() {
	start := time.Now()

	err := InitEgressOperators([]accessdata.Operator{{Value: accessdata.StartTimeOperator}, {Value: accessdata.DurationOperator, Name: "duration_ms"},
		{Value: accessdata.TrafficOperator}, {Value: accessdata.RouteNameOperator}, {Value: accessdata.RetryOperator},
		{Value: accessdata.RetryRateLimitOperator}, {Value: accessdata.RetryRateBurstOperator}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start1 time.Time
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start1, time.Since(start), map[string]string{accessdata.ActName: "handler-route", accessdata.RetryName: "true", accessdata.RetryRateLimitName: "123", accessdata.RetryRateBurstName: "67"}, nil, nil, ""))

	//Output:
	//test: Write() -> [{"start_time":"0001-01-01 00:00:00.000000","duration_ms":0,"traffic":"egress","route_name":"handler-route","retry":true,"retry_rate_limit":123,"retry_rate_burst":67}]

}

func ExampleLog_Request() {
	req, _ := http.NewRequest("", "www.google.com/search/documents", nil)
	req.Header.Add("customer", "Ted's Bait & Tackle")

	var start time.Time
	err := InitEgressOperators([]accessdata.Operator{{Value: accessdata.RequestProtocolOperator}, {Value: accessdata.RequestMethodOperator}, {Value: accessdata.RequestUrlOperator},
		{Value: accessdata.RequestPathOperator}, {Value: accessdata.RequestHostOperator}, {Value: "%REQ(customer)%"}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start, time.Since(start), map[string]string{accessdata.ActName: "handler-route"}, nil, nil, ""))
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start, time.Since(start), map[string]string{accessdata.ActName: "handler-route"}, req, nil, ""))

	//Output:
	//test: Write() -> [{"protocol":null,"method":null,"url":null,"path":null,"host":null,"customer":null}]
	//test: Write() -> [{"protocol":"HTTP/1.1","method":"GET","url":"www.google.com/search/documents","path":"www.google.com/search/documents","host":null,"customer":"Ted's Bait & Tackle"}]

}

func ExampleLog_Response() {
	resp := &http.Response{StatusCode: 404, ContentLength: 1234}

	err := InitEgressOperators([]accessdata.Operator{{Value: accessdata.ResponseStatusCodeOperator}, {Value: accessdata.ResponseBytesReceivedOperator}, {Value: accessdata.StatusFlagsOperator}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	var start time.Time
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start, time.Since(start), map[string]string{accessdata.ActName: "handler-route"}, nil, nil, "UT"))
	Log[TestOutputHandler](accessdata.NewHttpEgressEntry(start, time.Since(start), map[string]string{accessdata.ActName: "handler-route"}, nil, resp, "UT"))

	//Output:
	//test: Write() -> [{"status_code":0,"bytes_received":0,"status_flags":"UT"}]
	//test: Write() -> [{"status_code":404,"bytes_received":1234,"status_flags":"UT"}]

}

func _Example_Log_State() {
	t := time.Duration(time.Millisecond * 500)
	i := reflect.TypeOf(t)
	a := any(t)

	fmt.Printf("test 1 -> %v\n", a)

	fmt.Printf("test 2 -> %v\n", i)

	//Output:
	//fail
}

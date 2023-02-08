package middleware

import (
	"fmt"
	"github.com/idiomatic-go/motif/accessdata"
	"github.com/idiomatic-go/motif/accesslog"
	"net/http"
)

var (
	isEnabled    = false
	googleUrl    = "https://www.google.com/search?q=test"
	instagramUrl = "https://www.instagram.com"

	config = []accessdata.Operator{
		//{Value: accessdata.StartTimeOperator},
		//{Value: accessdata.DurationOperator},
		{Value: accessdata.TrafficOperator},
		{Value: accessdata.RouteNameOperator},

		{Value: accessdata.RequestMethodOperator},
		{Value: accessdata.RequestHostOperator},
		{Value: accessdata.RequestPathOperator},
		{Value: accessdata.RequestProtocolOperator},

		{Value: accessdata.ResponseStatusCodeOperator},
		{Value: accessdata.StatusFlagsOperator},
		{Value: accessdata.ResponseBytesReceivedOperator},
		{Value: accessdata.ResponseBytesSentOperator},

		/*
			{Value: accessdata.TimeoutDurationOperator},
			{Value: accessdata.RateLimitOperator},
			{Value: accessdata.RateBurstOperator},
			{Value: accessdata.RetryOperator},
			{Value: accessdata.RetryRateLimitOperator},
			{Value: accessdata.RetryRateBurstOperator},
			{Value: accessdata.FailoverOperator},

		*/
	}
)

func init() {
	err := accesslog.InitEgressOperators(config)
	if err != nil {
		fmt.Printf("init() -> [:%v]\n", err)
	}

	SetLogFn(func(entry *accessdata.Entry) {
		accesslog.Log[accesslog.TestOutputHandler, accessdata.JsonFormatter](entry)
	},
	)
}

func Example_No_Wrapper() {
	req, _ := http.NewRequest("GET", googleUrl, nil)

	// Testing - check for a nil wrapper or round tripper
	w := wrapper{}
	resp, err := w.RoundTrip(req)
	fmt.Printf("test: RoundTrip(wrapper:nil) -> [resp:%v] [err:%v]\n", resp, err)

	// Testing - no wrapper, calling Google search
	resp, err = http.DefaultClient.Do(req)
	fmt.Printf("test: RoundTrip(handler:false) -> [status_code:%v] [err:%v]\n", resp.StatusCode, err)

	//Output:
	//test: RoundTrip(wrapper:nil) -> [resp:<nil>] [err:invalid handler round tripper configuration : http.RoundTripper is nil]
	//test: RoundTrip(handler:false) -> [status_code:200] [err:<nil>]

}

func Example_Default() {
	req, _ := http.NewRequest("GET", instagramUrl, nil)

	if !isEnabled {
		isEnabled = true
		WrapDefaultTransport()
	}
	resp, err := http.DefaultClient.Do(req)
	fmt.Printf("test: RoundTrip(handler:true) -> [status_code:%v] [err:%v]\n", resp.StatusCode, err)

	//Output:
	//test: Write() -> [{"traffic":"egress","route_name":null,"method":"GET","host":"www.instagram.com","path":null,"protocol":"HTTP/1.1","status_code":200,"status_flags":null,"bytes_received":-1,"bytes_sent":0}]
	//test: RoundTrip(handler:true) -> [status_code:200] [err:<nil>]

}

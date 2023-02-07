package accessdata

import (
	"fmt"
	"net/http"
	"time"
)

func Example_Value_Origin() {
	op := OriginRegionOperator
	SetOrigin(Origin{"region", "zone", "", ""})
	data := Entry{Origin: &opt.origin}
	fmt.Printf("test: Value(\"region\") -> [%v]\n", data.Value(op))

	//Output:
	//test: Value("region") -> [region]
}

func _Example_Value_Duration() {
	start := time.Now()

	time.Sleep(time.Second * 2)
	data := &Entry{}
	data.Duration = time.Since(start)
	fmt.Printf("test: Value(\"Duration\") -> [%v]\n", data.Value(DurationOperator))
	fmt.Printf("test: Value(\"DurationString\") -> [%v]\n", data.Value(DurationStringOperator))

	//Output:
	//test: Value("Duration") -> [2011]
	//test: Value("DurationString") -> [2.0117949s]

}

func Example_Value_Actuator() {
	name := "test-route"
	op := RouteNameOperator

	data := Entry{}
	fmt.Printf("test: Value(\"%v\") -> [%v]\n", name, data.Value(op))

	data = Entry{ActState: map[string]string{ActName: name}}
	fmt.Printf("test: Value(\"%v\") -> [route_name:%v]\n", name, data.Value(op))

	data = Entry{ActState: map[string]string{TimeoutName: "500"}}
	fmt.Printf("test: Value(\"%v\") -> [timeout:%v]\n", name, data.Value(TimeoutDurationOperator))

	//Output:
	//test: Value("test-route") -> []
	//test: Value("test-route") -> [route_name:test-route]
	//test: Value("test-route") -> [timeout:500]
}

func Example_Value_Request() {
	op := RequestMethodOperator

	data := &Entry{}
	fmt.Printf("test: Value(\"method\") -> [%v]\n", data.Value(op))

	req, _ := http.NewRequest("POST", "www.google.com", nil)
	//req.Header.Add(RequestIdHeaderName, uuid.New().String())
	req.Header.Add(RequestIdHeaderName, "123-456-789")
	req.Header.Add(FromRouteHeaderName, "calling-route")
	data = &Entry{}
	data.addRequest(req)
	fmt.Printf("test: Value(\"method\") -> [%v]\n", data.Value(op))

	fmt.Printf("test: Value(\"headers\") -> [request-id:%v] [from-route:%v]\n", data.Value(RequestIdOperator), data.Value(RequestFromRouteOperator))

	//Output:
	//test: Value("method") -> []
	//test: Value("method") -> [POST]
	//test: Value("headers") -> [request-id:123-456-789] [from-route:calling-route]
}

func Example_Value_Response() {
	op := ResponseStatusCodeOperator

	data := &Entry{}
	fmt.Printf("test: Value(\"code\") -> [%v]\n", data.Value(op))

	resp := &http.Response{StatusCode: 200}
	data = &Entry{}
	data.addResponse(resp)
	fmt.Printf("test: Value(\"code\") -> [%v]\n", data.Value(op))

	//Output:
	//test: Value("code") -> [0]
	//test: Value("code") -> [200]
}

func Example_Value_Request_Header() {
	req, _ := http.NewRequest("", "www.google.com", nil)
	req.Header.Add("customer", "Ted's Bait & Tackle")
	data := Entry{}
	data.addRequest(req)
	fmt.Printf("test: Value(\"customer\") -> [%v]\n", data.Value("%REQ(customer)%"))

	//Output:
	//test: Value("customer") -> [Ted's Bait & Tackle]
}

func Example_Entry() {
	var start time.Time
	e := NewEgressEntry(start, 0, nil, 200, "urn:postgres:query.access-log", "123-456-789", "GET", "RL")
	fmt.Printf("test: String() -> {%v}\n", e)

	//Output:
	//test: String() -> {start:0001-01-01 00:00:00.000000 ,duration:0 ,traffic:egress ,route: ,request-id:123-456-789 ,status-code:200 ,method:GET ,host:postgres ,path:query.access-log ,timeout: ,rate-limit: ,rate-burst: ,retry: ,retry-rate-limit: ,retry-rate-burst: ,status-flags:RL}

}

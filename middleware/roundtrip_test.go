package middleware

import (
	"fmt"
	"net/http"
)

var (
	isEnabled    = false
	googleUrl    = "https://www.google.com/search?q=test"
	instagramUrl = "https://www.instagram.com"
)

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

func _Example_Default() {
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

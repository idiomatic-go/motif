package runtime

import (
	"errors"
	"fmt"
	"net/http"
)

func ExampleStatus_String() {
	s := NewStatus(StatusPermissionDenied, "", nil)
	fmt.Printf("test: NewStatus() -> [%v]\n", s)

	s = NewStatus(StatusOutOfRange, "", errors.New("error - 1"), errors.New("error - 2"))
	fmt.Printf("test: NewStatus() -> [%v]\n", s)

	//Output:
	//test: NewStatus() -> [OK]
	//test: NewStatus() -> [OutOfRange [error - 1 error - 2]]

}

func ExampleStatus_Http() {
	location := "test"
	err := errors.New("http error")
	fmt.Printf("test: NewHttpStatus(nil) -> [%v]\n", NewHttpStatus(nil, location, nil))
	fmt.Printf("test: NewHttpStatus(nil) -> [%v]\n", NewHttpStatus(nil, location, err))

	resp := http.Response{StatusCode: http.StatusBadRequest}
	fmt.Printf("test: NewHttpStatus(resp) -> [%v]\n", NewHttpStatus(&resp, location, nil))
	fmt.Printf("test: NewHttpStatus(resp) -> [%v]\n", NewHttpStatus(&resp, location, err))

	//Output:
	//test: NewHttpStatus(nil) -> [Invalid Content]
	//test: NewHttpStatus(nil) -> [Internal Error [http error]]
	//test: NewHttpStatus(resp) -> [Bad Request]
	//test: NewHttpStatus(resp) -> [Internal Error [http error]]

}

func ExampleStatus_SetMetadata() {
	s := NewStatusOK()

	s.SetMetadata("content-type", "text/plain")
	fmt.Printf("test: SetMetadata() -> %v\n", s.md)

	s = NewStatusOK()
	resp := &http.Response{}
	resp.Header = make(http.Header)
	resp.Header.Add("content-length", "1234")
	resp.Header.Add("host", "www.google.com")
	s.SetMetadataFromResponse(resp, "host", "content-length")
	fmt.Printf("test: SetMetadata() -> %v\n", s.md)

	//Output:
	//test: SetMetadata() -> map[content-type:[text/plain]]
	//test: SetMetadata() -> map[content-length:[1234] host:[www.google.com]]

}

func ExampleStatus_AddMetadata() {
	s := NewStatusOK()

	s.SetMetadata("content-type", "text/plain", "content-length", "1234", "host", "www.google.com")
	h := make(http.Header)
	s.AddMetadata(h, "content-length", "host")
	fmt.Printf("test: AddMetadata(h) -> %v\n", h)

	//Output:
	//test: AddMetadata(h) -> map[Content-Length:[1234] Host:[www.google.com]]

}

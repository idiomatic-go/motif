package runtime

import (
	"errors"
	"fmt"
	"net/http"
)

type address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a address) GetZip() string {
	return a.Zip
}

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

	s.SetMetadata("content-type", "text/plain", "content-length")
	fmt.Printf("test: SetMetadata() -> %v\n", s.md)

	s = NewStatusOK()
	resp := &http.Response{}
	resp.Header = make(http.Header)
	resp.Header.Add("content-length", "1234")
	resp.Header.Add("host", "www.google.com")
	s.SetMetadataFromResponse(resp, "host", "content-length")
	fmt.Printf("test: SetMetadata() -> %v\n", s.md)

	//Output:
	//test: SetMetadata() -> map[content-length:[] content-type:[text/plain]]
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

func ExampleStatus_Content() {
	str := "test string content"
	s := NewStatusOK()

	s.SetContent(str)
	fmt.Printf("test: SetContent(%v) -> [content:%v] [type:%v]\n", str, string(s.Content()), s.MetadataValue(ContentType))

	s.RemoveContent()
	s.SetContent([]byte(str))
	fmt.Printf("test: SetContent(%v) -> [content:%v] [type:%v]\n", str, string(s.Content()), s.MetadataValue(ContentType))

	s.RemoveContent()
	s.SetContent(12345)
	fmt.Printf("test: SetContent(12345) -> [content:%v] [type:%v]\n", string(s.Content()), s.MetadataValue(ContentType))

	s.RemoveContent()
	s.SetContent(address{
		Street: "123 Oak Street",
		City:   "Anytown",
		State:  "USA",
		Zip:    "01234",
	})
	fmt.Printf("test: SetContent(address) -> [content:%v] [type:%v]\n", string(s.Content()), s.MetadataValue(ContentType))

	//Output:
	//test: SetContent(test string content) -> [content:test string content] [type:text/plain]
	//test: SetContent(test string content) -> [content:test string content] [type:application/json]
	//test: SetContent(12345) -> [content:12345] [type:application/json]
	//test: SetContent(address) -> [content:{"Street":"123 Oak Street","City":"Anytown","State":"USA","Zip":"01234"}] [type:application/json]

}

func ExampleStatus_TemplateParameter() {
	fmt.Printf("test: testTemplate() -> %v\n", len(testTemplate[address]()))

	//Output:
	//test: testTemplate() -> 0

}

func testTemplate[T address]() string {
	var t T
	switch a := any(t).(type) {
	case address:
		return a.GetZip()
	}
	return ""
}

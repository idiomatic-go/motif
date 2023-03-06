package messaging

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/motif/runtime"
	"time"
)

var msgTest = Message{To: "to-uri", From: "from-uri", Content: []any{
	"text content",
	500,
	Credentials(func() (username, password string, err error) { return "", "", nil }),
	time.Second,
	nil,
	runtime.Handle[runtime.DebugError](),
	errors.New("this is a content error message"),
	func() bool { return false },
	runtime.NewStatusError("location", errors.New("error message")).SetDuration(time.Second * 2),
	ActuatorApply(func(ctx context.Context, statusCode func() int, uri, requestId, method string) (func(), context.Context, bool) {
		return func() {}, ctx, false
	}),
	runtime.HandleWithContext[runtime.DebugError](),
	DatabaseUrl{"postgres://username:password@database.cloud.timescale.com/database?sslmode=require"},
}}

func ExampleAccessCredentials() {
	fmt.Printf("test: AccessCredentials(nil) -> %v\n", AccessCredentials(nil) != nil)
	fmt.Printf("test: AccessCredentials(msg) -> %v\n", AccessCredentials(&Message{To: "to-uri"}) != nil)
	fmt.Printf("test: AccessCredentials(msg) -> %v\n", AccessCredentials(&msgTest) != nil)

	//Output:
	//test: AccessCredentials(nil) -> false
	//test: AccessCredentials(msg) -> false
	//test: AccessCredentials(msg) -> true
}

func ExampleAccessDatabaseUrl() {
	fmt.Printf("test: AccessDatabaseUrl(nil) -> %v\n", AccessDatabaseUrl(nil))
	fmt.Printf("test: AccessDatabaseUrl(msg) -> %v\n", AccessDatabaseUrl(&Message{To: "to-uri"}))
	fmt.Printf("test: AccessDatabaseUrl(msg) -> %v\n", AccessDatabaseUrl(&msgTest))

	//Output:
	//test: AccessDatabaseUrl(nil) -> {}
	//test: AccessDatabaseUrl(msg) -> {}
	//test: AccessDatabaseUrl(msg) -> {postgres://username:password@database.cloud.timescale.com/database?sslmode=require}

}

func ExampleAccessActuatorApply() {
	fmt.Printf("test: AccessActuatorApply(nil) -> [valid:%v]\n", AccessActuatorApply(nil) != nil)
	fmt.Printf("test: AccessActuatorApply(msg) -> [valid:%v]\n", AccessActuatorApply(&Message{To: "to-uri"}) != nil)
	fmt.Printf("test: AccessActuatorApply(msg) -> [valid:%v]\n", AccessActuatorApply(&msgTest) != nil)

	//Output:
	//test: AccessActuatorApply(nil) -> [valid:false]
	//test: AccessActuatorApply(msg) -> [valid:false]
	//test: AccessActuatorApply(msg) -> [valid:true]

}

func ExampleNewStatusCodeFny() {
	var status *runtime.Status

	fn := NewStatusCode(&status)
	status = runtime.NewStatusCode(runtime.StatusDeadlineExceeded)
	fmt.Printf("test: NewStatusCode(&status) -> [statusCode:%v]\n", fn())

	//Output:
	//test: NewStatusCode(&status) -> [statusCode:4]

}

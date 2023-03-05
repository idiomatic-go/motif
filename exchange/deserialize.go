package exchange

import (
	"encoding/json"
	"errors"
	"github.com/idiomatic-go/motif/runtime"
	"io"
)

var deserializeLoc = pkgPath + "/deserialize"

// Deserialize - templated function, providing deserialization of a request/response body
func Deserialize[E runtime.ErrorHandler, T any](body io.ReadCloser) (T, *runtime.Status) {
	var e E
	var t T

	if body == nil {
		return t, e.Handle(deserializeLoc, errors.New("body is nil")).SetCode(runtime.StatusInvalidContent)
	}
	switch ptr := any(&t).(type) {
	case *[]byte:
		buf, err := ReadAll(body)
		if err != nil {
			return t, e.Handle(deserializeLoc, err).SetCode(runtime.StatusIOError)
		}
		*ptr = buf
	default:
		err := json.NewDecoder(body).Decode(&t)
		if err != nil {
			return t, e.Handle(deserializeLoc, err).SetCode(runtime.StatusJsonDecodeError)
		}
	}
	return t, runtime.NewStatusOK()
}

// ReadAll - read all the body, with a deferred close
func ReadAll(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	return io.ReadAll(body)
}

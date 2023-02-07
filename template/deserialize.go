package template

import (
	"encoding/json"
	"errors"
	"github.com/idiomatic-go/motif/runtime"
	"io"
)

var deserializeLoc = pkgPath + "/deserialize"

func Deserialize[T any](body io.ReadCloser) (T, *runtime.Status) {
	var t T

	if body == nil {
		return t, runtime.NewStatusError(deserializeLoc, errors.New("body is nil")).SetCode(runtime.StatusInvalidContent)
	}
	switch ptr := any(&t).(type) {
	case *[]byte:
		buf, err := ReadAll(body)
		if err != nil {
			return t, runtime.NewStatusError(deserializeLoc, err).SetCode(runtime.StatusIOError)
		}
		*ptr = buf
	default:
		err := json.NewDecoder(body).Decode(&t)
		if err != nil {
			return t, runtime.NewStatusError(deserializeLoc, err).SetCode(runtime.StatusJsonDecodeError)
		}
	}
	return t, runtime.NewStatusOK()
}

func ReadAll(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	return io.ReadAll(body)
}

package runtime

import (
	"errors"
	"fmt"
)

const (
	comment   = "//"
	delimiter = ":"
)

func ValidateConfig(m map[string]string, err error, keys ...string) (errs []error) {
	if m == nil {
		return []error{errors.New("config map is nil")}
	}
	if err != nil {
		errs = append(errs, errors.New(fmt.Sprintf("config map read error: %v", err)))
	}
	for _, k := range keys {
		if v, ok := m[k]; ok {
			if v == "" {
				errs = append(errs, errors.New(fmt.Sprintf("config map error: value for key does not exist [%v]", k)))
			}
		} else {
			errs = append(errs, errors.New(fmt.Sprintf("[config map error: key does not exist [%v]", k)))
		}
	}
	return
}

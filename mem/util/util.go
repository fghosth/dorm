package util

import (
	// "fmt"
	"errors"
)

type Dstring struct {
}

func (ds *Dstring) FUPer(str string) (string, error) {
	errEmpty := errors.New("字符串为空")
	v := []byte(str)
	if len(v) == 0 {
		return "", errEmpty
	}
	if v[0] < 97 {
		v[0] += 32
	}
	return string(v), nil
}

package interceptor

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	r := errors.New("hello")
	switch r.(type) {
	case error:
		fmt.Println("err")
	default:
		fmt.Println("string")
	}
}

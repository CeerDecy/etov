package handle

import (
	"fmt"
	"testing"
)

func TestTools(t *testing.T) {
	fmt.Println(parseParams(`{"a":"1","b":"2","c":"3"}`))
}

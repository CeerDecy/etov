package tools

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	md5 := MD5([]byte("123456"), []byte("U(*!,,@z"))
	fmt.Println(md5)
	md5 = MD5([]byte("123456"), nil)
	fmt.Println(md5)
}

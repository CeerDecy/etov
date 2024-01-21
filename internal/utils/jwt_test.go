package utils

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := GenerateTokenUsingHs256(10001)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)

	hs256, err := ParseTokenHs256(token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hs256.UserID)

}

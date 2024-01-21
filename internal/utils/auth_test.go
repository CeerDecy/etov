package utils

import "testing"

func TestSalt(t *testing.T) {
	t.Logf(GenerateSalt(8))
}

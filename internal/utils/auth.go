package utils

import "math/rand"

func GenerateSalt() string {
	salt := make([]byte, 8)
	for i := range salt {
		salt[i] = byte(rand.Intn(94) + 33)
	}
	return string(salt)
}

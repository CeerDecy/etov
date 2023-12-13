package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(v []byte) string {
	m := md5.New()
	m.Write(v)
	return hex.EncodeToString(m.Sum(nil))
}

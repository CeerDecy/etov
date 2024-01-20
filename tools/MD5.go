package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(v []byte, salt []byte) string {
	m := md5.New()
	m.Write(v)
	return hex.EncodeToString(m.Sum(salt))
}

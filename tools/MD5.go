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

func MD5Str(v string, salt string) string {
	m := md5.New()
	m.Write([]byte(v))
	return hex.EncodeToString(m.Sum([]byte(salt)))
}

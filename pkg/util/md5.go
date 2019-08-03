package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) (md string) {
	m := md5.New()
	m.Write([]byte(value))

	md = hex.EncodeToString(m.Sum(nil))
	return md
}

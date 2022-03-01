package pkg

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte("chengxisheng"))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

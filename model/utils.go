package model

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePasswordHash(pwd string) string {
	return Md5(pwd)
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

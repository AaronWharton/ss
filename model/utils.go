package model

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePasswordHash(pwd string) string {
	hash := md5.New()
	hash.Write([]byte(pwd))
	pwdHash:=hex.EncodeToString(hash.Sum(nil))
	return pwdHash
}

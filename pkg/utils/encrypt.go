package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

//使用sha256加盐加密密码
//Encrypting passwords with sha256 and salt.
func Encrypt(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt)) //密码与盐自定义组合
	res := hex.EncodeToString(hash.Sum(nil))
	return res
}

package tools

import (
	"crypto/sha512"
	"crypto/md5"
	"fmt"
)

//用一个花里胡哨的加密方式来加密密码，密码最多32位
func CryptoPwd(password string) string {
	shaed := sha512.Sum512_256([]byte(password))
	pwd := md5.Sum(shaed[:])
	return fmt.Sprintf("%x", pwd)
}

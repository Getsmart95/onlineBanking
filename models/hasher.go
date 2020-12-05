package dbupdate

import (
	"crypto/md5"
	"fmt"
)

func MakeHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

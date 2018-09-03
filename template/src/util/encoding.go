package util

import (
	"crypto/md5"
	"fmt"
)

// MD5 will return md5 sum string
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

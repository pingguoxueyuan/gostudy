package util

import (
	"crypto/md5"
	"fmt"
)

func Md5(data []byte) (result string) {
	md5Sum := md5.Sum(data)
	result = fmt.Sprintf("%x", md5Sum)
	return
}

package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Md5Encrypt MD5加密
func Md5Encrypt(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

package encrypt

import (
	"crypto/sha256"
	"fmt"
	"io"
)

//SHA256Encrypt sha256加密
func SHA256Encrypt(str string) string {
	h := sha256.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

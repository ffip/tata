package md5

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Encrypt(value string) string {
	h := md5.New()
	_, _ = h.Write([]byte(value))
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

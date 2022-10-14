package env

import (
	"github.com/ffip/tata/lib/algo/crypt/aes"
)

// Parse ==> parse cfg string.
func Parse(in string) (out []byte) {
	//获取明文内容并赋值到cfg.Data
	out = aes.DeCrypt(in, "obIf5qZWI0hW8Mdh2zvjbiC1g4cFbBxu")
	return
}

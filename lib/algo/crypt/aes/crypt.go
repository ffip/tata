package aes

import (
	"encoding/hex"
	"strings"

	"github.com/ffip/tata/lib/algo/crypt/aes/aesep5"
	"github.com/ffip/tata/lib/text"
)

// DeCrypt ==> DeCrypt string.
func DeCrypt(cipher string, pass string) []byte {
	//将HEX解码成[]byte
	hash, err := hex.DecodeString(cipher)
	if err != nil {
		//如果解码失败返回失败状态值
		return nil
	}
	plain := aesep5.AesDecrypt(hash, text.Atob(pass))
	if plain == nil {
		//如果解密失败返回失败状态值
		return nil
	}
	//返回解密解密结果及状态
	return plain
}

// EnCrypt ==> EnCrypt string.
func EnCrypt(plain []byte, pass string) string {
	//将加密key转换为[]byte
	key := text.Atob(pass)
	//加密明文
	hash := aesep5.AesEncrypt(plain, key)
	//将密文结果编码成HEX
	cipher := hex.EncodeToString(hash)
	return strings.ToUpper(cipher)
}

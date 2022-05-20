package aes

import (
	"fmt"
	"testing"

	"bitbucket.org/pwq/tata/lib/text"
)

func TestCryptToken(t *testing.T) {
	account := `{"token":"c1c0a3a37b859f0cf6e0918d42cf4d8c","ts":1572684441}`
	origin := text.Atob(account)
	hash := EnCrypt(origin, `4bdb2edf505adff111c7056e73a1adb3`)

	fmt.Println("hash:", hash)
	plain := DeCrypt(hash, `4bdb2edf505adff111c7056e73a1adb3`)
	fmt.Println("plain:", text.Abot(plain))
}

package client

import (
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	s := Client{}
	s.Request.Method = "GET"
	s.Request.Url = "https://www.baidu.com"
	s.Do()
	fmt.Println(s.Result.Body)
}

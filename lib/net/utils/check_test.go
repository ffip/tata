package utils

import (
	"fmt"
	"testing"
)

func TestTcpPortCheck(t *testing.T) {
	port := 80
	addr, _ := GetRandomTCPAddress(port, 8080)
	fmt.Println(addr.Port)
}

package utils

import (
	"fmt"
	"math/rand"
	"net"
)

// GetRandomTCPAddress ==>  Generates a TCP Address with localhost:random-port
func GetRandomTCPAddress(min, max int) (addr *net.TCPAddr, err error) {
	return net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", min+rand.Intn(max-min)))

}

// GetRandomUDPAddress ==>  Generates a UDP Address with localhost:random-port
func GetRandomUDPAddress(min, max int) (addr *net.UDPAddr, err error) {
	return net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", min+rand.Intn(max-min)))
}

// GetInterfaceAddresses ==>  Get Device Network Interfaces
func GetInterfaceAddresses() (addrs []net.IP, err error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, ifc := range ifs {
		iaddrs, err := ifc.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range iaddrs {
			ip, _, _ := net.ParseCIDR(addr.String())
			if ip.To4() != nil {
				addrs = append(addrs, ip.To4())
			}
		}
	}

	return
}

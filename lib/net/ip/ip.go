package ip

import (
	"net"
	"strings"
)

// getLocalIpV4 获取 IPV4 IP，没有则返回空
func GetLocalIpV4() string {
	inters, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, inter := range inters {
		// 判断网卡是否开启，过滤本地环回接口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			// 获取网卡下所有的地址
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//判断是否存在IPV4 IP 如果没有过滤
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

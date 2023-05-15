package utils

import (
	"fmt"
	"net"
)

func GetIp() (ipList []string, err error) {
	addrs, getAddressErr := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("无法获取网络接口：", err)
		ipList = nil
		err = getAddressErr
		return
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { // 当前计算机的IPv4地址：
				ipList = append(ipList, ipnet.IP.String())
			}
		}
	}
	return
}
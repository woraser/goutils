/*
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @package     helper
 * @date        2018-01-25 19:19
 */
package helper

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
)

// get local ip address
func LocalIP() (ips []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return
}

// check ip is private?
func IsPrivateIP(ip string) bool {
	longip := Ip2long(ip)
	//10.0.0.0-10.255.255.255
	if longip&0xFF000000 == 0x0A000000 {
		return true
	}
	//172.16.0.0-172.31.255.255
	if longip&0xFFF00000 == 0xAC100000 {
		return true
	}
	//192.168.0.0-192.168.255.255
	if longip&0xFFFF0000 == 0xC0A80000 {
		return true
	}
	return false
}

// convert ip from string to uint2
func Ip2long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}
	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])
	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}
	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)
	return
}

// convert ip from uint32 to string
func Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}

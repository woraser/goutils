package helper

import (
	"testing"
	"runtime"
	"fmt"
	"github.com/kr/pretty"
)

func TestServerIP(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	serverIP := LocalIP()
	fmt.Printf("serverIP %# v\n", pretty.Formatter(serverIP))
	for _, ip := range serverIP {
		b := "true"
		if !IsPrivateIP(ip) {
			b = "false"
		}
		fmt.Printf("localIP:%s\tisPrivate:%s\tlong=%d\n", ip, b, Ip2long(ip))
	}
}

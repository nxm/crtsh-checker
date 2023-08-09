package misc

import (
	"net"
	"time"
)

func CheckDomainAlive(domain string) bool {
	timeout := time.Second * 5
	_, err80 := net.DialTimeout("tcp", domain+":80", timeout)
	_, err443 := net.DialTimeout("tcp", domain+":443", timeout)

	if err80 != nil && err443 != nil {
		return false
	}

	return true
}

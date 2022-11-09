package scanner

import (
	"fmt"
	"net"
	"time"
)

func tcpConnect(ip string, port int) error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.SetDeadline(time.Now())
	if err != nil {
		return err
	}
	return nil
}

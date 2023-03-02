package plugins

import (
	"fmt"
	"geektime-web/flag/val"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

func SshScan(ip string, port string) {
	for _, user := range val.SshUsers {
		for _, password := range val.Passwords {
			err := SshConnect(user, password, ip, port)
			if err == nil {
				Results["ssh"] = fmt.Sprintf("ssh %s:%s  user:%s password:%s\n", ip, port, user, password)
				return
			} else {
				fmt.Printf("ssh %s:%s  user:%s password:%s %v \n", ip, port, user, password, err)
			}
		}
	}
}

func SshConnect(user string, password string, ip string, port string) error {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout: time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", ip, port), config)
	//defer conn.Close()
	if err != nil {
		return err
	}
	session, err := conn.NewSession()
	//defer session.Close()
	if err != nil {
		return err
	}
	b, _ := session.CombinedOutput("ls")
	fmt.Println(string(b))
	return nil
}

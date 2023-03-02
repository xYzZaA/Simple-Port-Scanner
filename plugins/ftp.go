package plugins

import (
	"fmt"
	"geektime-web/flag/val"
	"github.com/jlaffaye/ftp"
	"time"
)

func FtpScan(ip string, port string) {
	err := FtpConnect(ip, port, "anonymous", "")
	if err == nil {
		Results["ftp"] = fmt.Sprintf("ftp %s:%s  user:anonymous  password:\n", ip, port)
		return
	}
	for _, ftpUser := range val.FtpUsers {
		for _, password := range val.Passwords {
			err := FtpConnect(ip, port, ftpUser, password)
			if err == nil {
				Results["ftp"] = fmt.Sprintf("ftp %s:%s  user:%s password:%s\n", ip, port, ftpUser, password)
				return
			} else {
				fmt.Printf("ftp %s:%s  user:%s password:%s %v \n", ip, port, ftpUser, password, err)
			}
		}
	}
}

func FtpConnect(ip string, port string, user string, password string) error {

	conn, err := ftp.Dial(fmt.Sprintf("%s:%s", ip, port), ftp.DialWithTimeout(time.Second))
	defer conn.Quit()
	if err != nil {
		return err
	}
	if err := conn.Login(user, password); err != nil {
		return err
	}
	return nil
}

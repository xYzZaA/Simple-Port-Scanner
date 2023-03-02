package plugins

import (
	"database/sql"
	"fmt"
	"geektime-web/flag/val"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func MysqlScan(ip string, port string) {
	for _, user := range val.MysqlUsers {
		for _, password := range val.Passwords {
			err := MysqlConnect(user, password, ip, port)
			if err == nil {
				Results["mysql"] = fmt.Sprintf("mysql %s:%s  user:%s password:%s\n", ip, port, user, password)
				return
			} else {
				fmt.Printf("mysql %s:%s  user:%s password:%s %v \n", ip, port, user, password, err)
			}
		}
	}
}

// 用户名:密码@(地址:端口)/数据库名称
func MysqlConnect(user string, password string, ip string, port string) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/?timeout=%v", user, password, ip, port, time.Second))
	defer db.Close()
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}

package plugins

import (
	"fmt"
	"geektime-web/flag/val"
	"github.com/gomodule/redigo/redis"
	"strings"
)

func RedisScan(ip string, port string) {
	if val.RedisUser != "" {
		users := strings.Split(val.RedisUser, ",")
		for _, user := range users {
			for _, password := range val.Passwords {
				err := RedisConnect(user, password, ip, port)
				if err == nil {
					Results["redis"] = fmt.Sprintf("redis %s:%s user:%s password:%s\n", ip, port, user, password)
					return
				} else {
					fmt.Printf("redis %s:%s user:%s password:%s %v \n", ip, port, user, password, err)
				}
			}
		}
	} else {
		for _, password := range val.Passwords {
			err := RedisConnect("", password, ip, port)
			if err == nil {
				Results["redis"] = fmt.Sprintf("redis %s:%s  password:%s\n", ip, port, password)
				return
			} else {
				fmt.Printf("redis %s:%s  password:%s %v \n", ip, port, password, err)
			}
		}
	}
}

func RedisConnect(user string, password string, ip string, port string) error {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return err
	}
	defer conn.Close()
	if user != "" {
		res, err := conn.Do("auth", user, password)
		if err != nil {
			return err
		} else
		{
			fmt.Println(res)
			return nil
		}
	}
	res, err := conn.Do("auth", password)
	if err != nil {
		return err
	} else {
		fmt.Println(res)
		return nil
	}
}

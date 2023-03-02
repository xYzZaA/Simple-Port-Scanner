package plugins

import (
	"fmt"
	"geektime-web/flag/val"
	"strings"
)

type ScanFunc func(ip string, port string)

var Plugins = map[int]ScanFunc{
	21:   FtpScan,
	3306: MysqlScan,
	517:  SshScan,
	6379: RedisScan,
}

var Results = map[string]string{
	"ftp":   "ftp Username and password not found\n",
	"mysql": "mysql Username and password not found\n",
	"ssh":   "ssh Username and password not found\n",
	"redis": "redis password not found\n",
}

func PrintResults() {
	if val.Blast == "all" {
		for _, result := range Results {
			fmt.Printf("%v", result)
		}
		return
	}
	list := strings.Split(val.Blast, ",")
	for _, li := range list {
		fmt.Println(Results[li])
	}
}

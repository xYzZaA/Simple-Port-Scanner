package sflag

import (
	"flag"
	"geektime-web/flag/val"
)

func FlagInit() {
	flag.StringVar(&val.Address, "a", "127.0.0.1", "Enter the address you want to search")
	flag.StringVar(&val.Port, "p", "", "port")
	flag.IntVar(&val.ThreadNum, "n", 300, "Enter an integer to control the number of goroutine")
	flag.StringVar(&val.Mode, "m", "connect", "Select scan mode, syn or connect, syn for linux systems only")
	//flag.Usage = func() {
	//	fmt.Fprintf()
	//	flag.PrintDefaults()
	//}
	//如果想要自定义 -help 输出的提示信息，就在这里修改
	flag.Parse()
}

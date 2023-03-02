package sflag

import (
	"flag"
	"geektime-web/flag/val"
)

func FlagInit() {
	flag.StringVar(&val.Address, "a", "127.0.0.1", "Enter the address you want to search")
	flag.StringVar(&val.Port, "p", "", "port")
	flag.IntVar(&val.ThreadNum, "n", 450, "Enter an integer to control the number of goroutine")
	flag.StringVar(&val.Mode, "m", "connect", "Select scan mode, syn or connect, syn for linux systems only")
	flag.BoolVar(&val.OnlyIcmp, "oi", false, "No port scanning, only host survival detection")
	flag.BoolVar(&val.NoBlast, "nobt", false, "Whether to perform blasting")
	flag.StringVar(&val.Blast, "bt", "all", "Specify the blast range, the default is all")
	flag.BoolVar(&val.Output, "o", false, "Output the results to a txt file")
	flag.StringVar(&val.Dict, "dict", "", "Specify the password dictionary")
	flag.StringVar(&val.RedisUser, "ru", "", "Higher versions of redis need to enter the username of redis, which can receive more than one, separated by ,")
	//flag.Usage = func() {
	//	fmt.Fprintf()
	//	flag.PrintDefaults()
	//}
	//如果想要自定义 -help 输出的提示信息，就在这里修改
	flag.Parse()
}

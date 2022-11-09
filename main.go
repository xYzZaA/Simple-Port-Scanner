package main

import (
	"fmt"
	"geektime-web/flag/scanner"
	"geektime-web/flag/sflag"
	"time"
)

func main() {
	start := time.Now()
	sflag.FlagInit()
	p := scanner.NewPort()
	p.FoundPort()
	end := time.Since(start)
	fmt.Printf("\n程序执行时长: 	%s \n", end)
}

//  保证main函数中始终是
//  p := scanner.NewPort()
//	p.FoundPort()
//  两步完成工作

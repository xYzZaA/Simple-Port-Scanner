package scanner

import (
	"fmt"
	"geektime-web/flag/val"
	"sync"
)

type Scanner struct {
	wg     sync.WaitGroup   //保证所有协程执行完成
	ch     chan struct{}    //控制协程个数
	m      sync.Mutex       //写保护
	result map[string][]int //存放结果
}

func NewPort() *Scanner {
	return &Scanner{
		wg:     sync.WaitGroup{},
		ch:     make(chan struct{}, val.ThreadNum),
		m:      sync.Mutex{},
		result: make(map[string][]int, 0),
	}
} //暴露的创建结构体的方法
// FoundPort 一次扫1000个端口会出问题，出现很多没开的接口，改成500*2次扫描，如果要保全速度，可以相应的减少建立连接的超时时间
// 还有问题，再对分组测出来的端口做二测
func (p *Scanner) FoundPort() { //这里还能改改，不需要让FoundPort知道task的内部结构
	tasks, _ := val.GetTasks()
	for _, task := range tasks {
		for ip, ports := range task {
			for _, port := range ports {
				p.ch <- struct{}{}
				go p.ScanPort(ip, port)
				p.wg.Add(1)
			}
			p.wg.Wait()
		}
	}
	p.PrintPorts()
}
func (p *Scanner) ScanPort(ip string, port int) {
	var whichSide error
	if val.Mode == "connect" || (val.Mode == "syn" && val.System == "windows") {
		whichSide = tcpConnect(ip, port)
	} else if val.Mode == "syn" && val.System == "linux" {
		whichSide = SynScan(ip, port)
	}
	if whichSide != nil {
	} else {
		p.m.Lock()
		p.SaveResults(ip, port)
		p.m.Unlock()
	}
	p.wg.Done()
	<-p.ch
}
func (p *Scanner) SaveResults(ip string, port int) {
	if v, ok := p.result[ip]; ok {
		temp := append(v, port)
		p.result[ip] = temp
	} else {
		ports := make([]int, 0)
		ports = append(ports, port)
		p.result[ip] = ports
	}
} //保存结果

func (p *Scanner) PrintPorts() {
	fmt.Println("IP             port          service")
	for ip, ports := range p.result {
		for _, port := range ports {
			fmt.Printf("%-15s", ip)
			fmt.Printf("%-14d", port)
			fmt.Printf("%-s\n", val.ReturnService(port))
		}
	}
} //最后的打印输出

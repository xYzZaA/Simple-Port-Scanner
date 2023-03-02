package scanner

import (
	"bufio"
	"fmt"
	"geektime-web/flag/plugins"
	"geektime-web/flag/val"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Scanner struct {
	wg     sync.WaitGroup   //保证所有协程执行完成
	ch     chan string      //传数据的
	m      sync.Mutex       //写保护
	result map[string][]int //存放结果
	writer *bufio.Writer    //写入文件,result.txt
}

func NewPort() *Scanner {
	return &Scanner{
		wg:     sync.WaitGroup{},
		ch:     make(chan string, val.ThreadNum), //试试ch放数据(ip+port)，协程读数据
		m:      sync.Mutex{},
		result: make(map[string][]int, 0),
		writer: NewWriter(),
	}
} //暴露的创建结构体的方法

func NewWriter() *bufio.Writer {
	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("fail to create file")
		os.Exit(0)
	}
	writer := bufio.NewWriter(f)
	return writer
}

// FoundPort 一次扫1000个端口会出问题，出现很多没开的接口
// 还有问题，再对分组测出来的端口做二测
func (p *Scanner) FoundPort() { //这里还能改改，不需要让FoundPort知道task的内部结构
	parse()

	ips, _ := val.GetIpList(val.Address)
	ports, _ := val.GetPortList(val.Port)
	if ips == nil {
		fmt.Println("invalid port")
		os.Exit(0)
	}
	if len(ports) == 0 {
		ports = val.PortCommon
	}
	Hosts, err := checkLive(ips)
	if err != nil {
		fmt.Println("must run with admin")
		os.Exit(0)
	}

	fmt.Printf("alive host len is %d \n", len(Hosts))

	if val.OnlyIcmp {
		for _, host := range Hosts {
			fmt.Printf("%s   live\n", host)
			p.writer.WriteString(fmt.Sprintf("%s   live\n", host))
		}
		p.writer.Flush()
		os.Exit(0)
	}

	for i := 0; i < val.ThreadNum; i++ {
		go p.ScanPort()
	}
	for _, host := range Hosts {
		for _, port := range ports {
			p.wg.Add(1)
			p.ch <- fmt.Sprintf("%s:%d", host, port)
		}
	}
	p.wg.Wait()
	p.PrintPorts()

	if !val.NoBlast {
		for ip, ports := range p.result {
			for _, port := range ports {
				switch port {
				case 21:
					if val.Blast == "all" || strings.Contains(val.Blast, "ftp") {
						go p.AddScanFunc(ip, port)
						p.wg.Add(1)
					}
				case 3306:
					if val.Blast == "all" || strings.Contains(val.Blast, "mysql") {
						go p.AddScanFunc(ip, port)
						p.wg.Add(1)
					}
				case 517:
					if val.Blast == "all" || strings.Contains(val.Blast, "ssh") {
						go p.AddScanFunc(ip, port)
						p.wg.Add(1)
					}
				case 6379:
					if val.Blast == "all" || strings.Contains(val.Blast, "redis") {
						go p.AddScanFunc(ip, port)
						p.wg.Add(1)
					}
				}
			}
		}
		p.wg.Wait()
		plugins.PrintResults()
	}
}
func (p *Scanner) ScanPort() {
	for ipPort := range p.ch {
		ip := strings.Split(ipPort, ":")[0]
		temp := strings.Split(ipPort, ":")[1]
		port, _ := strconv.Atoi(temp)
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
	}
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
			if val.Output {
				p.writer.WriteString(fmt.Sprintf("%-15s:%-14d %-s\n", ip, port, val.ReturnService(port)))
				p.writer.Flush()
			}
		}
	}
} //最后的打印输出

func (p *Scanner) AddScanFunc(ip string, port int) {
	f := plugins.Plugins[port]
	f(ip, strconv.Itoa(port))
	p.wg.Done()
}

func parse() {
	if val.Dict != "" {
		var password []string
		f, err := os.Open(val.Dict)
		if err != nil {
			fmt.Println("fail to open dictionary")
			os.Exit(0)
		}
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			password = append(password, scanner.Text())
		}
		val.Passwords = password
	}
}

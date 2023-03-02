package scanner

import (
	"net"
	"time"
)

func checkLive(ips []string) ([]string, error) {

	hostLive := make([]string, 0)
	conn1, err := net.ListenPacket("ip:icmp", "0.0.0.0")
	if err != nil {
		return hostLive, err
	}
	defer conn1.Close()

	icmp := make([]byte, 40)
	icmp[0] = 8
	icmp[1] = 0
	icmp[2] = 0
	icmp[3] = 0
	icmp[4] = 0
	icmp[5] = 0
	icmp[6] = 0
	icmp[7] = 0
	check := checkSum(icmp[0:40])
	icmp[2] = byte(check >> 8)   //高八位
	icmp[3] = byte(check & 0xff) //低八位

	for _, ip := range ips {
		go func(ip string) {
			dstIp, _ := net.ResolveIPAddr("ip", ip)
			conn1.WriteTo(icmp, dstIp)
		}(ip)
	}

	go func() {
		for {
			answer := make([]byte, 100)
			_, addr, _ := conn1.ReadFrom(answer)
			if addr != nil {
				hostLive = append(hostLive, addr.String())
			}
		}
	}()

	timeBegin := time.Now()
	for {
		if len(hostLive) == len(ips) {
			break
		}
		if time.Now().Sub(timeBegin) > time.Second*3 {
			break
		}
	}
	return hostLive, nil
}

func checkSum(msg []byte) uint16 {
	sum := 0
	length := len(msg)
	for i := 0; i < length-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if length%2 == 1 {
		sum += int(msg[length-1]) * 256 //剩下的一个字节,作为高8位,低八位补00,所以要乘以256
	}
	sum = (sum >> 16) + (sum & 0xffff) //高十六位与低十六位相加
	sum = sum + (sum >> 16)            //如果高十六位还有数(溢出)接着加
	answer := uint16(^sum)
	return answer
}

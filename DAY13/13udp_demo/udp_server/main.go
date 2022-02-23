package main

import (
	"fmt"
	"net"
	"strings"
)

/*
UDP协议（User Datagram Protocol）中文名称是用户数据报协议，
是OSI（Open System Interconnection，开放式系统互联）参考模型中一种无连接的传输层协议，
不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，
但是UDP协议的实时性比较好，通常用于视频直播相关领域。
*/
// UDP服务端
// UDP/server/main.go

// UDP server端
func main() {

	//1、本地启动端口
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0), //ip地址，指定地址为0.0.0.0的地址，这个地址事实上表示不确定地址，或“所有地址”、“任意地址”。
		Port: 30000,                //端口
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer conn.Close()

	//不需要建立连接，直接收发数据
	for {
		var data [1024]byte
		n, addr, err := conn.ReadFromUDP(data[:]) // 读取客户端发送的数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			return
		}
		fmt.Println(data[:n]) //打印接收到的客户端数据
		reply := strings.ToUpper(string(data[:n]))

		//发送数据给客户端
		conn.WriteToUDP([]byte(reply), addr)
	}
}

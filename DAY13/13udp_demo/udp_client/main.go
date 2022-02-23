package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//UDP客户端

func main() {

	//拨号，连接服务端
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容：")
		msg, _ := reader.ReadString('\n')

		//发送数据给服务端
		socket.Write([]byte(msg))

		//读取服务端回复的数据
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("redv reply msg failed,err：", err)
			return
		}
		fmt.Println("服务端回复：", string(reply[:n]))
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//TCP服务端
/* 一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。
因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理。 */

/* TCP服务端程序的处理流程：
1、监听端口
2、接收客户端请求建立链接
3、创建goroutine处理链接。 */

//操作：
//1、本地端口启动服务
//2、等待别人来跟我建立连接
//3、与客户端通信

// 我们使用Go语言的net包实现的TCP服务端代码如下：
/*
func main() {

	//1、本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000") //"127.0.0.1:20000"是本地回环ip+本服务端的端口号
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	//2、等待别人来跟我建立连接
	conn, err := listen.Accept() // 建立连接
	if err != nil {
		fmt.Println("accept failed, err:", err)
		return
	}

	//3、与客户端通信
	var a [128]byte
	n, err := conn.Read(a[:])
	if err != nil {
		fmt.Println("read from conn failed,err:", err)
		return
	}
	fmt.Println(string(a[:n]))
}
*/
// tcp/server/main.go

// TCP server端（服务端）

// 处理函数

func process(conn net.Conn) {
	var buf [128]byte
	reader := bufio.NewReader(os.Stdin)
	//conn.Write()发送消息给客户端
	conn.Write([]byte("你好，我是服务端，请问需要什么帮助"))
	defer conn.Close() // 关闭连接
	for {

		//服务端读取客户端消息
		n, err := conn.Read(buf[:]) // 读取客户端数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		fmt.Println("客户端：", string(buf[:n]))

		//服务端输入消息
		fmt.Print("服务端输入消息：")
		msg, _ := reader.ReadString('\n') // 读取服务端输入
		msg = strings.TrimSpace(msg)
		if msg == "exit" { // 如果输入exit就退出
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {

	//服务端监听：127.0.0.1 [ 本机虚拟网卡的ip ]
	//本机只能通过127.0.0.1访问位于本机的服务端，不能通过本机IP访问

	//服务端监听：本机的IP [ 指的是本机物理网卡(无线网卡和有限网卡)所绑定的网络协议地址ip ]
	//本机不能通过127.0.0.1访问位于本机的服务端，本网络主机（同网段）的客户端都能通过服务端的IP访问位于本机的服务端

	//服务端监听：0.0.0.0 [ 即代表本机的所有ip地址，127.0.0.1和本机ip都包含了 ]
	//本机能通过127.0.0.1和IP访问位于本机的服务端，本网络主机的客户端也可以通过服务端设备的IP访问

	//1、本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	//2、等待别人来跟我建立连接
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}

		//3、与客户端通信
		go process(conn) // 启动一个goroutine处理连接
	}
}

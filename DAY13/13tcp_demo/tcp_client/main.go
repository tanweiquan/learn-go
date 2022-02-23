package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp客户端

//tcp client (客户端)
/* 一个TCP客户端进行TCP通信的流程如下：
1、建立与服务端的链接
2、进行数据收发
3、关闭链接 */

// 操作：
//1、与server端建立连接
//2、发送数据
//3、关闭连接
/*
func main() {

	//1、与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000") //"127.0.0.1:20000"是要连的服务端的ip+端口号
	if err != nil {
		fmt.Println("dial failed，err:", err)
		return
	}

	//2、发送数据
	conn.Write([]byte("hello 网页"))

	//3、关闭连接
	conn.Close()
} */

// tcp/client/main.go
// 客户端
func main() {

	//1、与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}

	//2、发送数据
	var mserver [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {

		//客户端读取服务端消息
		n, err := conn.Read(mserver[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		fmt.Println("服务端：", string(mserver[:n]))

		//客户端输入消息
		fmt.Print("客户端请输入消息：")
		msg, _ := reader.ReadString('\n') // 读取用户输入
		msg = strings.TrimSpace(msg)
		if msg == "exit" { // 如果输入exit就退出
			break
		}
		//conn.Write()发送消息给服务端
		conn.Write([]byte(msg))
	}

	//3、关闭连接
	conn.Close() // 关闭连接
}

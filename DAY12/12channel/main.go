package main

import (
	"fmt"
)

//channel
// 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。
// channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
// Go 语言中的通道（channel）是一种特殊的类型。
// 通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
// 每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

//channel是一种类型，一种引用类型。声明通道类型的格式如下：
//var 变量 chan 元素类型
//chan（通道）必须使用make函数初始化才能使用,格式为：
//x:=make(chan 元素类型, [缓冲大小])
//chan（通道）操作：
//1、发送：ch <- 10 (把10发送到ch中)
//2、接收：x:= <- ch (从ch中接收值并赋值给变量x) 或 <- ch (从ch中接收值，忽略结果)。
//3、关闭：close(ch)
/* 关闭后的通道有以下特点：
对一个关闭的通道再发送值就会导致panic。
对一个关闭的通道进行接收会一直获取值直到通道为空。
对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
关闭一个已经关闭的通道会导致panic。
*/

/* 当一个标识符被声明为chan后，它就是通道，
当开启多个goroutine使用这一个通道去取值时，通道会不断往缓冲区存值，而不像变量那样会重置后赋值 */

var a chan int   //声明一个传递整型的通道
var b chan []int //声明一个传递int切片的通道
var c chan bool  // 声明一个传递布尔型的通道
func main() {
	fmt.Println(a)          //<nil>
	b = make(chan []int)    //chan的初始化
	fmt.Println(b)          //0xc000086060
	c = make(chan bool, 16) //带缓冲区的chan初始化
	fmt.Println(c)          //0xc000018150
}

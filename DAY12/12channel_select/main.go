package main

import "fmt"

//select
//在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。
//1、使用select语句能提高代码的可读性。
//2、可处理一个或多个channel的发送/接收操作。
//3、如果多个case同时满足，select会随机选择一个。
//4、对于没有case的select{}会一直等待，可用于阻塞main函数。
/* select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。
select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下：
select{
    case <-ch1:
        ...
    case data := <-ch2:
        ...
    case ch3<-data:
        ...
    default:
        默认操作
}
*/

//select只执行一个case，如果所有case都满足时，会随机执行一个不阻塞通道的case

func main() {
	ch := make(chan int, 1) //声明ch是一个缓冲区为1，发送int数据的通道
	for i := 0; i < 10; i++ {
		select { //选择一个执行，当都满足时，随机选择一个执行，没有可执行则一直等待，造成阻塞
		case x := <-ch: //第一次循环时，ch中还没有值，因此，第一次循环不执行下面打印x值的操作
			fmt.Println(x) //打印结果是0,2,4,6,8，因为chan不存值，假如从chan读取数据后，chan已经空了
		case ch <- i: //第一次循环时，执行对ch发送i的值的操作，ch中就有了0，后面继续执行，对ch通道发送了0,2,4,6,8
		default:
			fmt.Println("通道阻塞或者通信完成")
		}
	}
}

/* 以上描述了 select 语句的语法：
每个 case 都必须是一个通信
所有 channel 表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行，其他被忽略。
如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
否则：
如果有 default 子句，则执行该语句。
如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。 */

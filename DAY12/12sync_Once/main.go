package main

import (
	"fmt"
	"sync"
)

//sync.Once
/*在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等
Go语言中的sync包中提供了一个针对只执行一次场景的解决方案：sync.Once */

//sync.Once的签名为：func (o *Once) Do(f unc()) {}
//使用sync.Once示例：
var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i //ch1通道中被依次传值0到99
	}
	close(ch1)
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1 //将ch1中的值依次赋值给x
		if !ok {
			break
		}
		ch2 <- x * x //通道ch2中被x*x依次传值
	}
	//通道只能关闭一次，当两个goroutine执行这个函数时，为防止重复关闭ch2，使用sync.Once
	once.Do(func() { close(ch2) }) //在once.Do(闭包)，确保操作只执行一次
}
func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a) //这步骤执行完后，通道a中被依次传值0到99
	//下面开启了两个goroutine执行任务，当通道里没有值时，其中的一个goroutine会关闭通道
	//将通道a中的值依次赋值给了x，然后通道b中被x*x依次传值
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}

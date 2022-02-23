package main

import "fmt"

//单向通道
//有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。
//Go语言中提供了单向通道来处理这种情况。例如，我们把上面的例子改造如下：
//声明单向通道时，发送和接收的区分：
//变量 chan<- 类型(只写单向通道，只能对其写入对应类型的值) 可以对其执行发送操作但是不能执行接收操作 如 var s chan<- int,可以对s发送数据
//变量 <-chan 类型(只读单向通道，只能从其读取对应类型的值) 可以对其执行接收操作但是不能执行发送操作 如 var m <-chan int,可以从m接收数据

func counter(out chan<- int) { //这里的函数声明了参数是单向通道，只能对通道out发送数据
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) { //这里的函数声明了参数是单向通道，分别对通道out发送数据和从通道in接收数据
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) { //这里的函数声明了参数是单向通道，可以从通道in接收数据
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

/* chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
<-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。
在函数传参及任何赋值操作中可以将双向通道转换为单向通道，但反过来是不可以的。 */

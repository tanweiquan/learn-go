package main

import (
	"fmt"
	"time"
)

//goroutine：goroutine是先登记后操作，即先登记函数为协程函数，后运行被登记为协程的函数
//Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
//一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。
//goroutine创建时，会马上执行go后面的函数传参操作(有参传参，无参阻塞，等有参再传)，创建完成后，再运行函数
//创建是需要时间的，因此再创建期间，代码继续按行执行下去，当创建真正完成后马上就运行
//当遇到for循环时，goroutine创建期间，for继续循环几轮，即创建了几个goroutine
//多个goroutine执行一个函数，因为每行的(每个开启)goroutine速度不一样，因此造成了随机性
//因为不同机器或不同版本的golnag,创建的每个goroutine的运行速度不同，会导致在有goroutine的for循环中输出结果的顺序是不一样的

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	//注意这里Windows系统执行代码是看不出来区别的，用Mac系统才看得出
	//多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
	for i := 0; i < 100; i++ { //开启100个单独的goroutine去执行hello函数(任务)
		go hello(i)
	}
	fmt.Println("沙河")
	time.Sleep(time.Second)
	//main函数结束了，由main函数启动的goroutine也都结束了
}

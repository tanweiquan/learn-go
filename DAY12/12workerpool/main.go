package main

import (
	"fmt"
	"time"
)

//worker pool（goroutine池）
//在工作中我们通常会使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。

/*
接到5个任务后，通道传值，goroutine就开始运行，因为是运行同一个函数，
因此当其中一个任务完成后，另外一个goroutine会继续接任务
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 开启3个goroutine
	// 这就是worker pool(goroutine池)
	for w := 1; w <= 3; w++ { //开启了3个goroutine后，因为通道没有传值，就一直阻塞，等通道传值后运行
		go worker(w, jobs, results)
	}
	//这时，goroutine运行了，3个goroutine先后运行完，先完成的抢先继续执行任务（同一个函数）
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results //当上面任务运行后，results里的缓冲区存了5个值，为了通道不阻塞而使goroutine能运行完，要将缓冲区的值都丢弃
	}
}

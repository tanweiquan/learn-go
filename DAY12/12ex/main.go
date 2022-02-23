package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/* 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
1、开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2、开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3、主goroutine从resultChan取出结果并打印到终端输出 */
// job...
type job struct {
	value int64
}

// result...
type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)       //声明了是传指针的通道
var resultChan = make(chan *result, 100) //声明了是传指针的通道
var wg sync.WaitGroup

func A(a chan<- *job) {
	defer wg.Done()
	// 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		a <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}
func B(a <-chan *job, resultChan chan<- *result) { //goroutine池要启动的函数，多个goroutine在完成任务后继续抢任务
	defer wg.Done()
	for {
		job := <-a
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult

	}
}
func main() {
	/* n := 123
	for n > 0 {
		fmt.Println(n % 10)// 3 2 1
		n = n / 10
	} */
	wg.Add(1)
	go A(jobChan)
	wg.Add(24)
	// 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	for i := 0; i < 24; i++ {
		go B(jobChan, resultChan)
	}
	// 主goroutine从resultChan取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
}

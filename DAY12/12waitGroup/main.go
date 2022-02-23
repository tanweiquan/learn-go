package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//协程同步(goroutine同步)
//waitGroup
//Add(delta int)添加协程记录
//Done()移除协程记录
//Wait()同步等待所有记录的协程全部结束

/*使用格式：
 var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
} */

/* func f() {
	for i := 0; i < 5; i++ { //打印出r1和r2各5个随机数
		rand.Seed(time.Now().UnixNano()) //随机数种子，保证每次执行的时候都有点不一样
		r1 := rand.Int()                 //返回一个int类型的随机数
		r2 := rand.Intn(10)              //返回0到10的随机数(0<=x<10)
		fmt.Println(r1)
		fmt.Println(r2)
	}
} */
var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done() //每执行完函数就减1
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}
func main() {
	//f()
	//wg.Add(10)//启动10个goroutine，一共加10
	for i := 0; i < 10; i++ {
		wg.Add(1) //每启动一个goroutine，计数器就加1
		go f1(i)
	}
	//如何知道这10个goroutime都结束了
	//time.Sleep(?)
	wg.Wait() //等待wg的计数器减为0
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//进程：软件一打开就创建了一个进程
//线程：一个进程里有多个线程处理任务

//goroutine的初始栈的大小为2k,可以轻松创建十万个goroutine,而OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）
//goroutine调度模型：GPM
//G:G是goroutine,以及所在P的绑定等信息
//P:P管理着一组goroutine队列
//M:M（machine）是Go运行时对操作系统内核线程的虚拟,M与内核线程一般是一一映射的关系
/* P管理着一组G挂载在M上运行。
当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G挂载在新建的M上。
当旧的G阻塞完成或者认为其已经死掉时,回收旧的M */

/* P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。
在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。 */

//GOMAXPROCS
//runtime.GOMAXPROCS，不设置时，默认是CPU的逻辑核心数，默认是跑满整个CPU线程数的
//Go语言中的操作系统线程和goroutine的关系：
//1、一个操作系统线程对应用户态多个goroutine
//2、go程序可以同时使用多个操作系统线程
//3、goroutine和OS线程是多对多的关系，即m:n (把m个goroutine分配给n个OS线程去执行)[OS是操作系统]

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Printf("a:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Printf("b:%d\n", i)
	}
}

func main() {

	runtime.GOMAXPROCS(2) //将逻辑核心数设为2，此时两个任务并行执行
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
	// fmt.Println(runtime.NumCPU())打印电脑cpu线程数(核心数)
	// runtime.GOMAXPROCS(1)
	//两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。即串行执行(两个任务随机先做一个，再做一个,即a()与b()谁先后执行取决于go版本和电脑系统)
	// wg.Add(2)
	// go a()
	// go b()
	// wg.Wait()

	//下面代码打印，先打印A和先打印B都有可能，即AAABBBBBAAAABBB或者BAAABBBAAABBAAA又或者其他，都有可能
	runtime.GOMAXPROCS(1)
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			// i打破for循环后赋值为10
			fmt.Println("A:", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	//输出结构中，全部A输出的结果都是10，B则随机输出0~9，这跟变量的作用域有关
	/*因为在A的for循环中，当全部的goroutine都登记完，for才退出，
	for退出时，被goroutine登记的函数开始执行时，此时for退出时i++，取得for退出时给出的参数i=10*/
	//不管设置了多少个内核线程，当goroutine有多个时，goroutine执行顺序总是随机的
}

package main

import (
	"context"
	"fmt"
	"time"
)

/*
字段	              含义
Deadline	         返回一个time.Time，表示当前Context的应该结束的时间，ok则表示是否有结束时间
Done	             当Context被取消或者超时时候返回的一个struct{}类型的只读channel
Err	                 返回context被取消时的错误
Value	             context实现共享数据存储的地方，是协程安全的,是context自带的k-v存储功能
*/

func f1(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func f2(ctx context.Context, x time.Duration) {
	c := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("调用后返回的时间为%v\n", time.Now().Second())
				return
			case c <- n:
				n++
			}
		}
	}()
	time.Sleep(x)

}

func f3(ctx context.Context, y time.Duration) {
	c := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("调用后返回的时间为%v\n", time.Now().Second())
				return
			case c <- n:
				n++
			}
		}
	}()
	time.Sleep(y)
}

func f4(ctx context.Context, key interface{}) {
	/* ret, ok := ctx.Value(key).(interface{})//根据key读取value的值
	if !ok {
		fmt.Println("找不到key对应的值")
	} else {
		fmt.Println(ret)
	} */
	ret1, ok := ctx.Value(key).(int) //根据key读取value的值
	if !ok {
		ret2, ok := ctx.Value(key).(string)
		if !ok {
			fmt.Println("找不到key对应的值")
		} else {
			fmt.Println(ret2)
		}
	} else {
		fmt.Println(ret1)
	}

}

// 为了更方便的创建Context，包里头定义了Background来作为所有Context的父节点，它是一个emptyCtx的实例。
func main() {
	//............................................................
	// WithCancel()
	// WithCancel返回带有新Done通道的父节点的副本。
	// 当Context被取消或者超时的时候返回的Done是一个关闭的通道。(当调用返回的cancel函数，或关闭父上下文的Done通道时，将关闭返回上下文的Done通道，以先发生的方式为准。)
	ctx1, cancel1 := context.WithCancel(context.Background()) //传入background父节点，返回ctx父节点副本

	for n := range f1(ctx1) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// 当我们取完需要的整数后调用cancel
	cancel1()
	//............................................................
	// WithDeadline()
	// 返回父上下文的副本，并将deadline调整为不迟于d。
	// 如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。
	// 当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。
	d := time.Now().Add(5 * time.Second) //定义函数截止工作的时刻
	ctx2, cancel2 := context.WithDeadline(context.Background(), d)

	fmt.Printf("现在的时间为%v\n", time.Now().Second())
	x := 7 * time.Second
	f2(ctx2, x)

	// 以到截止时间取消context还是调用cancel取消context，以最先发生的情况为准
	cancel2()
	//............................................................
	// WithTimeout
	// 取消此上下文将释放与其相关的资源，
	// 因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制
	dm := 2 * time.Second //定义函数执行的时间长度
	ctx3, cancel3 := context.WithTimeout(context.Background(), dm)
	fmt.Printf("现在的时间为%v\n", time.Now().Second())
	y := 6 * time.Second
	f3(ctx3, y)
	cancel3()
	//............................................................
	// WithValue
	// k-v存值
	// 返回父节点的副本，其中与key关联的值为val,且为避免context的包冲突，key不能是内置类型，只能自己定义类型
	type key string                                         //定义一个类型
	var k1 key = "hello"                                    //声明key的类型与对应的值
	var v1 = 22222                                          //根据key给value赋值
	ctx4 := context.WithValue(context.Background(), k1, v1) //首次在context的父节点的副本存值
	var k2 key = "ok"
	var v2 = "hello"
	ctx4 = context.WithValue(ctx4, k2, v2) //接着在context的父节点副本的副本中再存值，此时副本的副本存了两条k-v记录
	f4(ctx4, k1)                           //调用函数取得k1对应的值v1
	f4(ctx4, k2)                           //调用函数取得k2对应的值v2
}

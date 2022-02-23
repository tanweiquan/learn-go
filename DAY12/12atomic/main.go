package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup

//atomic
//atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。
//这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。

// var lock sync.Mutex
func add() {
	// lock.Lock()
	// x++
	// lock.Unlock()
	atomic.AddInt64(&x, 1) //相当于上面有锁，且x=x+1
	wg.Done()
}
func main() {
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}

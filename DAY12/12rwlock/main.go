package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁：rwlock
// 使用场景：读大于写（读多写少）
// 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。
/* 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，
当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。
读写锁在Go语言中使用sync包中的RWMutex类型。 */

/* 读写锁分为两种:读锁和写锁。
当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。 */

var (
	x  int64
	wg sync.WaitGroup
	// lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()    // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(5 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                  // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock() // 加读锁
	fmt.Println(x)
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now() //声明了start是程序开始时的时间
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()           //声明了end是程序结束时的时间
	fmt.Println(end.Sub(start)) //打印start到end的时间，即程序运行时间
}

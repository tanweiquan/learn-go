package main

import "fmt"

//闭包(闭包=函数+外部变量的引用),闭包就是能够读取其他函数内部变量的函数。
//闭包是一个函数，这个函数包含了他外部作用域的变量

// 闭包的两个定义：
// 第一：必须要有在函数外部定义，但在函数内部引用的“自由变量”
// 第二：脱离了形成闭包的上下文，闭包也能照常使用这些自由变量

//底层的原理
//1.函数可以作为返回值
//2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找

func create() func(int) int {
	var x int
	return func(y int) int { //返回值为一个函数，函数类型为func(int)int
		x += y
		return x
	}
}
func main() {
	/* 函数create的返回值是一个函数，但这个函数内部使用了外部定义的变量x，
	即使create执行结束，通过ret依然能正常调用这个闭包函数。并使用定义在f1函数内部的局部变量x，所以这里符合闭包的定义。
	通常称这个变量x为捕获变量。 */
	f1 := create()     //f1为func(y int)int{...} //注意这里f1是一个闭包
	fmt.Println(f1(1)) //这里传1进去当做y的值，打印结果为1
	fmt.Println(f1(1)) //这里传1进去当做y的值，打印结果为2
	f2 := create()     //f2为func(y int)int{...} //注意这里f2是另外一个闭包，与f1不是同一个闭包
	fmt.Println(f2(1)) //这里传1进去当做y的值，打印结果为1
	fmt.Println(f2(1)) //这里传1进去当做y的值，打印结果为2
	f3 := create()     //f1为func(y int)int{...}
	f4 := create()     //f2为func(y int)int{...}
	fmt.Println(f3(1)) //这里传1进去当做y的值，打印结果为1
	fmt.Println(f4(1)) //这里传1进去当做y的值，打印结果为1
	fmt.Println(f4(1)) //这里传1进去当做y的值，打印结果为2
	fmt.Println(f3(1)) //这里传1进去当做y的值，打印结果为2
	//以上实例证明了闭包的重点知识：在函数返回变量后，被返回的变量叫捕获变量，可以同一个闭包所捕获调用
}
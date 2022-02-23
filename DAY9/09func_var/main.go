package main

import "fmt"

//变量的作用域
var x = 100 //定义一个全局变量
//定义一个函数
//函数内部定义的变量只能该函数内部使用
//如果局部变量和全局变量重名，优先访问局部变量。
func f1() {
	//函数中寻找变量的顺序：
	//1、先在函数内部查找
	//2、找不到就往函数外部找，一直找到全局
	fmt.Println(x)
}
func main() {
	f1()
	//语句块中的变量的作用域
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
	if i := 10; i < 20 {
		fmt.Println("hello")
	}
}

package main

import "fmt"

/*
Go语言中推荐使用驼峰式命名：
var myFistName string (小驼峰，第一个单词小写，其他单词大写)
var MyFistName string (大驼峰，第一个单词大写，其他单词也大写)
*/
//可以全局声明变量的
var a int     //普通的变量声明，此时a的赋值要放在函数里了，格式：var 变量名 变量类型
var b int = 2 //声明的同时赋值，格式变量名 变量类型 = 表达式
var c = 3     //类型推导，电脑根据值判断变量类型
//快速声明变量上面的形式，此声明形式都可以用    (常用)
var (
	d = 4
	e = 5
	f = 6
)

//一行声明多个变量
var i1, i2 = "你好", 9

//匿名变量：下划线_
func foo() (int, string) {
	return 10, "quan"
}
func main() {
	x, _ := foo() //x=10
	_, y := foo() //y=quan
	a = 1         //a在函数里赋值
	g := 7        //简短声明变量，只能在函数中用     （常用）

	fmt.Printf("a:%d\n", a) //格式化输出，其中%d是占位符，使用a这个变量去替换占位符,仅打印""中的内容，\n是换行符，
	fmt.Println(b)          //打印后加一个换行符
	fmt.Print(c)            //普通打印
	fmt.Println(d, e, f, g) //注意每个变量声明后一定要使用
	fmt.Println(i1, i2)
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

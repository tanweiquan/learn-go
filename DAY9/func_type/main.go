package main

import "fmt"

//函数类型
func f1() { //没有参数，也没有返回值的函数，其函数类型为func()
	fmt.Println("hello,沙河")
}
func f2() int { //没有参数，有返回值的函数，其函数类型为func()int
	return 10
}

//函数也可以作为参数的类型
func f3(x func() int) { //参数类型为func()int,但没有返回值的函数,其函数类型为func(func() int)
	ret := x()
	fmt.Println(ret)
}
func ff(a, b int) int { //参数a,b的类型为int,返回值为int的函数
	return a + b
}

//函数也可以作为返回值
func f4(x func() int) func(int, int) int { //参数类型为func()int，返回值类型为func(int,int)int的函数
	return ff //其函数类型为func(func() int) func(int, int) int
}

func f5(x func() int) func(int, int) int { //这个函数跟上面的函数f4是一样的
	ret := func(a, b int) int { //注意这里时因为在函数里，所以声明的函数只能是匿名函数
		return a + b
	}
	return ret
}

//注意了，b:=f2中b接收的是整个函数；b1:=f2(),代表的是b1接收了f1()的返回值，函数名带括号是调用的意思
func main() {
	a := f1
	fmt.Printf("%T\n", a) //函数类型为func()
	b := f2               //注意这里f2后不带括号，因此a接收的是整个函数
	fmt.Printf("%T\n", b) //函数类型为func() int
	c := f3
	fmt.Printf("%T\n", c) //函数类型为func(func() int)
	d := f4
	fmt.Printf("%T\n", d) //函数类型为func(func() int) func(int, int) int
	e := f4(f2)           //因为函数f2的类型为func()int，符合函数f4的参数类型，因此可以作为参数传入
	fmt.Printf("%T\n", e) //函数类型为func(int, int) int
	f := f5
	fmt.Printf("%T\n", f) //函数类型为func(func() int) func(int, int) int
}

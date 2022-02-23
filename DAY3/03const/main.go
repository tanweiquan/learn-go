package main

import "fmt"

//常量
//定义了常量之后就不能修改
//程序运行期间不会改变的量
//常量声明const
const pi = 3.1415
const e = 2.7182
const (
	f = 13
	g = 14
)
const (
	h = 15 //15
	i      //15
	j      //15
	k = 10
)

//iota是常量计算器。1、当出现const关键字时，iota将被重置为零。2、const中每新增一行常量声明，iota就计数一次
const (
	a  = iota //0
	b  = iota //1
	c  = iota //2
	n  = 100  //100，此时iota计数变为3
	n1 = iota //4
	_  = iota //5，匿名常量，赋值但丢弃不使用
	m  = iota //6
	o         //7
)

//多个常量声明在一行
const (
	a1, a2 = iota + 1, iota + 2 //a1:1,a2:2
	b1, b2 = iota + 1, iota + 2 //b1:2,b2:3
)

//定义数量级,符号<<表示左移,<<n表示左移n位,<<n=*(2^n)
const (
	_  = iota             //0
	KB = 1 << (10 * iota) //1<<(10*iota)=1*（2^(10*1))=2^10=1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	const l = 18 //在函数里也可以声明常量
	fmt.Println(pi + 1)
	fmt.Println(e + 1)
	fmt.Println(f, g, h, i, j)
	fmt.Println(a, b, c)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(n, n1, m, o)
	fmt.Println(a1, a2, b1, b2)
	fmt.Println(KB, MB, GB, TB, PB)
}

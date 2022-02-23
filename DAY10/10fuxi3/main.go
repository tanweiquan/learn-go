package main

import "fmt"

//结构体复习3
//结构体是值类型，作为参数传参时相当于拷贝副本
type person struct {
	name string
	age  int
}

//不同的传参区别

//1.值接收者
// func (p person) guonian() {
// 	p.age++ //改的是p1的副本
// }

//2.指针接收者
//2.1.需要修改结构体变量的值的时候；
//2.2.结构体本身比较大，拷贝占用内存大的时候；
//2.3.保持一致性，如果一个方法使用了指针接收者，其他方法也要使用指针接收者。
func (p *person) zhenguonian() { //*person是指针，传参的副本的内存地址指向的是p2的值，从而能跳过副本到达值，修改了值
	p.age++ //根据指针修改了p2的值
}
func main() {
	//声明变量
	// p1 := person{
	// 	name: "小黄",
	// 	age:  5,
	// }
	p2 := person{
		name: "小黑",
		age:  5,
	}
	//调用方法
	// p1.guonian()
	p2.zhenguonian() // 相当于(&p2).zhenguonian()
	//查看区别
	// fmt.Println(p1.age) //5
	fmt.Println(p2.age) //6
}

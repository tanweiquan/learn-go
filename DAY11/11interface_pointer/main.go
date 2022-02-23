package main

import "fmt"

//使用值接收者和使用指针接收者的区别
//1、使用值接收者实现接口：能存：结构体类型的变量、结构体指针类型的变量
//2、使用指针接收者实现接口：只能存：结构体指针类型的变量
//注意，我们常用的是指针接收者来实现方法的，因此使用指针接收者来实现接口一定要熟练
type animal interface { //定义一个animal类型的接口
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}
type chicken struct {
	feet int8
}

//使用值接收者实现了接口的所有方法
func (c cat) move() {
	fmt.Println("走猫步！")
}
func (c cat) eat(food string) {
	fmt.Printf("吃%s\n", food)
}

//使用指针接收者实现了接口的所有方法（常用）
func (c *chicken) move() {
	fmt.Println("鸡动！")
}
func (c *chicken) eat(food string) {
	fmt.Printf("吃%s\n", food)
}

func main() {
	var a1 animal

	//值接收者:既能存结构体类型变量，也能存结构体指针类型的变量
	c1 := cat{"tom", 4}
	a1 = c1         //cat
	fmt.Println(a1) //{tom 4}
	c2 := &cat{"lanmao", 4}
	a1 = c2         //&cat
	fmt.Println(a1) //&{lanmao 4}

	//指针接收者：只能存结构体指针类型的变量(常用)
	ch1 := chicken{2}
	a1 = &ch1       //&ch1
	fmt.Println(a1) //&{2}
}

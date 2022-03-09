package main

import "fmt"

//接口(一种类型)

/*格式：
type 接口名 interface{
	方法名1(参数1，参数2...)(返回值1，返回值2...)
	方法名2(参数1，参数2...)(返回值1，返回值2...)
	...
}
*/
//接口的实现
/*
一个变量如果实现了接口中规定的所有的方法，那么这个变量就实现了这个接口，可以称为这个接口类型的变量
但如果是一个变量的指针实现接口的方法，那则是这个这个变量的指针是接口类型，而变量不是接口类型，但变量也可以调用指针的方法
*/

//接口是一种特殊的类型，它规定了变量有哪些方法
//解决了当我们不关心一个变量什么类型，只关心能调用它的什么方法

//没有规定传什么数据，只规定了用了什么方法
//引出接口的实例
//定义一个能叫的类型
type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型的(方法的签名，可以有一个，也可以有多个)
}
type cat struct{}
type dog struct{}
type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}
func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

//不管传进来什么参数，都能调用方法
func da(x speaker) { //定义了一个speaker类型的变量x
	//接收一个参数，传进来什么，我就打什么
	x.speak() //挨打了就叫
}
func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(c1)
	da(d1)
	da(p1)

	var ss speaker //定义一个接口类型为speaker的变量：ss
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss) //{}
}

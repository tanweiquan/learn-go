package main

import "fmt"

//接口的实现
//接口的实现：通过接口类型的变量保存数据(动态类型/动态值)
//接口是一种类型(特殊的类型)

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

func (c cat) move() {
	fmt.Println("走猫步！")
}
func (c cat) eat(food string) {
	fmt.Printf("吃%s\n", food)
}
func (c chicken) move() {
	fmt.Println("鸡动！")
}
func (c chicken) eat(food string) {
	fmt.Printf("吃%s\n", food)
}
func main() {

	//接口类型的变量能存值：type a interface{}; var b a; c=b;这里c的值的类型和值本身都赋值给接口类型的变量a了
	//接口类型的变量保存的分为两部分：值的类型和值本身  (动态类型/动态值)
	//这样就实现了接口变量能存储不同的值，没给接口变量赋类型和值前，其接口变量类型为<nil>,即里面的类型和值都是空的

	var a1 animal //定义一个接口类型变量a1
	tq := cat{    //定义一个cat类型的变量bc
		name: "淘气",
		feet: 4,
	}
	a1 = tq         //接口保存的分为两部分：值的类型和值本身
	fmt.Println(a1) //{淘气 4}
	a1.eat("猫粮")    //吃猫粮
	a1.move()       //走猫步!
	ji := chicken{
		feet: 2,
	}
	a1 = ji
	fmt.Println(a1) //{2}
	a1.eat("鸡饲料")   //吃鸡饲料
	a1.move()       //鸡动!

}

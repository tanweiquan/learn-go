package main

import "fmt"

//同一结构体可以实现多个接口
//接口还可以嵌套
type animal interface {
	mover
	eater
}
type mover interface { //其他人或者官方实现接口时经常加er：如mover
	move()
}
type eater interface {
	eat(string)
}
type cat struct {
	name string
	feet int8
}

//cat实现了mover接口
func (c cat) move() {
	fmt.Println("走猫步！")
}

//cat实现了eater接口
func (c cat) eat(food string) {
	fmt.Printf("吃%s\n", food)
}
func main() {
	var a1 animal
	c1 := cat{"tom", 4}
	a1 = c1
	fmt.Println(a1) //{tom 4}
	c2 := &cat{"lanmao", 4}
	a1 = c2
	fmt.Println(a1) //&{lanmao 4}
}

package main

import "fmt"

//给自定义类型加方法
//不能给别的包里面的类型添加方法，只能给自己包里的自定义类型添加方法

/*只能给自己自定义类型加方法，因此下面这里的int是外面的包，不能这样
 func (a int) sayhello() {
	fmt.Println("hello,这是一个int")
} */
type myint int

func (i myint) hello() {
	fmt.Println("hello,这是一个int")
}
func main() {
	m := myint(100) //相当于var m myint = 100
	m.hello()
}

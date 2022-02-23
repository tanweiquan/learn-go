package main

import "fmt"

//除了var name struct{}格式外，还可以构造函数
//结构体属于值类型
type person struct {
	x string
	y int
}

//构造函数:使用时用约定俗成的new开头的函数名

//返回的是结构体还是结构体指针是需要考量的
//当结构体很小，占用内存小，返回可以为结构体。当结构体很大，占用内存很多，可以返回指针

//这里返回的是结构体
func new1Person(x string, y int) person {
	return person{
		x: x,
		y: y,
	}
}

//这里返回的是指针
func new2Person(x string, y int) *person {
	return &person{
		x: x,
		y: y,
	}
}
func main() {
	p1 := new1Person("元帅", 18)
	p2 := new1Person("周琳", 20)
	fmt.Println(p1, p2) //{元帅 18} {周琳 20}
	p3 := *new2Person("哈哈", 18)
	p4 := *new2Person("呵呵", 20)
	fmt.Println(p3, p4) //{哈哈 18} {呵呵 20}
}

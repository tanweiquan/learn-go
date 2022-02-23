package main

import "fmt"

//结构体模拟实现其他语言中的"继承"

type animal struct {
	name string
}

//给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//狗类
type dog struct {
	feet   uint8
	animal //因为结构体嵌套，animal有的方法，dog此时也有了(继承)
}

//给狗实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪~\n", d.name)
}

func main() {
	//造狗
	d1 := dog{
		4,
		animal{
			"小黄",
		},
	}
	fmt.Println(d1) //{4 {小黄}}
	d1.move()       //小黄会动！
	d1.wang()       //小黄在叫：汪汪汪~
}

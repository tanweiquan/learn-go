package main

import "fmt"

//复习结构体
type tmp struct { //有名字的结构体
	x int
	y int
}
type person struct {
	name string
	age  int
}

//当结构体字段很多且要重复赋值给结构体时，可以用构造(结构体变量)函数，返回值是对应的结构体类型
func newPerson(n string, m int) person {
	return person{
		name: n,
		age:  m,
	}
}
func main() {
	var a = tmp{
		10,
		20,
	}
	fmt.Println(a)

	var b = struct { //匿名结构体
		x int
		y int
	}{10, 20} //传值{10，20}给x,y,跟上面类似
	fmt.Println(b)

	//因为有
	var x int32 = 12
	y := int32(12)
	fmt.Println(x, y)

	//所以
	//1.
	var p1 person //结构体实例化
	p1.name = "鲍德里"
	p1.age = 23
	//2.
	p2 := person{name: "鲍德里", age: 23} //结构体实例化
	//3.
	//调用构造函数生成person类型变量
	p3 := newPerson("娜扎", 26)
	fmt.Println(p1, p2, p3)
}

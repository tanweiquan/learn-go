package main

import "fmt"

//复习结构体2

type person struct {
	name string
	age  int
}

//方法
//接收者使用对应类型的首字母小写
//指定了接收者之后，只有接收者这个类型的变量才能调用这个方法
func (p person) dream(str string) { //方法名是dream
	fmt.Printf("%s的梦想是%s", p.name, str)
}
func main() {
	//声明p1是person类型的变量,并给字段赋值
	p1 := person{
		name: "周玲",
		age:  26,
	}
	//p1调用方法
	p1.dream("学好go语言") //周玲的梦想是学好go语言

}

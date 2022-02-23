package main

import (
	"fmt"
)

//方法
/*标准格式：
func (接收者变量 接收者类型) 方法名 (参数列表)(返回值参数){
	函数体
}
*/

//标识符：变量名 函数名 类型名 方法名
//go语言中如果标识符首字母是大写的，就表示对外部包可见的(暴露的，公有的，能被外部包调用的)
//如fmt.Pringln()的P大写，只有这样才能访问外部fmt包

/*对于这种标识符大写的，要写注释，具体注释格式如下：
//Dog 这是一个狗的结构体
type Dog struct {
	name string
}
*/
type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法是作用于特定类型的函数(在函数名前加接收者,即作用于接收者)
//接收者表示的是调用该方法的具体类型变量，多用类型名的首字母小写表示
//此时方法wang绑定结构体dog
func (d dog) wang() { //作用于dog类型的接收者d
	fmt.Printf("%s:汪汪汪~", d.name)
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//值接收者：传拷贝进去
/*
func (p person) guonian() {
	p.age++
}
*/

//指针接收者：传内存地址进去
func (p *person) zhenguonian() {
	p.age++
}

//使用指针接收者：
//1.需要修改接收者中的值
//2.接收者是拷贝代价比较大的对象
//3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用

func main() {
	d1 := newDog("周琳")
	d1.wang() //周琳:汪汪汪
	fmt.Println("......")
	p1 := newPerson("元帅", 18)
	fmt.Println(p1.age) //18
	// p1.guonian()     //传值进去
	// fmt.Println(p1.age) //18
	p1.zhenguonian()    //传内存地址进去
	fmt.Println(p1.age) //19
}

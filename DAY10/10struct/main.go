package main

import "fmt"

//结构体定义：一种可以封装多个基本数据类型的自定义数据类型
/*结构体的形式如下：
type 类型名 struct{
    字段名 字段类型
    字段名 字段类型
    ...
}
*/
//类型名：标识自定义结构体的名称，在同一个包内不能重复
//字段名：表示结构体字段名。结构体中的字段名必须唯一,不能重复。
//字段类型：表示结构体字段的具体类型
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个person类型变量p
	var p person
	//通过字段赋值
	p.name = "周玲"
	p.age = 26
	p.gender = "女"
	p.hobby = []string{"篮球", "游戏", "足球"}
	fmt.Println(p)        //{周玲 26 女 [篮球 游戏 足球]}
	fmt.Printf("%T\n", p) //类型：main.person
	//访问变量p的字段
	fmt.Println(p.name) //周玲
	//另一个例子：
	var p2 person
	p2.name = "沙河"
	p2.age = 26
	fmt.Println(p2)        //{沙河 26  []}
	fmt.Printf("%T\n", p2) //类型：main.person
	//匿名结构体：没有名字的结构体，适合在函数内临时定义的(注意了这里直接声明，没有用type)
	var s struct {
		x string
		y int
	}
	s.x = "哈哈哈"
	s.y = 26
	fmt.Println(s)        //{哈哈哈 26}
	fmt.Printf("%T\n", s) //类型：struct { x string; y int }
}
